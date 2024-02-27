package funding

type (
	Currency struct {
		Ccy               string `json:"ccy"`
		Name              string `json:"name"`
		LogoLink          string `json:"logoLink"`
		Chain             string `json:"chain"`
		MinDep            string `json:"minDep"`
		MinWd             string `json:"minWd"`
		MaxWd             string `json:"maxWd"`
		MinFee            string `json:"minFee"`
		MaxFee            string `json:"maxFee"`
		CanDep            bool   `json:"canDep"`
		CanWd             bool   `json:"canWd"`
		CanInternal       bool   `json:"canInternal"`
		WdQuota           string `json:"wdQuota"`
		UsedWdQuota       string `json:"usedWdQuota"`
		DepQuotaFixed     string `json:"depQuotaFixed"`
		UsedDepQuotaFixed string `json:"usedDepQuotaFixed"`
	}

	Balance struct {
		Ccy       string `json:"ccy"`
		Bal       string `json:"bal"`
		FrozenBal string `json:"frozenBal"`
		AvailBal  string `json:"availBal"`
	}
	Transfer struct {
		TransID string          `json:"transId"`
		Ccy     string          `json:"ccy"`
		Amt     okx.JSONFloat64 `json:"amt"`
		From    okx.AccountType `json:"from,string"`
		To      okx.AccountType `json:"to,string"`
	}
	TransferState struct {
		TransID string          `json:"transId"`
		Ccy     string          `json:"ccy"`
		Amt     okx.JSONFloat64 `json:"amt"`
		Type    string          `json:"type"`
		From    okx.AccountType `json:"from,string"`
		To      okx.AccountType `json:"to,string"`
		SubAcct string          `json:"subAcct"`
		State   string          `json:"state"`
	}
	Bill struct {
		BillID string          `json:"billId"`
		Ccy    string          `json:"ccy"`
		Bal    okx.JSONFloat64 `json:"bal"`
		BalChg okx.JSONFloat64 `json:"balChg"`
		Type   okx.BillType    `json:"type,string"`
		TS     okx.JSONTime    `json:"ts"`
	}
	DepositAddress struct {
		Addr     string          `json:"addr"`
		Tag      string          `json:"tag,omitempty"`
		Memo     string          `json:"memo,omitempty"`
		PmtID    string          `json:"pmtId,omitempty"`
		Ccy      string          `json:"ccy"`
		Chain    string          `json:"chain"`
		CtAddr   string          `json:"ctAddr"`
		Selected bool            `json:"selected"`
		To       okx.AccountType `json:"to,string"`
		TS       okx.JSONTime    `json:"ts"`
	}
	DepositHistory struct {
		Ccy                 string           `json:"ccy"`
		Chain               string           `json:"chain"`
		TxID                string           `json:"txId"`
		From                string           `json:"from"`
		To                  string           `json:"to"`
		DepId               string           `json:"depId"`
		Amt                 okx.JSONFloat64  `json:"amt"`
		State               okx.DepositState `json:"state,string"`
		ActualDepBlkConfirm string           `json:"actualDepBlkConfirm"`
		TS                  okx.JSONTime     `json:"ts"`
	}
	Withdrawal struct {
		Ccy   string          `json:"ccy"`
		Chain string          `json:"chain"`
		WdID  okx.JSONInt64   `json:"wdId"`
		Amt   okx.JSONFloat64 `json:"amt"`
	}
	WithdrawalHistory struct {
		Ccy   string              `json:"ccy"`
		Chain string              `json:"chain"`
		TxID  string              `json:"txId"`
		From  string              `json:"from"`
		To    string              `json:"to"`
		Tag   string              `json:"tag,omitempty"`
		PmtID string              `json:"pmtId,omitempty"`
		Memo  string              `json:"memo,omitempty"`
		Amt   okx.JSONFloat64     `json:"amt"`
		Fee   okx.JSONFloat64     `json:"fee"`
		WdID  okx.JSONInt64       `json:"wdId"`
		State okx.WithdrawalState `json:"state,string"`
		TS    okx.JSONTime        `json:"ts"`
	}
	PiggyBank struct {
		Ccy  string          `json:"ccy"`
		Amt  okx.JSONFloat64 `json:"amt"`
		Side okx.ActionType  `json:"side,string"`
	}
	PiggyBankBalance struct {
		Ccy      string          `json:"ccy"`
		Amt      okx.JSONFloat64 `json:"amt"`
		Earnings okx.JSONFloat64 `json:"earnings"`
	}
	Detail struct {
		Ccy    string          `json:"ccy"`
		Amt    okx.JSONFloat64 `json:"amt"`
		CnvAmt okx.JSONFloat64 `json:"cnvAmt"`
		Fee    okx.JSONFloat64 `json:"fee"`
	}
	SmallAssetConvert struct {
		TotalCnvAmt string    `json:"totalCnvAmt"`
		Details     []*Detail `json:"details"`
	}
)
