package creditcard

import (
	"encoding/json"
	"net/url"

	"github.com/toastcheng/ecpay-sdk-go/ecpay/utils"
)

// Statement defines the struct of trade info.
type Statement struct {
	MerchantTradeNo string `json:"MerchantTradeNo,omitempty"`
	DateType        string `json:"TimeStamp,omitempty"`
	BeginDate       string `json:"BeginDate,omitempty"`
	EndDate         string `json:"EndDate,omitempty"`
	PaymentType     string `json:"PaymentType,omitempty"`
	PlatformStatus  string `json:"PlatformStatus,omitempty"`
	PaymentStatus   string `json:"PaymentStatus,omitempty"`
	AllocateStatus  string `json:"AllocateStatus,omitempty"`
	MediaFormated   string `json:"MediaFormated,omitempty"`
	CharSet         string `json:"CharSet,omitempty"`
}

// Validate validate if the trade struct is valid.
func (s Statement) Validate() (bool, error) {
	return true, nil
}

// ToFormData transform the Trade struct to url.Values
func (s Statement) ToFormData(merchantID, hashKey, hashIV string) url.Values {
	req := url.Values{}
	mp := map[string]string{}
	databytes, _ := json.Marshal(s)
	json.Unmarshal(databytes, &mp)
	for k, v := range mp {
		req.Set(k, v)
	}
	req.Set("MerchantID", merchantID)
	req.Set("CheckMacValue", utils.GetCheckMacValue(req, hashKey, hashIV))

	return req
}
