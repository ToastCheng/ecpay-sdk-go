package payment

type ChoosePaymentType string

const (
	ALL        ChoosePaymentType = "ALL"
	CREDIT     ChoosePaymentType = "Credit"
	WEB_ATM    ChoosePaymentType = "WebATM"
	ATM        ChoosePaymentType = "ATM"
	CVS        ChoosePaymentType = "CVS"
	BARCODE    ChoosePaymentType = "BARCODE"
	GOOGLE_PAY ChoosePaymentType = "GooglePay"
)
