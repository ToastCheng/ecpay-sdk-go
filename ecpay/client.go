package ecpay

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Client implements client for making ECPay api.
type Client struct {
	// merchantID (特店編號).
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

func (c *Client) do(p Payload) (*http.Response, error) {
	if ok, err := p.Validate(); !ok {
		return nil, err
	}

	form := p.ToFormData()
	form.Set("MerchantID", c.merchantID)
	form.Set("CheckMacValue", GetCheckMacValue(form, c.hashKey, c.hashIV))

	formStr := form.Encode()
	formStr = strings.ReplaceAll(formStr, "-", "%2d")
	formStr = strings.ReplaceAll(formStr, "_", "%5f")
	formStr = strings.ReplaceAll(formStr, "*", "%2a")
	formStr = strings.ReplaceAll(formStr, "(", "%28")
	formStr = strings.ReplaceAll(formStr, ")", "%29")
	formStr = strings.ReplaceAll(formStr, "+", "%20")

	var endpoint string
	switch p.(type) {
	case Order:
		endpoint = c.endpoint + "/Cashier/AioCheckOut/V5"
	case TradeInfo:
		endpoint = c.endpoint + "/Cashier/QueryTradeInfo/V5"
	case PaymentInfo:
		endpoint = c.endpoint + "/Cashier/QueryPaymentInfo"
	case Trade:
		endpoint = c.endpoint + "/CreditDetail/QueryTrade/V2"
	case CreditCardPeriodInfo:
		endpoint = c.endpoint + "/Cashier/QueryCreditCardPeriodInfo"
	case Statement:
		endpoint = c.vendor + "/PaymentMedia/TradeNoAio"
	case CreditCardAction:
		endpoint = c.endpoint + "/CreditDetail/DoAction"
	case CreditCardStatement:
		endpoint = c.endpoint + "/CreditDetail/FundingReconDetail"
	default:
		endpoint = c.endpoint
	}
	req, _ := http.NewRequest("POST", endpoint, strings.NewReader(formStr))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-encoding", "gzip, deflate, br")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("failed to create order: %v", err)
		return nil, err
	}

	return resp, nil
}

// AioCheckOut sends an order to ECPay server (產生訂單).
func (c *Client) AioCheckOut(order Order) (string, error) {
	resp, err := c.do(order)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

// QueryTradeInfo queries a single trade info (查詢訂單).
func (c *Client) QueryTradeInfo(info TradeInfo) (string, error) {
	resp, err := c.do(info)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

// QueryTrade queries a single creadit card trade (查詢信用卡單筆明細記錄).
func (c *Client) QueryTrade(trade Trade) (map[string]interface{}, error) {
	resp, err := c.do(trade)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

// QueryPaymentInfo queries payment info of ATM/CVS/Barcode (查詢 ATM/CVS/BARCODE 取號結果).
func (c *Client) QueryPaymentInfo(info PaymentInfo) (map[string]interface{}, error) {
	resp, err := c.do(info)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	m, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		log.Fatal(err)
	}
	dataBytes, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	json.Unmarshal(dataBytes, &result)

	return result, nil
}

// QueryCreditCardPeriodInfo queries credit card periodic payment (信用卡定期定額訂單查詢).
func (c *Client) QueryCreditCardPeriodInfo(info CreditCardPeriodInfo) (map[string]interface{}, error) {
	resp, err := c.do(info)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}

// DoAction fires an credit card refund action (信用卡請退款功能).
func (c *Client) DoAction(action CreditCardAction) (map[string]interface{}, error) {
	resp, err := c.do(action)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	m, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		log.Fatal(err)
	}
	dataBytes, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	json.Unmarshal(dataBytes, &result)

	return result, nil
}

// TradeNoAio downloads the member statement (下載特店對帳媒體檔).
func (c *Client) TradeNoAio(statement Statement) (string, error) {
	resp, err := c.do(statement)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

// FundingReconDetail downloads the member statement (下載信用卡撥款對帳資料檔).
func (c *Client) FundingReconDetail(statement CreditCardStatement) (string, error) {
	resp, err := c.do(statement)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}
