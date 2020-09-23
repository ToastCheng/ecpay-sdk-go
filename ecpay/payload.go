package ecpay

import (
	"net/url"
)

// Payload defines the interface of payload for communicate with ECPay server.
type Payload interface {
	Validate() (bool, error)
	ToFormData(merchantID, hashKey, hashIV string) url.Values
}
