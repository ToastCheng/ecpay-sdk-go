package payment

// ActionType defines the struct of credit card action.
type ActionType string

const (
	// ActionTypeC (關帳).
	ActionTypeC ActionType = "C"
	// ActionTypeR (退刷).
	ActionTypeR ActionType = "R"
	// ActionTypeE (取消).
	ActionTypeE ActionType = "E"
	// ActionTypeN (放棄).
	ActionTypeN ActionType = "N"
)
