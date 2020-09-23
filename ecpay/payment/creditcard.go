package payment

import (
	"encoding/json"
	"net/url"
)

// CreditCardPeriodInfo defines the struct of trade info.
type CreditCardPeriodInfo struct {
	MerchantTradeNo string `json:"MerchantTradeNo,omitempty"`
	TimeStamp       string `json:"TimeStamp,omitempty"`
	PlatformID      string `json:"PlatformID,omitempty"`
}

// Validate validate if the trade struct is valid.
func (c CreditCardPeriodInfo) Validate() (bool, error) {
	return true, nil
}

// ToFormData transform the Trade struct to url.Values
func (c CreditCardPeriodInfo) ToFormData() url.Values {
	req := url.Values{}
	mp := map[string]string{}
	databytes, _ := json.Marshal(c)
	json.Unmarshal(databytes, &mp)
	for k, v := range mp {
		req.Set(k, v)
	}

	return req
}
