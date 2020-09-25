package ecpay

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
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
		MerchantTradeNo: fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		// StoreID:           "",
		MerchantTradeDate: FormatDatetime(time.Now()),
		PaymentType:       PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          FormatItemName([]string{"商品1*", "商品2"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ChoosePaymentTypeAll,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php'",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: NeedExtraPaidInfoTypeYes,
		InvoiceMark:       InvoiceMarkTypeNo,
		ATM: &ATMParam{
			ExpireDate:     7,
			PaymentInfoURL: "https://www.ecpay.com.tw/payment_info_url.php",
		},
		CVSBarcode: &CVSOrBarcodeParam{
			StoreExpireDate: 15,
			PaymentInfoURL:  "https://www.ecpay.com.tw/payment_info_url.php",
		},
		IgnorePayment: FormatIgnorePayment(IgnorePaymentOption{
			CVS: true,
		}),
		Credit: &CreditParam{
			BindingCard: BindingCardTypeNo,
		},
	}

	resp, err := client.AioCheckOut(order)
	assert.NoError(t, err)
	assert.NotContains(t, resp, "交易失敗 Transaction failed")
}

func TestCreateOrderATM(t *testing.T) {
	client := getTestClient()
	order := Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: FormatDatetime(time.Now()),
		PaymentType:       PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          FormatItemName([]string{"商品1", "商品2"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ChoosePaymentTypeATM,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php'",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: NeedExtraPaidInfoTypeYes,
		InvoiceMark:       InvoiceMarkTypeNo,
		ATM: &ATMParam{
			ExpireDate:     7,
			PaymentInfoURL: "https://www.ecpay.com.tw/payment_info_url.php",
		},
	}

	resp, err := client.AioCheckOut(order)

	assert.NoError(t, err)
	assert.NotContains(t, resp, "交易失敗 Transaction failed")
}

func TestCreateOrderBarcode(t *testing.T) {
	client := getTestClient()
	order := Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: FormatDatetime(time.Now()),
		PaymentType:       PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          FormatItemName([]string{"商品1", "商品2"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ChoosePaymentTypeBarCode,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php'",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: NeedExtraPaidInfoTypeYes,
		InvoiceMark:       InvoiceMarkTypeNo,
		CVSBarcode: &CVSOrBarcodeParam{
			StoreExpireDate: 15,
			PaymentInfoURL:  "https://www.ecpay.com.tw/payment_info_url.php",
		},
	}

	resp, err := client.AioCheckOut(order)

	assert.NoError(t, err)
	assert.NotContains(t, resp, "交易失敗 Transaction failed")
}

func TestCreateOrderCVS(t *testing.T) {
	client := getTestClient()
	order := Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: FormatDatetime(time.Now()),
		PaymentType:       PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          FormatItemName([]string{"商品1", "商品2"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ChoosePaymentTypeCVS,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php'",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: NeedExtraPaidInfoTypeYes,
		InvoiceMark:       InvoiceMarkTypeNo,
		CVSBarcode: &CVSOrBarcodeParam{
			StoreExpireDate: 300,
			PaymentInfoURL:  "https://www.ecpay.com.tw/payment_info_url.php",
		},
	}

	resp, err := client.AioCheckOut(order)

	assert.NoError(t, err)
	assert.NotContains(t, resp, "交易失敗 Transaction failed")
}

func TestCreateOrderCredit(t *testing.T) {
	client := getTestClient()
	order := Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: FormatDatetime(time.Now()),
		PaymentType:       PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          FormatItemName([]string{"商品1", "商品2"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ChoosePaymentTypeCredit,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php'",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: NeedExtraPaidInfoTypeYes,
		InvoiceMark:       InvoiceMarkTypeNo,
		Credit: &CreditParam{
			BindingCard: BindingCardTypeNo,
			Redeem:      RedeemTypeNo,
			UnionPay:    UnionPayTypeSelect,
		},
	}

	resp, err := client.AioCheckOut(order)

	assert.NoError(t, err)
	assert.NotContains(t, resp, "交易失敗 Transaction failed")
}

func TestCreateOrderCreditInstallment(t *testing.T) {
	client := getTestClient()
	order := Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: FormatDatetime(time.Now()),
		PaymentType:       PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          FormatItemName([]string{"商品1", "商品2"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ChoosePaymentTypeCredit,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php'",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: NeedExtraPaidInfoTypeYes,
		InvoiceMark:       InvoiceMarkTypeNo,
		Credit: &CreditParam{
			BindingCard: BindingCardTypeNo,
			CreditInstallment: FormatCreditInstallmentOption(
				CreditInstallmentOption{
					Month3:  true,
					Month6:  true,
					Month12: true,
					Month18: true,
					Month24: true,
				},
			),
		},
	}

	resp, err := client.AioCheckOut(order)

	assert.NoError(t, err)
	assert.NotContains(t, resp, "交易失敗 Transaction failed")
}

func TestCreateOrderCreditPeriod(t *testing.T) {
	client := getTestClient()
	order := Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: FormatDatetime(time.Now()),
		PaymentType:       PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          FormatItemName([]string{"商品1", "商品2"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ChoosePaymentTypeCredit,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php'",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: NeedExtraPaidInfoTypeYes,
		InvoiceMark:       InvoiceMarkTypeNo,
		Credit: &CreditParam{
			BindingCard:     BindingCardTypeNo,
			PeriodAmount:    2000,
			PeriodType:      PeriodTypeMonth,
			Frequency:       1,
			ExecTimes:       2,
			PeriodReturnURL: "https://www.ecpay.com.tw/receive.php",
		},
	}

	resp, err := client.AioCheckOut(order)

	assert.NoError(t, err)
	assert.NotContains(t, resp, "交易失敗 Transaction failed")
}

func TestCreateOrderGooglePay(t *testing.T) {
	client := getTestClient()
	order := Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: FormatDatetime(time.Now()),
		PaymentType:       PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          FormatItemName([]string{"商品1...(*)", "商品2!!"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ChoosePaymentTypeGooglePay,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: NeedExtraPaidInfoTypeYes,
		InvoiceMark:       InvoiceMarkTypeNo,
	}

	resp, err := client.AioCheckOut(order)

	assert.NoError(t, err)
	assert.NotContains(t, resp, "交易失敗 Transaction failed")
}

func TestCreateOrderWebATM(t *testing.T) {
	client := getTestClient()
	order := Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: FormatDatetime(time.Now()),
		PaymentType:       PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          FormatItemName([]string{"商品1", "商品2"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ChoosePaymentTypeWebATM,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php'",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: NeedExtraPaidInfoTypeYes,
		InvoiceMark:       InvoiceMarkTypeNo,
	}

	resp, err := client.AioCheckOut(order)

	assert.NoError(t, err)
	assert.NotContains(t, resp, "交易失敗 Transaction failed")
}

func TestCreateOrderInvoice(t *testing.T) {
	client := getTestClient()
	order := Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: FormatDatetime(time.Now()),
		PaymentType:       PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          FormatItemName([]string{"商品1", "商品2"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ChoosePaymentTypeATM,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: NeedExtraPaidInfoTypeYes,
		InvoiceMark:       InvoiceMarkTypeYes,
		Invoice: &InvoiceParam{
			RelateNumber:       "Tea9527",
			CustomerID:         "TEA_0000001",
			CustomerIdentifier: "53348111",
			CustomerName:       "客戶名稱",
			CustomerAddr:       "客戶地址",
			CustomerPhone:      "0912345678",
			CustomerEmail:      "abc@ecpay.com.tw",
			ClearanceMark:      ClearanceMarkTypeCustoms,
			TaxType:            TaxTypeZero,
			Donation:           DonationTypeNo,
			Print:              PrintTypeYes,
			InvoiceItemName:    FormatInvoiceItem([]string{"測試商品1", "測試商品2"}),
			InvoiceItemCount:   FormatInvoiceItem([]string{"2", "3"}),
			InvoiceItemWord:    FormatInvoiceItem([]string{"個", "包"}),
			InvoiceItemPrice:   FormatInvoiceItem([]string{"35", "10"}),
			InvoiceItemTaxType: FormatInvoiceItem([]string{string(TaxTypeFree), string(TaxTypeFree)}),
			InvoiceRemark:      FormatInvoiceItem([]string{"測試商品1的說明", "測試商品2的說明"}),
			DelayDay:           1,
		},
	}

	resp, err := client.AioCheckOut(order)

	assert.NoError(t, err)
	assert.NotContains(t, resp, "交易失敗 Transaction failed")
}

func TestQueryTradeInfo(t *testing.T) {
	client := getTestClient()
	info := TradeInfo{
		MerchantTradeNo: "kncs20180804103309",
		TimeStamp:       string(time.Now().Unix()),
	}
	resp, err := client.QueryTradeInfo(info)

	assert.NoError(t, err)
	assert.Equal(t, "kncs20180804103309", resp["MerchantTradeNo"])
}

func TestQueryTrade(t *testing.T) {
	client := getTestClient()
	trade := Trade{
		CreditRefundID:  "10123456",
		CreditAmount:    100,
		CreditCheckCode: "59997889",
	}

	_, err := client.QueryTrade(trade)
	assert.NoError(t, err)
}

func TestQueryPaymentInfo(t *testing.T) {
	client := getTestClient()
	info := PaymentInfo{
		MerchantTradeNo: "kncs20180804103309",
		TimeStamp:       time.Now().Unix(),
	}

	_, err := client.QueryPaymentInfo(info)
	assert.NoError(t, err)
}

func TestQueryCreditCardPeriodInfo(t *testing.T) {
	client := getTestClient()
	info := CreditCardPeriodInfo{
		MerchantTradeNo: "kncs20180804103309",
		TimeStamp:       time.Now().Unix(),
	}

	_, err := client.QueryCreditCardPeriodInfo(info)
	assert.NoError(t, err)
}

func TestDoAction(t *testing.T) {
	client := getTestClient()
	a := CreditCardAction{
		MerchantTradeNo: "2000132",
		TradeNo:         "NO123",
		Action:          ActionTypeC,
		TotalAmount:     100,
	}

	resp, err := client.DoAction(a)
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{"2000132"}, resp["Merchant"])
}

func TestFundingReconDetail(t *testing.T) {
	client := getTestClient()
	statement := Statement{
		MerchantTradeNo: "NO123",
		PayDateType:     PayDateTypeClose,
		BeginDate:       "2018-02-12",
		EndDate:         "2018-02-12",
		MediaFormated:   "1",
	}

	_, err := client.FundingReconDetail(statement)
	assert.NoError(t, err)
}

func TestDownloadMerchantBalance(t *testing.T) {
	client := getTestClient()
	statement := CreditCardStatement{
		DateType:        DateTypeOrder,
		BeginDate:       "2018-02-12",
		EndDate:         "2018-02-12",
		PaymentType:     MerchantPaymentTypeCreditCard,
		MerchantTradeNo: "NO123",
		MediaFormated:   MediaFormatedTypeNew,
		AllocateStatus:  AllocateStatusTypeDone,
		PaymentStatus:   PaymentStatusTypePaid,
	}

	_, err := client.TradeNoAio(statement)
	assert.NoError(t, err)
}
