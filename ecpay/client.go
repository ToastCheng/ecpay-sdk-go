package ecpay

import (
	"ecpay/api"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

// ClientOption .
type ClientOption func(*Client) error

// WithSandbox .
func WithSandbox(client *Client) error {
	client.endpoint = "https://payment-stage.ecpay.com.tw"
	client.vendor = "https://vendor-stage.ecpay.com.tw"
	return nil
}

// NewClient .
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

// AioCheckOut .
func (c *Client) CreateOrder() ([]byte, error) {

	form := api.NewECPayOrderFormData(api.OrderRequest{})
	formStr := form.Encode()
	formStr = strings.ReplaceAll(formStr, "-", "%2d")
	formStr = strings.ReplaceAll(formStr, "_", "%5f")
	formStr = strings.ReplaceAll(formStr, "*", "%2a")
	formStr = strings.ReplaceAll(formStr, "(", "%28")
	formStr = strings.ReplaceAll(formStr, ")", "%29")
	formStr = strings.ReplaceAll(formStr, "+", "%20")

	ecpayReq, _ := http.NewRequest("POST", c.endpoint+"/Cashier/AioCheckOut/V5", strings.NewReader(formStr))

	ecpayReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	ecpayReq.Header.Add("accept", "*/*")
	ecpayReq.Header.Add("accept-encoding", "gzip, deflate, br")
	ecpayReq.Header.Add("user-agent", "PostmanRuntime/7.24.0")
	ecpayReq.Header.Add("cache-control", "no-cache")
	ecpayReq.Header.Add("postman-token", "40975c86-58b2-4b7e-aee6-9add8ed2cc09")

	ecpayResp, err := c.httpClient.Do(ecpayReq)
	if err != nil {
		log.Printf("failed to create order: %v", err)
		return nil, err
	}
	defer ecpayResp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(ecpayResp.Body)
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}

	return bodyBytes, nil

	// w.Header().Set("Content-Type", "text/html")
	// w.Write(bodyBytes)

}

// QueryTradeInfo 查詢訂單(可依特店的需求來決定是否撰寫此 API)
func (c *Client) QueryTradeInfo() {
	ecpayReq, _ := http.NewRequest("POST", c.endpoint+"/Cashier/QueryTradeInfo/V5", strings.NewReader(formStr))
	ecpayReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

}

// QueryTrade 查詢信用卡單筆明細記錄
func (c *Client) QueryTrade() {
	ecpayReq, _ := http.NewRequest("POST", c.endpoint+"/CreditDetail/QueryTrade/V2", strings.NewReader(formStr))
	ecpayReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

}

// QueryPaymentInfo 查詢 ATM/CVS/BARCODE 取號結果
func (c *Client) QueryPaymentInfo() {
	ecpayReq, _ := http.NewRequest("POST", c.endpoint+"/Cashier/QueryPaymentInfo", strings.NewReader(formStr))
	ecpayReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

}

// QueryPaymentInfo 信用卡請退款功能 (若不撰寫此 API，則可透過廠商後台功能處理)
func (c *Client) DoAction() {
	ecpayReq, _ := http.NewRequest("POST", c.endpoint+"/CreditDetail/DoAction", strings.NewReader(formStr))
	ecpayReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

}

// QueryCreditCardPeriodInfo 信用卡定期定額訂單查詢
func (c *Client) QueryCreditCardPeriodInfo() {
	ecpayReq, _ := http.NewRequest("POST", c.endpoint+"/Cashier/QueryCreditCardPeriodInfo", strings.NewReader(formStr))
	ecpayReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

}

// TradeNoAio 下載特店對帳媒體檔
func (c *Client) TradeNoAio() {
	ecpayReq, _ := http.NewRequest("POST", c.vendor+"/PaymentMedia/TradeNoAio", strings.NewReader(formStr))
	ecpayReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

}

// FundingReconDetail 下載信用卡撥款對帳資料檔
func (c *Client) FundingReconDetail() {
	ecpayReq, _ := http.NewRequest("POST", c.vendor+"/CreditDetail/FundingReconDetail", strings.NewReader(formStr))
	ecpayReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

}
