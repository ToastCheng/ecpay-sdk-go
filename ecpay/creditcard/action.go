package creditcard

import (
	"encoding/json"
	"net/url"

	"github.com/toastcheng/ecpay-sdk-go/ecpay/utils"
)

// Action defines the struct of trade info.
type Action struct {
	MerchantTradeNo string     `json:"MerchantTradeNo,omitempty"`
	TradeNo         string     `json:"TradeNo,omitempty"`
	Action          ActionType `json:"Action,omitempty"`
	PlatformID      string     `json:"PlatformID,omitempty"`
	TotalAmount     string     `json:"TotalAmount,omitempty"`
}

// Validate validate if the trade struct is valid.
func (a Action) Validate() (bool, error) {
	return true, nil
}

// ToFormData transform the Trade struct to url.Values
func (a Action) ToFormData(merchantID, hashKey, hashIV string) url.Values {
	req := url.Values{}
	mp := map[string]string{}
	databytes, _ := json.Marshal(a)
	json.Unmarshal(databytes, &mp)
	for k, v := range mp {
		req.Set(k, v)
	}
	req.Set("MerchantID", merchantID)
	req.Set("CheckMacValue", utils.GetCheckMacValue(req, hashKey, hashIV))

	return req
}
