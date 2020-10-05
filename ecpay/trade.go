package ecpay

import (
	"encoding/json"
	"net/url"
)

// Trade defines the struct of trade.
type Trade struct {
	CreditRefundID  string `json:"CreditRefundId,omitempty"`
	CreditAmount    int    `json:"CreditAmount,omitempty"`
	CreditCheckCode string `json:"CreditCheckCode,omitempty"`
}

// Validate validate if the trade struct is valid.
func (o Trade) Validate() (bool, error) {
	// TODO: handle error
	return true, nil
}

// ToFormData transform the Trade struct to url.Values
func (o Trade) ToFormData() url.Values {
	req := url.Values{}
	mp := map[string]interface{}{}
	databytes, _ := json.Marshal(o)
	json.Unmarshal(databytes, &mp)
	req = setUrlValues(req, mp)

	return req
}
