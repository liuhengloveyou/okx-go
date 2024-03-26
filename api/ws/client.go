package ws

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/drinkthere/okx"
	"github.com/drinkthere/okx/events"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// ClientWs is the websocket api client
//
// https://www.okx.com/docs-v5/en/#websocket-api
type ClientWs struct {
	url           sync.Map
	apiKey        string
	secretKey     []byte
	passphrase    string
	conn          sync.Map
	ctx           context.Context
	Cancel        context.CancelFunc
	DoneChan      chan interface{}
	ErrChan       chan *events.Error
	SubscribeChan chan *events.Subscribe
	UnsubscribeCh chan *events.Unsubscribe
	LoginChan     chan *events.Login
	SuccessChan   chan *events.Success
	sendChan      sync.Map
	lastTransmit  sync.Map
	AuthRequested *time.Time
	Authorized    int32 // 0 for false, 1 for true
	Private       *Private
	Public        *Public
	Trade         *Trade
}

const (
	redialTick = 2 * time.Second
	writeWait  = 3 * time.Second
	pongWait   = 25 * time.Second
	PingPeriod = 15 * time.Second
)

// NewClient returns a pointer to a fresh ClientWs
func NewClient(ctx context.Context, apiKey, secretKey, passphrase string, url map[bool]okx.BaseURL) *ClientWs {
	ctx, cancel := context.WithCancel(ctx)
	c := &ClientWs{
		apiKey:     apiKey,
		secretKey:  []byte(secretKey),
		passphrase: passphrase,
		ctx:        ctx,
		Cancel:     cancel,
		DoneChan:   make(chan interface{}, 32),
	}

	c.url.Store(true, url[true])
	c.url.Store(false, url[false])
	c.sendChan.Store(true, make(chan []byte, 3))
	c.sendChan.Store(false, make(chan []byte, 3))
	c.lastTransmit.Store(true, time.Now())
	c.lastTransmit.Store(false, time.Now())
	c.Private = NewPrivate(c)
	c.Public = NewPublic(c)
	c.Trade = NewTrade(c)

	return c
}

// Connect into the server
//
// https://www.okx.com/docs-v5/en/#websocket-api-connect
func (c *ClientWs) Connect(p bool) error {
	if c.CheckConnect(p) {
		return nil
	}

	err := c.dial(p)
	if err == nil {
		return nil
	}

	ticker := time.NewTicker(redialTick)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err = c.dial(p)
			if err == nil {
				return nil
			}
		case <-c.ctx.Done():
			return c.handleCancel("connect")
		}
	}
}

// CheckConnect into the server
func (c *ClientWs) CheckConnect(p bool) bool {
	_, ok := c.conn.Load(p)
	return ok
}

// Login
//
// https://www.okx.com/docs-v5/en/#websocket-api-login
func (c *ClientWs) Login() error {
	if atomic.LoadInt32(&c.Authorized) == 1 {
		return nil
	}

	if c.AuthRequested != nil && time.Since(*c.AuthRequested).Seconds() < 30 {
		return nil
	}

	now := time.Now()
	c.AuthRequested = &now
	method := http.MethodGet
	path := "/users/self/verify"
	ts, sign := c.sign(method, path)
	args := []map[string]string{
		{
			"apiKey":     c.apiKey,
			"passphrase": c.passphrase,
			"timestamp":  ts,
			"sign":       sign,
		},
	}

	return c.Send(true, okx.LoginOperation, args)
}

// Subscribe
// Users can choose to subscribe to one or more channels, and the total length of multiple channels cannot exceed 4096 bytes.
//
// https://www.okx.com/docs-v5/en/#websocket-api-subscribe
func (c *ClientWs) Subscribe(p bool, ch []okx.ChannelName, args map[string]string) error {
	count := 1
	if len(ch) != 0 {
		count = len(ch)
	}

	tmpArgs := make([]map[string]string, count)
	tmpArgs[0] = args

	for i, name := range ch {
		tmpArgs[i] = map[string]string{}
		tmpArgs[i]["channel"] = string(name)
		for k, v := range args {
			tmpArgs[i][k] = v
		}
	}

	return c.Send(p, okx.SubscribeOperation, tmpArgs)
}

