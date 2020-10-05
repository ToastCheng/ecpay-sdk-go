package ecpay

import (
	"encoding/json"
	"net/url"
)

// CreditCardAction defines the struct of trade info.
type CreditCardAction struct {
	MerchantTradeNo string     `json:"MerchantTradeNo,omitempty"`
	TradeNo         string     `json:"TradeNo,omitempty"`
	Action          ActionType `json:"Action,omitempty"`
	PlatformID      string     `json:"PlatformID,omitempty"`
	TotalAmount     int        `json:"TotalAmount,omitempty"`
}

// Validate validate if the trade struct is valid.
func (a CreditCardAction) Validate() (bool, error) {
	return true, nil
}

// ToFormData transform the Trade struct to url.Values
func (a CreditCardAction) ToFormData() url.Values {
	req := url.Values{}
	mp := map[string]interface{}{}
	databytes, _ := json.Marshal(a)
	json.Unmarshal(databytes, &mp)
	req = setUrlValues(req, mp)

	return req
}
