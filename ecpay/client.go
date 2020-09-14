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
	endpoint   string
	httpClient *http.Client
	options    ClientOption
}

// ClientOption .
type ClientOption func(*Client) error

// WithSandbox .
func WithSandbox(client *Client) error {
	client.endpoint = "https://payment-stage.ecpay.com.tw/"
	return nil
}

// NewClient .
func NewClient(merchantID string, options ...ClientOption) (*Client, error) {
	c := &Client{
		merchantID: merchantID,
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

// Order .
func (c *Client) Order() {

	form := api.NewECPayOrderFormData(api.OrderRequest{})
	formStr := form.Encode()
	formStr = strings.ReplaceAll(formStr, "-", "%2d")
	formStr = strings.ReplaceAll(formStr, "_", "%5f")
	// formStr = strings.ReplaceAll(formStr, ".", "%2e")
	formStr = strings.ReplaceAll(formStr, "*", "%2a")
	formStr = strings.ReplaceAll(formStr, "(", "%28")
	formStr = strings.ReplaceAll(formStr, ")", "%29")
	formStr = strings.ReplaceAll(formStr, "+", "%20")

	// formStr := "MerchantID=2000132&MerchantTradeNo=icuhricuevciu1&MerchantTradeDate=2020%2F03%2F29%2015%3A54%3A00&PaymentType=aio&ChoosePayment=Credit&ItemName=%E7%8E%A9%E5%85%B7&TotalAmount=100&CheckMacValue=83BD2D02AE61D5E498664781FCBF0083DEDC9CBC5A07AA220D02ADCA10B2059F&ReturnURL=https%3A%2F%2Fgoogle.com&TradeDesc=%E5%A5%BD%E7%8E%A9&EncryptType=1"
	log.Print(formStr)
	client := &http.Client{}
	ecpayReq, _ := http.NewRequest("POST", api.NewOrderUrl, strings.NewReader(formStr))
	// ecpayReq, _ := http.NewRequest("POST", "https://en3am9d2przom.x.pipedream.net/", strings.NewReader(formStr))
	ecpayReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	ecpayReq.Header.Add("accept", "*/*")
	ecpayReq.Header.Add("accept-encoding", "gzip, deflate, br")
	ecpayReq.Header.Add("user-agent", "PostmanRuntime/7.24.0")
	ecpayReq.Header.Add("cache-control", "no-cache")
	ecpayReq.Header.Add("postman-token", "40975c86-58b2-4b7e-aee6-9add8ed2cc09")

	// log.Printf("ecpay req: %v", ecpayReq)

	ecpayResp, err := client.Do(ecpayReq)
	// ecpayResp, err := http.PostForm("https://en3am9d2przom.x.pipedream.net/", form)
	if err != nil {
		log.Printf("failed to create order: %v", err)
	}
	defer ecpayResp.Body.Close()

	// bodyBytes, err := ioutil.ReadAll(ecpayResp.Body)
	// if err != nil {
		// log.Fatal(err)
	}

	// w.Header().Set("Content-Type", "text/html")
	// w.Write(bodyBytes)

}

// ATM、CVS 或 BARCODE 的取號結果通知.
func (c *Client) x() {

}