// Unsubscribe into channel(s)
//
// https://www.okx.com/docs-v5/en/#websocket-api-unsubscribe
func (c *ClientWs) Unsubscribe(p bool, ch []okx.ChannelName, args map[string]string) error {
	tmpArgs := make([]map[string]string, len(ch))
	for i, name := range ch {
		tmpArgs[i] = make(map[string]string)
		tmpArgs[i]["channel"] = string(name)
		for k, v := range args {
			tmpArgs[i][k] = v
		}
	}

	return c.Send(p, okx.UnsubscribeOperation, tmpArgs)
}

// Send message through either connections
func (c *ClientWs) Send(p bool, op okx.Operation, args []map[string]string, extras ...map[string]string) error {
	if op != okx.LoginOperation {
		err := c.Connect(p)
		if err != nil {
			return err
		}
		if p {
			err = c.WaitForAuthorization()
			if err != nil {
				return err
			}
		}
	}

	data := map[string]interface{}{
		"op":   op,
		"args": args,
	}

	for _, extra := range extras {
		for k, v := range extra {
			data[k] = v
		}
	}

	j, err := json.Marshal(data)
	if err != nil {
		return err
	}

	sendChanInterface, _ := c.sendChan.Load(p)
	sendChan := sendChanInterface.(chan []byte)
	sendChan <- j
	return nil
}

// SetChannels to receive certain events on separate channel
func (c *ClientWs) SetChannels(errCh chan *events.Error, subCh chan *events.Subscribe, unSub chan *events.Unsubscribe, lCh chan *events.Login, sCh chan *events.Success) {
	c.ErrChan = errCh
	c.SubscribeChan = subCh
	c.UnsubscribeCh = unSub
	c.LoginChan = lCh
	c.SuccessChan = sCh
}

// WaitForAuthorization waits for the auth response and try to log in if it was needed
func (c *ClientWs) WaitForAuthorization() error {
	if atomic.LoadInt32(&c.Authorized) == 1 {
		return nil
	}

	if err := c.Login(); err != nil {
		return err
	}

	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()

	for range ticker.C {
		if atomic.LoadInt32(&c.Authorized) == 1 {
			return nil
		}
	}

	return nil
}

func (c *ClientWs) dial(p bool) error {
	urlInterface, _ := c.url.Load(p)
	url := urlInterface.(okx.BaseURL)

	conn, res, err := websocket.DefaultDialer.Dial(string(url), nil)
	if err != nil {
		var statusCode int
		if res != nil {
			statusCode = res.StatusCode
		}
		return fmt.Errorf("error %d: %w", statusCode, err)
	}
	defer res.Body.Close()

	c.conn.Store(p, conn)

	go func() {
		defer func() {
			// Cleaning the connection with ws
			c.Cancel()
			conn.Close()
		}()
		err := c.receiver(p)
		if err != nil {
			if !strings.Contains(err.Error(), "operation cancelled: receiver") {
				c.ErrChan <- &events.Error{
					Event: "error",
					Msg:   err.Error(),
				}
				fmt.Printf("receiver error: %v\n", err)
			}
		}
	}()

	go func() {
		defer func() {
			// Cleaning the connection with ws
			c.Cancel()
			conn.Close()
		}()
		err := c.sender(p)
		if err != nil {
			if !strings.Contains(err.Error(), "operation cancelled: sender") {
				c.ErrChan <- &events.Error{
					Event: "error",
					Msg:   err.Error(),
				}
				fmt.Printf("sender error: %v\n", err)
			}
		}
	}()

	return nil
}

