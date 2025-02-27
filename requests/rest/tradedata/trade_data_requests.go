package tradedata

import "github.com/liuhengloveyou/okx-go"

type (
	GetTakerVolume struct {
		Ccy      string             `json:"ccy"`
		Begin    int64              `json:"before,omitempty,string"`
		End      int64              `json:"limit,omitempty,string"`
		InstType okx.InstrumentType `json:"instType,string"`
		Period   okx.BarSize        `json:"period,string,omitempty"`
	}
	GetRatio struct {
		Ccy    string      `json:"ccy"`
		Begin  int64       `json:"before,omitempty,string"`
		End    int64       `json:"limit,omitempty,string"`
		Period okx.BarSize `json:"period,string,omitempty"`
	}
	GetOpenInterestAndVolumeStrike struct {
		Ccy     string      `json:"ccy"`
		ExpTime string      `json:"expTime"`
		Period  okx.BarSize `json:"period,string,omitempty"`
	}
)
