package order

type CarrierType string

const (
	CarrierTypeNone      CarrierType = ""
	CarrierTypeMember    CarrierType = "1"
	CarrierTypeCitizen   CarrierType = "2"
	CarrierTypeCellphone CarrierType = "3"
)

// DonationType defines the struct of donation.
type DonationType string

const (
	DonationTypeYes DonationType = "1"
	DonationTypeNo  DonationType = "2"
)

// InvoiceType defines the struct of invoice.
type InvoiceType string

const (
	// InvoiceTypeGeneral 一般稅額
	InvoiceTypeGeneral InvoiceType = "07"
	// InvoiceTypeSpecial 特種稅額
	InvoiceTypeSpecial InvoiceType = "08"
)

type PeriodType string

const (
	PeriodTypeYear  PeriodType = "Y"
	PeriodTypeMonth PeriodType = "M"
	PeriodTypeDay   PeriodType = "D"
)

// TaxType tax type.
type TaxType string

const (
	// TaxTypeDutiable 應稅
	TaxTypeDutiable TaxType = "1"
	// TaxTypeZero 零稅率
	TaxTypeZero TaxType = "2"
	// TaxTypeFree 免稅
	TaxTypeFree TaxType = "3"
	// TaxTypeMix 若為混合應稅與免稅或零稅率時(限收 銀機發票無法分辨時使用，且需通過申 請核可)
	TaxTypeMix TaxType = "9"
)

type UnionPayType string

const (
	UnionPayTypeChooseInWebPage UnionPayType = "0"
	UnionPayTypeYes             UnionPayType = "1"
	UnionPayTypeNo              UnionPayType = "2"
)

type PaymentType string

const (
	PaymentTypeAIO              PaymentType = "aio"
	PaymentTypeWebATMTaishin    PaymentType = "WebATM_TAISHIN"
	PaymentTypeWebATMBOT        PaymentType = "WebATM_BOT"
	PaymentTypeWebATMChinaTrust PaymentType = "WebATM_CHINATRUST"
	PaymentTypeWebATMCathay     PaymentType = "WebATM_CATHAY"
	PaymentTypeWebATMLand       PaymentType = "WebATM_LAND"
	PaymentTypeWebATMSinoPac    PaymentType = "WebATM_SINOPAC"
	PaymentTypeATMESUN          PaymentType = "ATM_ESUN"
	PaymentTypeATMFubon         PaymentType = "ATM_FUBON"
	PaymentTypeATMFirst         PaymentType = "ATM_FIRST"
	PaymentTypeATMCathay        PaymentType = "ATM_CATHAY"
	PaymentTypeCVSCVS           PaymentType = "CVS_CVS"
	PaymentTypeCVSFamily        PaymentType = "CVS_FAMILY"
	PaymentTypeCVSIBon          PaymentType = "CVS_IBON"
	PaymentTypeCreditCreditCard PaymentType = "Credit_CreditCard"
)

type ChoosePaymentType string

const (
	// ChoosePaymentTypeAll all type
	ChoosePaymentTypeAll ChoosePaymentType = "ALL"
	// ChoosePaymentTypeATM credit card
	ChoosePaymentTypeCredit ChoosePaymentType = "Credit"
	// ChoosePaymentTypeATM web ATM
	ChoosePaymentTypeWebATM ChoosePaymentType = "WebATM"
	// ChoosePaymentTypeATM ATM
	ChoosePaymentTypeATM ChoosePaymentType = "ATM"
	// ChoosePaymentTypeCVS convenience store voucher code
	ChoosePaymentTypeCVS ChoosePaymentType = "CVS"
	// ChoosePaymentTypeBarCode bar code
	ChoosePaymentTypeBarCode ChoosePaymentType = "BARCODE"
	// ChoosePaymentTypeGooglePay Google Pay
	ChoosePaymentTypeGooglePay ChoosePaymentType = "GooglePay"
)

type ChooseSubpaymentType string

const (
	ChooseSubpaymentTypeTaishin    ChooseSubpaymentType = "TAISHIN"
	ChooseSubpaymentTypeESUN       ChooseSubpaymentType = "ESUN"
	ChooseSubpaymentTypeBOT        ChooseSubpaymentType = "BOT"
	ChooseSubpaymentTypeFubon      ChooseSubpaymentType = "FUBON"
	ChooseSubpaymentTypeChinaTrust ChooseSubpaymentType = "CHINATRUST"
	ChooseSubpaymentTypeFirst      ChooseSubpaymentType = "FIRST"
	ChooseSubpaymentTypeCathay     ChooseSubpaymentType = "CATHAY"
	ChooseSubpaymentTypeMega       ChooseSubpaymentType = "MEGA"
	ChooseSubpaymentTypeLand       ChooseSubpaymentType = "LAND"
	ChooseSubpaymentTypeTachong    ChooseSubpaymentType = "TACHONG"
	ChooseSubpaymentTypeSinoPac    ChooseSubpaymentType = "SINOPAC"
	ChooseSubpaymentTypeCVS        ChooseSubpaymentType = "CVS"
	ChooseSubpaymentTypeOK         ChooseSubpaymentType = "OK"
	ChooseSubpaymentTypeFamily     ChooseSubpaymentType = "FAMILY"
	ChooseSubpaymentTypeHiLife     ChooseSubpaymentType = "HILIFE"
	ChooseSubpaymentTypeIBon       ChooseSubpaymentType = "IBON"
	ChooseSubpaymentTypeBarcode    ChooseSubpaymentType = "BARCODE"
)
