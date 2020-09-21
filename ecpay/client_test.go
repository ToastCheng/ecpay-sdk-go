package ecpay

import (
	"bytes"
	"ecpay/order"
	"ecpay/order/payment"
	"log"
	"strings"
	"testing"
)

func getTestClient() *Client {
	client, err := NewClient("2000132", "5294y06JbISpM5x9", "v77hoKGq4kWxNNIS", WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %v", err)
	}
	return client
}

func TestCreateOrderAll(t *testing.T) {
	client := getTestClient()
	order := order.Order{
		MerchantTradeNo:   "NO123",
		MerchantTradeDate: "2020/10/10 10:10:10",
		PlatformID:        "3002599",
		ChoosePayment:     payment.ALL,
		TotalAmount:       100,
		PaymentType:       payment.AIO,
		ItemNames:         []string{"item1", "item2"},
		TradeDesc:         "description",
		ReturnURL:         "https://abc.com",
		NeedExtraPaidInfo: false,
		Invoice: &order.InvoiceParam{
			CustomerEmail: "abc@gmail.com",
		},
	}

	resp, err := client.AioCheckOut(order)
	if err != nil {
		log.Fatalf("failed to AioCheckOut: %v", err)
	}
	respStr := bytes.NewBuffer(resp).String()
	if strings.Contains(respStr, "交易失敗 Transaction failed") {
		log.Fatalf("failed to AioCheckOut: %s", respStr)
	}
}
