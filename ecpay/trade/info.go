package trade

import (
	"encoding/json"
	"net/url"

	"github.com/toastcheng/ecpay-sdk-go/ecpay/utils"
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
func (t Info) ToFormData(merchantID, hashKey, hashIV string) url.Values {
	req := url.Values{}
	mp := map[string]string{}
	databytes, _ := json.Marshal(t)
	json.Unmarshal(databytes, &mp)
	for k, v := range mp {
		req.Set(k, v)
	}
	req.Set("MerchantID", merchantID)
	req.Set("CheckMacValue", utils.GetCheckMacValue(req, hashKey, hashIV))

	return req
}
