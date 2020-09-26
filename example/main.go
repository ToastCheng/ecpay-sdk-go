package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/toastcheng/ecpay-sdk-go/ecpay"
)

func main() {
	client, err := ecpay.NewClient("2000132", "5294y06JbISpM5x9", "v77hoKGq4kWxNNIS", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
		order := ecpay.Order{
			MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
			StoreID:           "",
			MerchantTradeDate: ecpay.FormatDatetime(time.Now()),
			PaymentType:       ecpay.PaymentTypeAIO,
			TotalAmount:       2000,
			TradeDesc:         "訂單測試",
			ItemName:          ecpay.FormatItemName([]string{"商品1*", "商品2"}),
			ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
			ChoosePayment:     ecpay.ChoosePaymentTypeAll,
			ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
			ItemURL:           "https://www.ecpay.com.tw/item_url.php",
			Remark:            "交易備註",
			OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
			NeedExtraPaidInfo: ecpay.NeedExtraPaidInfoTypeYes,
			InvoiceMark:       ecpay.InvoiceMarkTypeNo,
			ATM: &ecpay.ATMParam{
				ExpireDate:     7,
				PaymentInfoURL: "https://www.ecpay.com.tw/payment_info_url.php",
			},
			CVSBarcode: &ecpay.CVSOrBarcodeParam{
				StoreExpireDate: 15,
				PaymentInfoURL:  "https://www.ecpay.com.tw/payment_info_url.php",
			},
			IgnorePayment: ecpay.FormatIgnorePayment(ecpay.IgnorePaymentOption{
				CVS: true,
			}),
			Credit: &ecpay.CreditParam{
				BindingCard: ecpay.BindingCardTypeNo,
			},
		}
		html, e := client.AioCheckOut(order)
		if e != nil {
			log.Print(e)
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, html)

	})
	mux.HandleFunc("/periodinfo", func(w http.ResponseWriter, r *http.Request) {
		info := ecpay.CreditCardPeriodInfo{
			MerchantTradeNo: "kncs20180804103310",
			TimeStamp:       time.Now().Unix(),
		}

		resp, err := client.QueryCreditCardPeriodInfo(info)
		if err != nil {
			log.Print(err)
		}
		log.Print(resp)
		fmt.Fprintf(w, "%v", resp)
	})
	http.ListenAndServe(":8080", mux)
}
