package ecpay

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Payload defines the interface of payload for communicate with ECPay server.
type Payload interface {
	Validate() (bool, error)
	ToFormData(param string) url.Values
}

// Request defines the structure of request for communicate with ECPay server.
type Request struct {
	endpoint   string
	httpClient *http.Client
}

func (r *Request) Do(p Payload) (string, error) {
	if ok, err := p.Validate(); !ok {
		return "", err
	}

	form := p.ToFormData("2000132")

	formStr := form.Encode()
	formStr = strings.ReplaceAll(formStr, "-", "%2d")
	formStr = strings.ReplaceAll(formStr, "_", "%5f")
	formStr = strings.ReplaceAll(formStr, "*", "%2a")
	formStr = strings.ReplaceAll(formStr, "(", "%28")
	formStr = strings.ReplaceAll(formStr, ")", "%29")
	formStr = strings.ReplaceAll(formStr, "+", "%20")

	ecpayReq, _ := http.NewRequest("POST", r.endpoint, strings.NewReader(formStr))
	ecpayReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	ecpayReq.Header.Add("accept", "*/*")
	ecpayReq.Header.Add("accept-encoding", "gzip, deflate, br")
	ecpayReq.Header.Add("cache-control", "no-cache")

	ecpayResp, err := r.httpClient.Do(ecpayReq)
	if err != nil {
		log.Printf("failed to create order: %v", err)
		return "", err
	}
	defer ecpayResp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(ecpayResp.Body)
	if err != nil {
		return "", err
	}
	respStr := bytes.NewBuffer(bodyBytes).String()

	return respStr, nil
}
