package ecpay

import (
	"encoding/json"
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
	mp := map[string]string{}
	databytes, _ := json.Marshal(t)
	json.Unmarshal(databytes, &mp)
	for k, v := range mp {
		req.Set(k, v)
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
	mp := map[string]string{}
	databytes, _ := json.Marshal(t)
	json.Unmarshal(databytes, &mp)
	for k, v := range mp {
		req.Set(k, v)
	}

	return req
}
