package ecpay

import "time"

type Order struct {
	MerchantTradeNo   string
	StoreID           string
	MerchantTradeDate time.Time
	PaymentType       string
	TotalAmount       int
	TradeDesc         string
	ItemName          string
	ReturnURL         string
	ChoosePayment     string
	ClientBackURL     string
	ItemURL           string
	Remark            string
	ChooseSubPayment  string
	OrderResultURL    string
	NeedExtraPaidInfo bool
	DeviceSource      string
	IgnorePayment     string
	PlatformID        string
	InvoiceMark       bool
	CustomField1      string
	CustomField2      string
	CustomField3      string
	CustomField4      string
	EncryptType       string
}
