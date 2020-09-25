package ecpay

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
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
	order := Order{
		MerchantTradeNo:   "NO123",
		MerchantTradeDate: "2020/10/10 10:10:10",
		PlatformID:        "3002599",
		ChoosePayment:     ChoosePaymentTypeAll,
		TotalAmount:       100,
		PaymentType:       PaymentTypeAIO,
		ItemName:          FormatItemName([]string{"item1", "item2"}),
		TradeDesc:         "description",
		ReturnURL:         "https://abc.com",
		NeedExtraPaidInfo: NeedExtraPaidInfoTypeNo,
		IgnorePayment: FormatIgnorePayment(IgnorePaymentOption{
			CVS: true,
		}),
		Invoice: &InvoiceParam{
			CustomerEmail:   "abc@gmail.com",
			CarrierType:     CarrierTypeCellphone,
			InvoiceItemName: FormatInvoiceItem([]string{"商品1", "商品2"}),
		},
		Credit: &CreditParam{},
		ATM: &ATMParam{
			ExpireDate: 34,
		},
	}

	resp, err := client.AioCheckOut(order)
	if err != nil {
		log.Fatalf("failed to AioCheckOut: %v", err)
	}
	assert.Contains(t, resp, "交易失敗 Transaction failed")
}

func TestQueryTradeInfo(t *testing.T) {
	client := getTestClient()
	info := TradeInfo{
		MerchantTradeNo: "NO123",
		TimeStamp:       "2020/10/10 10:10:10",
	}

	_, err := client.QueryTradeInfo(info)
	if err != nil {
		log.Fatalf("failed to QueryTradeInfo: %v", err)
	}
}

func TestQueryTrade(t *testing.T) {
	client := getTestClient()
	trade := Trade{
		CreditRefundID: "NO123",
	}

	_, err := client.QueryTrade(trade)
	if err != nil {
		log.Fatalf("failed to QueryTradeInfo: %v", err)
	}
}

func TestQueryPaymentInfo(t *testing.T) {
	client := getTestClient()
	info := PaymentInfo{
		MerchantTradeNo: "NO123",
		TimeStamp:       "2020/10/10 10:10:10",
	}

	_, err := client.QueryPaymentInfo(info)
	if err != nil {
		log.Fatalf("failed to QueryTradeInfo: %v", err)
	}
}

func TestQueryCreditCardPeriodInfo(t *testing.T) {
	client := getTestClient()
	info := CreditCardPeriodInfo{
		MerchantTradeNo: "NO123",
		TimeStamp:       "2020/10/10 10:10:10",
	}

	_, err := client.QueryCreditCardPeriodInfo(info)
	if err != nil {
		log.Fatalf("failed to QueryTradeInfo: %v", err)
	}
}

func TestDoAction(t *testing.T) {
	client := getTestClient()
	a := CreditCardAction{
		MerchantTradeNo: "NO123",
		Action:          ActionTypeC,
	}

	_, err := client.DoAction(a)
	if err != nil {
		log.Fatalf("failed to QueryTradeInfo: %v", err)
	}
}

func TestTradeNoAio(t *testing.T) {
	client := getTestClient()
	statement := Statement{
		MerchantTradeNo: "NO123",
		DateType:        "2",
		BeginDate:       "2015-02-12",
		EndDate:         "2016-02-12",
		MediaFormated:   "1",
	}

	_, err := client.TradeNoAio(statement)
	if err != nil {
		log.Fatalf("failed to QueryTradeInfo: %v", err)
	}
}
