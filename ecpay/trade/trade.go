package trade

import (
	"encoding/json"
	"net/url"

	"github.com/toastcheng/ecpay-sdk-go/ecpay/utils"
)

// Trade defines the struct of trade.
type Trade struct {
	MerchantTradeNo string
	TimeStamp       string
}

// Validate validate if the trade struct is valid.
func (o Trade) Validate() (bool, error) {
	return true, nil
}

// ToFormData transform the Trade struct to url.Values
func (o Trade) ToFormData(merchantID, hashKey, hashIV string) url.Values {
	req := url.Values{}
	mp := map[string]string{}
	databytes, _ := json.Marshal(o)
	json.Unmarshal(databytes, &mp)
	for k, v := range mp {
		req.Set(k, v)
	}
	req.Set("MerchantID", merchantID)
	req["CheckMacValue"] = []string{utils.GetCheckMacValue(req, hashKey, hashIV)}

	return req
}