func (c *ClientWs) sender(p bool) error {
	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()

	for {
		sendChanInterface, _ := c.sendChan.Load(p)
		sendChan := sendChanInterface.(chan []byte)

		connInterface, _ := c.conn.Load(p)
		conn := connInterface.(*websocket.Conn)

		select {
		case data := <-sendChan:
			err := conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				return fmt.Errorf("failed to set write deadline for ws connection, error: %w", err)
			}

			w, err := conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return fmt.Errorf("failed to get next writer for ws connection, error: %w", err)
			}

			if _, err = w.Write(data); err != nil {
				return fmt.Errorf("failed to write data via ws connection, error: %w", err)
			}

			if err := w.Close(); err != nil {
				return fmt.Errorf("failed to close ws connection, error: %w", err)
			}
		case <-ticker.C:
			connInterface, _ := c.conn.Load(p)
			conn := connInterface.(*websocket.Conn)
			lastTransmitInterface, _ := c.lastTransmit.Load(p)
			lastTransmit := lastTransmitInterface.(*time.Time)
			if conn != nil && (lastTransmit == nil || (lastTransmit != nil && time.Since(*lastTransmit) > PingPeriod)) {
				go func() {
					sendChan <- []byte("ping")
				}()
			}
		case <-c.ctx.Done():
			return c.handleCancel("sender")
		}
	}
}

func (c *ClientWs) receiver(p bool) error {
	for {
		select {
		case <-c.ctx.Done():
			return c.handleCancel("receiver")
		default:
			connInterface, _ := c.conn.Load(p)
			conn := connInterface.(*websocket.Conn)
			err := conn.SetReadDeadline(time.Now().Add(pongWait))
			if err != nil {
				return fmt.Errorf("failed to set read deadline for ws connection, error: %w", err)
			}

			_, data, err := conn.ReadMessage()
			if err != nil {
				return fmt.Errorf("failed to read message from ws connection, error: %v\n", err)
			}

			now := time.Now()
			c.lastTransmit.Store(p, &now)

			if string(data) != "pong" {
				e := &events.Basic{}
				if err := json.Unmarshal(data, e); err != nil {
					return fmt.Errorf("failed to unmarshall message from ws, error: %w", err)
				}
				go c.process(data, e)
			}
		}
	}
}

func (c *ClientWs) sign(method, path string) (string, string) {
	t := time.Now().UTC().Unix()
	ts := fmt.Sprint(t)
	s := ts + method + path
	p := []byte(s)
	h := hmac.New(sha256.New, c.secretKey)
	h.Write(p)

	return ts, base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (c *ClientWs) handleCancel(msg string) error {
	go func() {
		c.DoneChan <- msg
	}()

	return fmt.Errorf("operation cancelled: %s", msg)
}

// TODO: break each case into a separate function
func (c *ClientWs) process(data []byte, e *events.Basic) bool {
	switch e.Event {
	case "error":
		e := events.Error{}
		_ = json.Unmarshal(data, &e)
		go func() {
			if c.ErrChan != nil {
				c.ErrChan <- &e
			}
		}()

		return true
	case "subscribe":
		e := events.Subscribe{}
		_ = json.Unmarshal(data, &e)
		if c.SubscribeChan != nil {
			c.SubscribeChan <- &e
		}

		return true
	case "unsubscribe":
		e := events.Unsubscribe{}
		_ = json.Unmarshal(data, &e)
		go func() {
			if c.UnsubscribeCh != nil {
				c.UnsubscribeCh <- &e
			}
		}()

		return true
	case "login":
		if time.Since(*c.AuthRequested).Seconds() > 30 {
			c.AuthRequested = nil
			_ = c.Login()
			break
		}

		atomic.StoreInt32(&c.Authorized, 1)

		e := events.Login{}
		_ = json.Unmarshal(data, &e)
		go func() {
			if c.LoginChan != nil {
				c.LoginChan <- &e
			}
		}()

		return true
	}

	if c.Private.Process(data, e) {
		return true
	}

	if c.Public.Process(data, e) {
		return true
	}

	if e.ID != "" {
		if e.Code != 0 {
			ee := *e
			ee.Event = "error"

			return c.process(data, &ee)
		}

		e := events.Success{}
		_ = json.Unmarshal(data, &e)

		if c.SuccessChan != nil {
			c.SuccessChan <- &e
		}

		return true
	}

	return false
}
