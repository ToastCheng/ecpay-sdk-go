package ecpay

import (
	"log"
	"testing"

	"github.com/toastcheng/ecpay-sdk-go/ecpay/order"
	"github.com/toastcheng/ecpay-sdk-go/ecpay/payment"
	"github.com/toastcheng/ecpay-sdk-go/ecpay/trade"

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
	order := order.Order{
		MerchantTradeNo:   "NO123",
		MerchantTradeDate: "2020/10/10 10:10:10",
		PlatformID:        "3002599",
		ChoosePayment:     order.ChoosePaymentTypeAll,
		TotalAmount:       100,
		PaymentType:       order.PaymentTypeAIO,
		ItemName:          FormatItemName([]string{"item1", "item2"}),
		TradeDesc:         "description",
		ReturnURL:         "https://abc.com",
		NeedExtraPaidInfo: order.NeedExtraPaidInfoTypeNo,
		IgnorePayment: FormatIgnorePayment(IgnorePaymentOption{
			CVS: true,
		}),
		Invoice: &order.InvoiceParam{
			CustomerEmail:   "abc@gmail.com",
			CarrierType:     order.CarrierTypeCellphone,
			InvoiceItemName: FormatInvoiceItem([]string{"商品1", "商品2"}),
		},
		Credit: &order.CreditParam{},
		ATM: &order.ATMParam{
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
	info := trade.Info{
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
	trade := trade.Trade{
		CreditRefundID: "NO123",
	}

	_, err := client.QueryTrade(trade)
	if err != nil {
		log.Fatalf("failed to QueryTradeInfo: %v", err)
	}
}

func TestQueryPaymentInfo(t *testing.T) {
	client := getTestClient()
	info := payment.Info{
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
	info := payment.CreditCardPeriodInfo{
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
	a := payment.CreditCardAction{
		MerchantTradeNo: "NO123",
		Action:          payment.ActionTypeC,
	}

	_, err := client.DoAction(a)
	if err != nil {
		log.Fatalf("failed to QueryTradeInfo: %v", err)
	}
}

func TestTradeNoAio(t *testing.T) {
	client := getTestClient()
	statement := payment.Statement{
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
