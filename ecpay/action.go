package ecpay

import (
	"encoding/json"
	"fmt"
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
