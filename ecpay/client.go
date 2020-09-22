package ecpay

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/toastcheng/ecpay-sdk-go/ecpay/order"
	"github.com/toastcheng/ecpay-sdk-go/ecpay/utils"
)

// Client implements client for making ECPay api.
type Client struct {
	merchantID string
	hashKey    string
	hashIV     string
	endpoint   string
	vendor     string
	httpClient *http.Client
	options    ClientOption
}

// ClientOption extra option to configure the ECPay client.
type ClientOption func(*Client) error

// WithSandbox configures the client, making it sending request to sandbox environment.
func WithSandbox(client *Client) error {
	client.endpoint = "https://payment-stage.ecpay.com.tw"
	client.vendor = "https://vendor-stage.ecpay.com.tw"
	return nil
}

// NewClient create a client for communicating to ECPay server.
func NewClient(merchantID, hashKey, hashIV string, options ...ClientOption) (*Client, error) {
	c := &Client{
		merchantID: merchantID,
		hashKey:    hashKey,
		hashIV:     hashIV,
		endpoint:   "https://payment.ecpay.com.tw",
		vendor:     "https://vendor.ecpay.com.tw",
		httpClient: &http.Client{},
	}
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

// AioCheckOut sends an order to ECPay server.
func (c *Client) AioCheckOut(order order.Order) (string, error) {

	if ok, err := order.Validate(); !ok {
		return "", err
	}

	r := &Request{endpoint: c.endpoint + "/Cashier/AioCheckOut/V5", httpClient: c.httpClient}
	res, _ := r.Do(order)

	return res, nil
}

// QueryTradeInfo 查詢訂單
func (c *Client) QueryTradeInfo() {

	form := url.Values{}
	form["MerchantID"] = []string{"1234567890"}
	form["MerchantTradeNo"] = []string{"12345678901234567890"}
	form["TimeStamp"] = []string{"1234567890"}
	form["PlatformID"] = []string{"1234567890"}

	form["CheckMacValue"] = []string{utils.GetCheckMacValue(form)}
	formStr := form.Encode()
	formStr = strings.ReplaceAll(formStr, "-", "%2d")
	formStr = strings.ReplaceAll(formStr, "_", "%5f")
	formStr = strings.ReplaceAll(formStr, "*", "%2a")
	formStr = strings.ReplaceAll(formStr, "(", "%28")
	formStr = strings.ReplaceAll(formStr, ")", "%29")
	formStr = strings.ReplaceAll(formStr, "+", "%20")
	ecpayReq, _ := http.NewRequest("POST", c.endpoint+"/Cashier/QueryTradeInfo/V5", strings.NewReader(formStr))
	ecpayReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

}

// QueryCreditCardPeriodInfo 信用卡定期定額訂單查詢
func (c *Client) QueryCreditCardPeriodInfo() {
	ecpayReq, _ := http.NewRequest("POST", c.endpoint+"/Cashier/QueryCreditCardPeriodInfo", nil)
	ecpayReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

}

// QueryTrade 查詢信用卡單筆明細記錄
func (c *Client) QueryTrade() {
	ecpayReq, _ := http.NewRequest("POST", c.endpoint+"/CreditDetail/QueryTrade/V2", nil)
	ecpayReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

}

// QueryPaymentInfo 查詢 ATM/CVS/BARCODE 取號結果
func (c *Client) QueryPaymentInfo() {
	ecpayReq, _ := http.NewRequest("POST", c.endpoint+"/Cashier/QueryPaymentInfo", nil)
	ecpayReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

}

// QueryPaymentInfo 信用卡請退款功能 (若不撰寫此 API，則可透過廠商後台功能處理)
func (c *Client) DoAction() {
	ecpayReq, _ := http.NewRequest("POST", c.endpoint+"/CreditDetail/DoAction", nil)
	ecpayReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

}

// TradeNoAio 下載特店對帳媒體檔
func (c *Client) TradeNoAio() {
	ecpayReq, _ := http.NewRequest("POST", c.vendor+"/PaymentMedia/TradeNoAio", nil)
	ecpayReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

}

// FundingReconDetail 下載信用卡撥款對帳資料檔
func (c *Client) FundingReconDetail() {
	ecpayReq, _ := http.NewRequest("POST", c.vendor+"/CreditDetail/FundingReconDetail", nil)
	ecpayReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

}
