package ecpay

import (
	"bytes"
	"log"
	"testing"
)

func TestCreateOrderAll(t *testing.T) {
	client, err := NewClient("2000132", "5294y06JbISpM5x9", "v77hoKGq4kWxNNIS", WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %v", err)
	}

	order := Order{
		MerchantTradeNo:   "NO123",
		MerchantTradeDate: "2020/10/10 10:10:10",
		PlatformID:        "3002599",
		ChoosePayment:     ALL,
		CustomerEmail:     "abc@gmail.com",
		TotalAmount:       100,
		PaymentType:       AIO,
		ItemNames:         []string{"item1", "item2"},
		TradeDesc:         "description",
		ReturnURL:         "https://abc.com",
		NeedExtraPaidInfo: false,
	}

	resp, err := client.AioCheckOut(order)
	if err != nil {
		log.Fatalf("failed to AioCheckOut: %v", err)
	}
	buf := bytes.NewBuffer(resp)
	log.Println(buf)
}
