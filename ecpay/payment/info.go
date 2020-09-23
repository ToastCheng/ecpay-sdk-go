package payment

import (
	"encoding/json"
	"net/url"
)

// Info defines the struct of trade info.
type Info struct {
	MerchantTradeNo string `json:"MerchantTradeNo,omitempty"`
	TimeStamp       string `json:"TimeStamp,omitempty"`
	PlatformID      string `json:"PlatformID,omitempty"`
}

// Validate validate if the trade struct is valid.
func (t Info) Validate() (bool, error) {
	return true, nil
}

// ToFormData transform the Trade struct to url.Values
func (t Info) ToFormData() url.Values {
	req := url.Values{}
	mp := map[string]string{}
	databytes, _ := json.Marshal(t)
	json.Unmarshal(databytes, &mp)
	for k, v := range mp {
		req.Set(k, v)
	}

	return req
}
