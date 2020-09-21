package payment

// ChoosePaymentType defines the field for ChoosePayment
type ChoosePaymentType string

const (
	// All all type
	All ChoosePaymentType = "ALL"
	// Credit credit card
	Credit ChoosePaymentType = "Credit"
	// WebATM web ATM
	WebATM ChoosePaymentType = "WebATM"
	// ATM ATM
	ATM ChoosePaymentType = "ATM"
	// CVS convenience store voucher code
	CVS ChoosePaymentType = "CVS"
	// BarCode bar code
	BarCode ChoosePaymentType = "BARCODE"
	// GooglePay Google Pay
	GooglePay ChoosePaymentType = "GooglePay"
)
