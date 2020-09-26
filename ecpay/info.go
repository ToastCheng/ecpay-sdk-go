package ecpay

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// TradeInfo defines the struct of trade info.
type TradeInfo struct {
	MerchantTradeNo string `json:"MerchantTradeNo,omitempty"`
	TimeStamp       string `json:"TimeStamp,omitempty"`
	PlatformID      string `json:"PlatformID,omitempty"`
}

// Validate validate if the trade struct is valid.
func (t TradeInfo) Validate() (bool, error) {
	// TODO: handle error
	return true, nil
}

// ToFormData transform the Trade struct to url.Values
func (t TradeInfo) ToFormData() url.Values {
	req := url.Values{}
	mp := map[string]interface{}{}
	databytes, _ := json.Marshal(t)
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

// PaymentInfo defines the struct of trade info.
type PaymentInfo struct {
	MerchantTradeNo string `json:"MerchantTradeNo,omitempty"`
	TimeStamp       int64  `json:"TimeStamp,omitempty"`
	PlatformID      string `json:"PlatformID,omitempty"`
}

// Validate validate if the trade struct is valid.
func (t PaymentInfo) Validate() (bool, error) {
	// TODO: handle error
	return true, nil
}

// ToFormData transform the Trade struct to url.Values
func (t PaymentInfo) ToFormData() url.Values {
	req := url.Values{}
	mp := map[string]interface{}{}
	databytes, _ := json.Marshal(t)
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

// CreditCardPeriodInfo defines the struct of trade info.
type CreditCardPeriodInfo struct {
	MerchantTradeNo string `json:"MerchantTradeNo,omitempty"`
	TimeStamp       int64  `json:"TimeStamp,omitempty"`
	PlatformID      string `json:"PlatformID,omitempty"`
}

// Validate validate if the trade struct is valid.
func (c CreditCardPeriodInfo) Validate() (bool, error) {
	// TODO: handle error
	return true, nil
}

// ToFormData transform the Trade struct to url.Values
func (c CreditCardPeriodInfo) ToFormData() url.Values {
	req := url.Values{}
	mp := map[string]interface{}{}
	databytes, _ := json.Marshal(c)
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
