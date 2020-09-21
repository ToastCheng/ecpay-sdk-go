package payment

type PaymentType string

const (
	AIO               PaymentType = "aio"
	WebATM_TAISHIN    PaymentType = "WebATM_TAISHIN"
	WebATM_BOT        PaymentType = "WebATM_BOT"
	WebATM_CHINATRUST PaymentType = "WebATM_CHINATRUST"
	WebATM_CATHAY     PaymentType = "WebATM_CATHAY"
	WebATM_LAND       PaymentType = "WebATM_LAND"
	WebATM_SINOPAC    PaymentType = "WebATM_SINOPAC"
	ATM_ESUN          PaymentType = "ATM_ESUN"
	ATM_FUBON         PaymentType = "ATM_FUBON"
	ATM_FIRST         PaymentType = "ATM_FIRST"
	ATM_CATHAY        PaymentType = "ATM_CATHAY"
	CVS_CVS           PaymentType = "CVS_CVS"
	CVS_FAMILY        PaymentType = "CVS_FAMILY"
	CVS_IBON          PaymentType = "CVS_IBON"
	Credit_CreditCard PaymentType = "Credit_CreditCard"
)
