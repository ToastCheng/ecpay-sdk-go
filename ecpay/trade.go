package ecpay

import (
	"encoding/json"
	"fmt"
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
	for k, v := range mp {
		switch t := v.(type) {
		case int:
			req.Set(k, string(t))
		case int64:
			req.Set(k, string(t))
		case float32, float64:
			req.Set(k, fmt.Sprintf("%.0f", t))
		case string:
			req.Set(k, t)
		}
	}
	return req
}
