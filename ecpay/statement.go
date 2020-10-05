package ecpay

import (
	"encoding/json"
	"net/url"
)

// Statement defines the struct of trade info.
type Statement struct {
	MerchantTradeNo string      `json:"MerchantTradeNo,omitempty"`
	PayDateType     PayDateType `json:"PayDateType,omitempty"`
	BeginDate       string      `json:"BeginDate,omitempty"`
	EndDate         string      `json:"EndDate,omitempty"`
	PaymentType     string      `json:"PaymentType,omitempty"`
	PlatformStatus  string      `json:"PlatformStatus,omitempty"`
	PaymentStatus   string      `json:"PaymentStatus,omitempty"`
	AllocateStatus  string      `json:"AllocateStatus,omitempty"`
	MediaFormated   string      `json:"MediaFormated,omitempty"`
	CharSet         string      `json:"CharSet,omitempty"`
}

// Validate validate if the trade struct is valid.
func (s Statement) Validate() (bool, error) {
	// TODO: handle error
	return true, nil
}

// ToFormData transform the Trade struct to url.Values
func (s Statement) ToFormData() url.Values {
	req := url.Values{}
	mp := map[string]interface{}{}
	databytes, _ := json.Marshal(s)
	json.Unmarshal(databytes, &mp)
	req = setUrlValues(req, mp)

	return req
}

// CreditCardStatement defines the struct of trade info.
type CreditCardStatement struct {
	MerchantTradeNo string              `json:"MerchantTradeNo,omitempty"`
	DateType        DateType            `json:"DateType,omitempty"`
	BeginDate       string              `json:"BeginDate,omitempty"`
	EndDate         string              `json:"EndDate,omitempty"`
	PaymentType     MerchantPaymentType `json:"PaymentType,omitempty"`
	PlatformStatus  PlatformStatusType  `json:"PlatformStatus,omitempty"`
	PaymentStatus   PaymentStatusType   `json:"PaymentStatus,omitempty"`
	AllocateStatus  AllocateStatusType  `json:"AllocateStatus,omitempty"`
	MediaFormated   MediaFormatedType   `json:"MediaFormated,omitempty"`
	CharSet         CharSetType         `json:"CharSet,omitempty"`
}

// Validate validate if the trade struct is valid.
func (s CreditCardStatement) Validate() (bool, error) {
	// TODO: handle error
	return true, nil
}

// ToFormData transform the Trade struct to url.Values
func (s CreditCardStatement) ToFormData() url.Values {
	req := url.Values{}
	mp := map[string]interface{}{}
	databytes, _ := json.Marshal(s)
	json.Unmarshal(databytes, &mp)
	req = setUrlValues(req, mp)

	return req
}
