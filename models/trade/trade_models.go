package trade

import "github.com/liuhengloveyou/okx-go"

type (
	PlaceOrder struct {
		OrdID   string        `json:"ordId"`
		ClOrdID string        `json:"clOrdId"`
		Tag     string        `json:"tag"`
		SMsg    string        `json:"sMsg"`
		SCode   okx.JSONInt64 `json:"sCode"`
	}
	CancelOrder struct {
		OrdID   string          `json:"ordId"`
		ClOrdID string          `json:"clOrdId"`
		SMsg    string          `json:"sMsg"`
		SCode   okx.JSONFloat64 `json:"sCode"`
	}
	AmendOrder struct {
		OrdID   string          `json:"ordId"`
		ClOrdID string          `json:"clOrdId"`
		ReqID   string          `json:"reqId"`
		SMsg    string          `json:"sMsg"`
		SCode   okx.JSONFloat64 `json:"sCode"`
	}
	ClosePosition struct {
		InstID  string           `json:"instId"`
		PosSide okx.PositionSide `json:"posSide"`
	}
	Order struct {
		InstID       string             `json:"instId"`
		Ccy          string             `json:"ccy"`
		OrdID        string             `json:"ordId"`
		AlgoID       string             `json:"algoId"`
		ClOrdID      string             `json:"clOrdId"`
		AlgoClOrdID  string             `json:"algoClOrdId"`
		TradeID      string             `json:"tradeId"`
		Tag          string             `json:"tag"`
		Category     string             `json:"category"`
		FeeCcy       string             `json:"feeCcy"`
		RebateCcy    string             `json:"rebateCcy"`
		QuickMgnType string             `json:"quickMgnType"`
		ReduceOnly   string             `json:"reduceOnly"`
		Px           okx.JSONFloat64    `json:"px"`
		Sz           okx.JSONFloat64    `json:"sz"`
		Pnl          okx.JSONFloat64    `json:"pnl"`
		AccFillSz    okx.JSONFloat64    `json:"accFillSz"`
		FillPx       okx.JSONFloat64    `json:"fillPx"`
		FillSz       okx.JSONFloat64    `json:"fillSz"`
		FillTime     okx.JSONFloat64    `json:"fillTime"`
		AvgPx        okx.JSONFloat64    `json:"avgPx"`
		Lever        okx.JSONFloat64    `json:"lever"`
		TpTriggerPx  okx.JSONFloat64    `json:"tpTriggerPx"`
		TpOrdPx      okx.JSONFloat64    `json:"tpOrdPx"`
		SlTriggerPx  okx.JSONFloat64    `json:"slTriggerPx"`
		SlOrdPx      okx.JSONFloat64    `json:"slOrdPx"`
		Fee          okx.JSONFloat64    `json:"fee"`
		Rebate       okx.JSONFloat64    `json:"rebate"`
		State        okx.OrderState     `json:"state"`
		TdMode       okx.TradeMode      `json:"tdMode"`
		PosSide      okx.PositionSide   `json:"posSide"`
		Side         okx.OrderSide      `json:"side"`
		OrdType      okx.OrderType      `json:"ordType"`
		InstType     okx.InstrumentType `json:"instType"`
		TgtCcy       okx.QuantityType   `json:"tgtCcy"`
		UTime        okx.JSONTime       `json:"uTime"`
		CTime        okx.JSONTime       `json:"cTime"`
	}
	TransactionDetail struct {
		InstID   string             `json:"instId"`
		OrdID    string             `json:"ordId"`
		TradeID  string             `json:"tradeId"`
		ClOrdID  string             `json:"clOrdId"`
		BillID   string             `json:"billId"`
		Tag      okx.JSONFloat64    `json:"tag"`
		FillPx   okx.JSONFloat64    `json:"fillPx"`
		FillSz   okx.JSONFloat64    `json:"fillSz"`
		FillPnl  string             `json:"fillPnl"`
		FillTime okx.JSONTime       `json:"fillTime"`
		Fee      okx.JSONFloat64    `json:"fee"`
		FeeCcy   string             `json:"feeCcy"`
		InstType okx.InstrumentType `json:"instType"`
		Side     okx.OrderSide      `json:"side"`
		PosSide  okx.PositionSide   `json:"posSide"`
		ExecType okx.OrderFlowType  `json:"execType"`
		TS       okx.JSONTime       `json:"ts"`
	}
	PlaceAlgoOrder struct {
		AlgoID string        `json:"algoId"`
		SMsg   string        `json:"sMsg"`
		SCode  okx.JSONInt64 `json:"sCode"`
	}
	CancelAlgoOrder struct {
		AlgoID string        `json:"algoId"`
		SMsg   string        `json:"sMsg"`
		SCode  okx.JSONInt64 `json:"sCode"`
	}
	AlgoOrder struct {
		InstID          string             `json:"instId"`
		Ccy             string             `json:"ccy"`
		OrdID           string             `json:"ordId"`
		AlgoID          string             `json:"algoId"`
		ClOrdID         string             `json:"clOrdId"`
		TradeID         string             `json:"tradeId"`
		Tag             string             `json:"tag"`
		Category        string             `json:"category"`
		FeeCcy          string             `json:"feeCcy"`
		RebateCcy       string             `json:"rebateCcy"`
		TimeInterval    string             `json:"timeInterval"`
		QuickMgnType    string             `json:"quickMgnType"`
		ReduceOnly      string             `json:"reduceOnly"`
		Px              okx.JSONFloat64    `json:"px"`
		PxVar           okx.JSONFloat64    `json:"pxVar"`
		PxSpread        okx.JSONFloat64    `json:"pxSpread"`
		PxLimit         okx.JSONFloat64    `json:"pxLimit"`
		Sz              okx.JSONFloat64    `json:"sz"`
		SzLimit         okx.JSONFloat64    `json:"szLimit"`
		ActualSz        okx.JSONFloat64    `json:"actualSz"`
		ActualPx        okx.JSONFloat64    `json:"actualPx"`
		Pnl             okx.JSONFloat64    `json:"pnl"`
		AccFillSz       okx.JSONFloat64    `json:"accFillSz"`
		FillPx          okx.JSONFloat64    `json:"fillPx"`
		FillSz          okx.JSONFloat64    `json:"fillSz"`
		FillTime        okx.JSONFloat64    `json:"fillTime"`
		AvgPx           okx.JSONFloat64    `json:"avgPx"`
		Lever           okx.JSONFloat64    `json:"lever"`
		TpTriggerPx     okx.JSONFloat64    `json:"tpTriggerPx"`
		TpOrdPx         okx.JSONFloat64    `json:"tpOrdPx"`
		SlTriggerPx     okx.JSONFloat64    `json:"slTriggerPx"`
		SlOrdPx         okx.JSONFloat64    `json:"slOrdPx"`
		TpTriggerPxType string             `json:"tpTriggerPxType"`
		SlTriggerPxType string             `json:"slTriggerPxType"`
		TriggerPx       okx.JSONFloat64    `json:"triggerPx"`
		CallbackRatio   okx.JSONFloat64    `json:"callbackRatio"`
		CallbackSpread  okx.JSONFloat64    `json:"callbackSpread"`
		ActivePx        okx.JSONFloat64    `json:"activePx"`
		OrdPx           okx.JSONFloat64    `json:"ordPx"`
		Fee             okx.JSONFloat64    `json:"fee"`
		Rebate          okx.JSONFloat64    `json:"rebate"`
		State           okx.OrderState     `json:"state"`
		TdMode          okx.TradeMode      `json:"tdMode"`
		ActualSide      okx.PositionSide   `json:"actualSide"`
		PosSide         okx.PositionSide   `json:"posSide"`
		Side            okx.OrderSide      `json:"side"`
		OrdType         okx.AlgoOrderType  `json:"ordType"`
		InstType        okx.InstrumentType `json:"instType"`
		TgtCcy          okx.QuantityType   `json:"tgtCcy"`
		CTime           okx.JSONTime       `json:"cTime"`
		TriggerTime     okx.JSONTime       `json:"triggerTime"`
	}
	FromData struct {
		Ccy    string          `json:"fromCcy"`
		Amount okx.JSONFloat64 `json:"fromAmt""`
	}
	EasyConvertListResult struct {
		FromData []FromData `json:"fromData"`
		ToCcy    []string   `json:"toCcy"`
	}
	EasyConvertProcess struct {
		FromCcy    string          `json:"fromCcy"`
		ToCcy      string          `json:"toCcy"`
		FillFromSz okx.JSONFloat64 `json:"fillFromSz"`
		FillToSz   okx.JSONFloat64 `json:"fillToSz"`
		Status     string          `json:"status"`
		UpdateTime okx.JSONTime    `json:"uTime"`
	}
)

type EasyConvertSource string

const EasyConvertSourceFunding = EasyConvertSource("2")
const EasyConvertSourceTrading = EasyConvertSource("1")
