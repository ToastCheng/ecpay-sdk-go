package ecpay

import (
	"ecpay/api"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/schema"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Config defines the configuration of datacollector server.
type Config struct {
	SQLUser     string
	SQLPassword string
	SQLEndpoint string
	SQLDB       string
	PaymentPort string
}

type impl struct {
	SQL *gorm.DB
}

func (s *impl) handleOrder(w http.ResponseWriter, r *http.Request) {
	req := api.OrderRequest{}
	switch (*r).Method {
	case "GET":
		schema.NewDecoder().Decode(&req, r.URL.Query())
	case "POST":
		json.NewDecoder(r.Body).Decode(&req)
	case "OPTIONS":
		return
	}

	log.Printf("OrderRequest: %v", req)

	form := api.NewECPayOrderFormData(req)
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

	bodyBytes, err := ioutil.ReadAll(ecpayResp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(bodyBytes)
}
