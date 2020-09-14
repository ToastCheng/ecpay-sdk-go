package api

// OrderRequest defines the payload of OrderRequest
type OrderRequest struct {
	MerchantID    string `json:"merchant_id" schema:"merchant_id"`       // 合作特電編號
	TotalAmount   int    `json:"total_amount" schema:"total_amount"`     // 交易金額
	TradeDesc     string `json:"trade_desc" schema:"trade_desc"`         // 交易描述（需要urlencode）
	ItemName      string `json:"item_name" schema:"item_name"`           // 商品名稱，多項需用#隔開
	ReturnURL     string `json:"return_url" schema:"return_url"`         // 付款完成通知回傳網址
	ChoosePayment string `json:"choose_payment" schema:"choose_payment"` // 選擇預設付款方式（Credit）
}

// OrderResponse defines the payload of OrderResponse
type OrderResponse struct {
	Message string `json:"message"`
}
