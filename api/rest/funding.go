package rest

import (
	"encoding/json"
	"github.com/liuhengloveyou/okx-go"
	requests "github.com/liuhengloveyou/okx-go/requests/rest/funding"
	responses "github.com/liuhengloveyou/okx-go/responses/funding"
	"net/http"
	"strings"
)

// Funding
//
// https://www.okx.com/docs-v5/en/#rest-api-funding
type Funding struct {
	client *ClientRest
}

// NewFunding returns a pointer to a fresh Funding
func NewFunding(c *ClientRest) *Funding {
	return &Funding{c}
}

// GetCurrencies
// Retrieve a list of all currencies. Not all currencies can be traded. Currencies that have not been defined in ISO 4217 may use a custom symbol.
//
// https://www.okx.com/docs-v5/en/#rest-api-funding-get-currencies
func (c *Funding) GetCurrencies() (response responses.GetCurrencies, err error) {
	p := "/api/v5/asset/currencies"

	res, err := c.client.Do(http.MethodGet, p, true)
	if err != nil {
		return
	}
	defer res.Body.Close()

	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetBalance
// Retrieve the balances of all the assets, and the amount that is available or on hold.
//
// https://www.okx.com/docs-v5/en/#rest-api-funding-get-balance
func (c *Funding) GetBalance(req requests.GetBalance) (response responses.GetBalance, err error) {
	p := "/api/v5/asset/balances"
	m := okx.S2M(req)
	if len(req.Ccy) > 0 {
		m["ccy"] = strings.Join(req.Ccy, ",")
	}
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// FundsTransfer
// This endpoint supports the transfer of funds between your funding account and trading account, and from the master account to sub-accounts. Direct transfers between sub-accounts are not allowed.
//
// https://www.okx.com/docs-v5/en/#rest-api-funding-funds-transfer
func (c *Funding) FundsTransfer(req requests.FundsTransfer) (response responses.FundsTransfer, err error) {
	p := "/api/v5/asset/transfer"
	m := okx.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// FundsTransferState
//
// https://www.okx.com/docs-v5/en/#funding-account-rest-api-get-funds-transfer-state
func (c *Funding) FundsTransferState(req requests.FundsTransferState) (response responses.FundsTransferState, err error) {
	p := "/api/v5/asset/transfer-state"
	m := okx.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// AssetBillsDetails
// Query the billing record, you can get the latest 1 month historical data.
//
// https://www.okx.com/docs-v5/en/#rest-api-funding-asset-bills-details
func (c *Funding) AssetBillsDetails(req requests.AssetBillsDetails) (response responses.AssetBillsDetails, err error) {
	p := "/api/v5/asset/bills"
	m := okx.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetDepositAddress
// Retrieve the deposit addresses of currencies, including previously-used addresses.
//
// https://www.okx.com/docs-v5/en/#rest-api-funding-get-deposit-address
func (c *Funding) GetDepositAddress(req requests.GetDepositAddress) (response responses.GetDepositAddress, err error) {
	p := "/api/v5/asset/deposit-address"
	m := okx.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetDepositHistory
// Retrieve the deposit history of all currencies, up to 100 recent records in a year.
//
// https://www.okx.com/docs-v5/en/#rest-api-funding-get-deposit-history
func (c *Funding) GetDepositHistory(req requests.GetDepositHistory) (response responses.GetDepositHistory, err error) {
	p := "/api/v5/asset/deposit-history"
	m := okx.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// Withdrawal
// Withdrawal of tokens.
//
// https://www.okx.com/docs-v5/en/#rest-api-funding-withdrawal
func (c *Funding) Withdrawal(req requests.Withdrawal) (response responses.Withdrawal, err error) {
	p := "/api/v5/asset/withdrawal"
	m := okx.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetWithdrawalHistory
// Retrieve the withdrawal records according to the currency, withdrawal status, and time range in reverse chronological order. The 100 most recent records are returned by default.
//
// https://www.okx.com/docs-v5/en/#rest-api-funding-get-withdrawal-history
func (c *Funding) GetWithdrawalHistory(req requests.GetWithdrawalHistory) (response responses.GetWithdrawalHistory, err error) {
	p := "/api/v5/asset/withdrawal-history"
	m := okx.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// PiggyBankPurchaseRedemption
//
// https://www.okx.com/docs-v5/en/#rest-api-funding-piggybank-purchase-redemption
func (c *Funding) PiggyBankPurchaseRedemption(req requests.PiggyBankPurchaseRedemption) (response responses.PiggyBankPurchaseRedemption, err error) {
	p := "/api/v5/asset/purchase_redempt"
	m := okx.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetPiggyBankBalance
//
// https://www.okx.com/docs-v5/en/#rest-api-funding-get-piggybank-balance
func (c *Funding) GetPiggyBankBalance(req requests.GetPiggyBankBalance) (response responses.GetPiggyBankBalance, err error) {
	p := "/api/v5/asset/piggy-balance"
	m := okx.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// SmallAssetConvert
//
// https://www.okx.com/docs-v5/en/#funding-account-rest-api-small-assets-convert
func (c *Funding) SmallAssetConvert(req requests.SmallAssetConvert) (response responses.SmallAssetConvert, err error) {
	p := "/api/v5/asset/convert-dust-assets"
	res, err := c.client.DoBatch(p, req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}
