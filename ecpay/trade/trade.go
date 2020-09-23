package trade

import (
	"net/url"

	"github.com/toastcheng/ecpay-sdk-go/ecpay/utils"
)

type Trade struct {
}

// Validate validate if the trade struct is valid.
func (o Trade) Validate() (bool, error) {
	return true, nil
}

// ToFormData transform the Trade struct to url.Values
func (o Trade) ToFormData(merchantID, hashKey, hashIV string) url.Values {
	ecpayReq := map[string][]string{}
	ecpayReq["CheckMacValue"] = []string{utils.GetCheckMacValue(ecpayReq, hashKey, hashIV)}

	return ecpayReq
}
