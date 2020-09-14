package api

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var NewOrderUrl string = "https://payment-stage.ecpay.com.tw/Cashier/AioCheckOut/V5"
var QueryOrderUrl string = "https://payment-stage.ecpay.com.tw/Cashier/QueryTradeInfo/V5"

// V5.1.40

// ECPayOrderRequest defines the payload of ECPayOrderRequest
type ECPayOrderRequest struct {
	MerchantID        string // 合作特電編號
	MerchantTradeNo   string // 特店交易編號
	MerchantTradeDate string // 交易時間 yyyy/MM/dd HH:mm:ss
	PaymentType       string // 交易類型，固定填aio
	TotalAmount       int    // 交易金額
	TradeDesc         string // 交易描述（需要urlencode）
	ItemName          string // 商品名稱，多項需用#隔開
	ReturnURL         string // 付款完成通知回傳網址
	ChoosePayment     string // 選擇預設付款方式（Credit）
	CheckMacValue     string // 檢查碼
	ClientBackURL     string // Client端返回特店的按鈕連結
	OrderResultURL    string // Client端回傳付款結果網址
	Redeem            string // 是否使用紅利折抵[Y|N]
	InvoiceMark       string // 電子發票開立註記[Y|N]
	// 分期付款
	CreditInstallment string `json:",omitempty"` // 刷卡分期期數[3|6|12|18|24]，多選用,隔開
	// 定期定額
	PeriodAmount    int    `json:",omitempty"` // 每次授權金額
	PeriodType      string `json:",omitempty"` // 週期種類[D|M|Y]
	Frequency       int    `json:",omitempty"` // 執行頻率
	ExecTimes       int    `json:",omitempty"` // 執行次數
	PeriodReturnURL string `json:",omitempty"` // 定期定額的執行結果回應URL

	// Invoice
	RelateNumber       string // 特店自訂編號
	CustomerID         string `json:",omitempty"`
	CustomerIdentifier string `json:",omitempty"`
	CustomerName       string `json:",omitempty"`
	CustomerAddr       string `json:",omitempty"`
	CustomerPhone      string `json:",omitempty"`
	CustomerEmail      string `json:",omitempty"`
	ClearanceMark      string `json:",omitempty"`
	TaxType            string `json:",omitempty"`
	CarruerType        string `json:",omitempty"`
	CarruerNum         string `json:",omitempty"`
	Donation           string `json:",omitempty"`
	LoveCode           string `json:",omitempty"`
	Print              string `json:",omitempty"`
	InvoiceItemName    string `json:",omitempty"`
	InvoiceItemCount   string `json:",omitempty"`
	InvoiceItemWord    string `json:",omitempty"`
	InvoiceItemPrice   string `json:",omitempty"`
	InvoiceItemTaxType string `json:",omitempty"`
	InvoiceRemark      string `json:",omitempty"`
	DelayDay           string `json:",omitempty"`
	InvType            string `json:",omitempty"`
}

// NewECPayOrderFormData creates an NewECPayOrderFormData
func NewECPayOrderFormData(req OrderRequest) url.Values {
	ecpayReq := map[string][]string{}
	ecpayReq["ChoosePayment"] = []string{req.ChoosePayment}
	ecpayReq["EncryptType"] = []string{"1"}
	// ecpayReq["OrderResultURL"] = []string{""}
	// ecpayReq["MerchantID"] = []string{""}
	ecpayReq["MerchantID"] = []string{req.MerchantID}
	ecpayReq["MerchantTradeNo"] = []string{
		fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()),
	}
	ecpayReq["MerchantTradeDate"] = []string{
		time.Now().Format("2006/01/02 15:04:05"),
	}
	ecpayReq["PaymentType"] = []string{"aio"}
	ecpayReq["TotalAmount"] = []string{strconv.Itoa(req.TotalAmount)}
	ecpayReq["TradeDesc"] = []string{req.TradeDesc}
	ecpayReq["ItemName"] = []string{req.ItemName}
	ecpayReq["ReturnURL"] = []string{req.ReturnURL}
	ecpayReq["CheckMacValue"] = []string{
		GetCheckMacValue(ecpayReq),
	}
	log.Print(ecpayReq)
	return ecpayReq
}

// ResultCallback defines the payload of ResultCallback
type ResultCallback struct {
	MerchantID           string // 特店編號
	MerchantTradeNo      string // 特店交易編號
	StoreID              string // 特店旗下店舖代號
	RtnCode              int    // 交易狀態 [1:success]
	RtnMsg               string // 交易訊息
	TradeNo              string // 綠界交易編號
	TradeAmt             int    // 交易金額
	PaymentDate          string // 付款時間
	PaymentType          string // 特店選擇的付款方式
	PaymentTypeChargeFee int    // 交易手續費
	TradeDate            string // 訂單成立時間
	SimulatePaid         int    // 是否為模擬付款
}

// QueryOrderRequest defines the payload of QueryOrderRequest
type QueryOrderRequest struct {
	MerchantID      string // 特店編號
	MerchantTradeNo string // 特店交易編號
	TimeStamp       int    // 驗證時間 UnixTimeStamp
	PlatformID      string // 平台商代號
	CheckMacValue   string // 檢查碼
}
