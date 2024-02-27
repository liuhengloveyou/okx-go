package market

import (
	"encoding/json"
	"fmt"
	"github.com/drinkthere/okx"
	"strconv"
	"time"
)

type (
	Ticker struct {
		InstID    string             `json:"instId"`
		Last      okx.JSONFloat64    `json:"last"`
		LastSz    okx.JSONFloat64    `json:"lastSz"`
		AskPx     okx.JSONFloat64    `json:"askPx"`
		AskSz     okx.JSONFloat64    `json:"askSz"`
		BidPx     okx.JSONFloat64    `json:"bidPx"`
		BidSz     okx.JSONFloat64    `json:"bidSz"`
		Open24h   okx.JSONFloat64    `json:"open24h"`
		High24h   okx.JSONFloat64    `json:"high24h"`
		Low24h    okx.JSONFloat64    `json:"low24h"`
		VolCcy24h okx.JSONFloat64    `json:"volCcy24h"`
		Vol24h    okx.JSONFloat64    `json:"vol24h"`
		SodUtc0   okx.JSONFloat64    `json:"sodUtc0"`
		SodUtc8   okx.JSONFloat64    `json:"sodUtc8"`
		InstType  okx.InstrumentType `json:"instType"`
		TS        okx.JSONTime       `json:"ts"`
	}
	IndexTicker struct {
		InstID  string          `json:"instId"`
		IdxPx   okx.JSONFloat64 `json:"idxPx"`
		High24h okx.JSONFloat64 `json:"high24h"`
		Low24h  okx.JSONFloat64 `json:"low24h"`
		Open24h okx.JSONFloat64 `json:"open24h"`
		SodUtc0 okx.JSONFloat64 `json:"sodUtc0"`
		SodUtc8 okx.JSONFloat64 `json:"sodUtc8"`
		TS      okx.JSONTime    `json:"ts"`
	}
	OrderBook struct {
		Asks []*OrderBookEntity `json:"asks"`
		Bids []*OrderBookEntity `json:"bids"`
		TS   okx.JSONTime       `json:"ts"`
	}
	OrderBookWs struct {
		Asks     []*OrderBookEntity `json:"asks"`
		Bids     []*OrderBookEntity `json:"bids"`
		Checksum int                `json:"checksum"`
		TS       okx.JSONTime       `json:"ts"`
	}
	OrderBookEntity struct {
		DepthPrice      float64
		Size            float64
		LiquidatedOrder int
		OrderNumbers    int
	}
	Candle struct {
		O      float64
		H      float64
		L      float64
		C      float64
		Vol    float64
		VolCcy float64
		TS     okx.JSONTime
	}
	Candlesticks struct {
		TS          okx.JSONTime
		O           float64
		H           float64
		L           float64
		C           float64
		Vol         float64
		VolCcy      float64
		VolCcyQuote float64
		Confirm     int64
	}
	IndexCandle struct {
		O  float64
		H  float64
		L  float64
		C  float64
		TS okx.JSONTime
	}
	Trade struct {
		InstID  string          `json:"instId"`
		TradeID okx.JSONFloat64 `json:"tradeId"`
		Px      okx.JSONFloat64 `json:"px"`
		Sz      okx.JSONFloat64 `json:"sz"`
		Side    okx.TradeSide   `json:"side"`
		TS      okx.JSONTime    `json:"ts"`
	}
	TotalVolume24H struct {
		VolUsd okx.JSONFloat64 `json:"volUsd"`
		VolCny okx.JSONFloat64 `json:"volCny"`
		TS     okx.JSONTime    `json:"ts"`
	}
	IndexComponent struct {
		Index      string          `json:"index"`
		Last       okx.JSONFloat64 `json:"last"`
		Components []*Component    `json:"components"`
		TS         okx.JSONTime    `json:"ts"`
	}
	Component struct {
		Exch   string          `json:"exch"`
		Symbol string          `json:"symbol"`
		SymPx  okx.JSONFloat64 `json:"symPx"`
		Wgt    okx.JSONFloat64 `json:"wgt"`
		CnvPx  okx.JSONFloat64 `json:"cnvPx"`
	}
)

func (o *OrderBookEntity) UnmarshalJSON(buf []byte) error {
	var (
		dp, s, lo, on string
		err           error
	)
	tmp := []interface{}{&dp, &s, &lo, &on}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in OrderBookEntity: %d != %d", g, e)
	}
	o.DepthPrice, err = strconv.ParseFloat(dp, 64)
	if err != nil {
		return err
	}
	o.Size, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	o.LiquidatedOrder, err = strconv.Atoi(lo)
	if err != nil {
		return err
	}
	o.OrderNumbers, err = strconv.Atoi(on)
	if err != nil {
		return err
	}

	return nil
}

func (c *Candlesticks) UnmarshalJSON(buf []byte) error {
	var (
		o, h, l, cl, vol, volCcy, volCcyQuote, confirm, ts string
		err                                                error
	)
	tmp := []interface{}{&ts, &o, &h, &l, &cl, &vol, &volCcy, &volCcyQuote, &confirm}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.TS) = time.UnixMilli(timestamp)

	c.O, err = strconv.ParseFloat(o, 64)
	if err != nil {
		return err
	}

	c.H, err = strconv.ParseFloat(h, 64)
	if err != nil {
		return err
	}

	c.L, err = strconv.ParseFloat(l, 64)
	if err != nil {
		return err
	}

	c.C, err = strconv.ParseFloat(cl, 64)
	if err != nil {
		return err
	}

	c.Vol, err = strconv.ParseFloat(vol, 64)
	if err != nil {
		return err
	}

	c.VolCcy, err = strconv.ParseFloat(volCcy, 64)
	if err != nil {
		return err
	}

	c.VolCcyQuote, err = strconv.ParseFloat(volCcyQuote, 64)
	if err != nil {
		return err
	}
	confirmInt, err := strconv.Atoi(confirm)
	if err != nil {
		return err
	}
	c.Confirm = int64(confirmInt)

	return nil
}

func (c *Candle) UnmarshalJSON(buf []byte) error {
	var (
		o, h, l, cl, vol, volCcy, ts string
		err                          error
	)
	tmp := []interface{}{&ts, &o, &h, &l, &cl, &vol, &volCcy}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.TS) = time.UnixMilli(timestamp)

	c.O, err = strconv.ParseFloat(o, 64)
	if err != nil {
		return err
	}

	c.H, err = strconv.ParseFloat(h, 64)
	if err != nil {
		return err
	}

	c.L, err = strconv.ParseFloat(l, 64)
	if err != nil {
		return err
	}

	c.C, err = strconv.ParseFloat(cl, 64)
	if err != nil {
		return err
	}

	c.Vol, err = strconv.ParseFloat(vol, 64)
	if err != nil {
		return err
	}

	c.VolCcy, err = strconv.ParseFloat(volCcy, 64)
	if err != nil {
		return err
	}

	return nil
}

func (c *IndexCandle) UnmarshalJSON(buf []byte) error {
	var (
		o, h, l, cl, ts string
		err             error
	)
	tmp := []interface{}{&ts, &o, &h, &l, &cl}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.TS) = time.UnixMilli(timestamp)

	c.O, err = strconv.ParseFloat(o, 64)
	if err != nil {
		return err
	}

	c.H, err = strconv.ParseFloat(h, 64)
	if err != nil {
		return err
	}

	c.L, err = strconv.ParseFloat(l, 64)
	if err != nil {
		return err
	}

	c.C, err = strconv.ParseFloat(cl, 64)
	if err != nil {
		return err
	}

	return nil
}
