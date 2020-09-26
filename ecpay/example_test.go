package ecpay_test

import (
	"fmt"
	"log"
	"time"

	"github.com/toastcheng/ecpay-sdk-go/ecpay"

	"github.com/google/uuid"
)

func ExampleNewClient() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}

	info := client.Info()
	fmt.Println(info["merchantID"])
	fmt.Println(info["hashKey"])
	fmt.Println(info["hashIV"])
	fmt.Println(info["endpoint"])
	fmt.Println(info["vendor"])
	// Output:
	// <MERCHANT_ID>
	// <HASH_KEY>
	// <HASH_IV>
	// https://payment-stage.ecpay.com.tw
	// https://vendor-stage.ecpay.com.tw

}

func ExampleClient_AioCheckOut_all() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}

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

	html, err := client.AioCheckOut(order)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)
}

func ExampleClient_AioCheckOut_atm() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}
	order := ecpay.Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: ecpay.FormatDatetime(time.Now()),
		PaymentType:       ecpay.PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          ecpay.FormatItemName([]string{"商品1", "商品2"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ecpay.ChoosePaymentTypeATM,
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
	}

	html, err := client.AioCheckOut(order)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)
}

func ExampleClient_AioCheckOut_barcode() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}
	order := ecpay.Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: ecpay.FormatDatetime(time.Now()),
		PaymentType:       ecpay.PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          ecpay.FormatItemName([]string{"商品1", "商品2"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ecpay.ChoosePaymentTypeBarCode,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: ecpay.NeedExtraPaidInfoTypeYes,
		InvoiceMark:       ecpay.InvoiceMarkTypeNo,
		CVSBarcode: &ecpay.CVSOrBarcodeParam{
			StoreExpireDate: 15,
			PaymentInfoURL:  "https://www.ecpay.com.tw/payment_info_url.php",
		},
	}

	html, err := client.AioCheckOut(order)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)

}

func ExampleClient_AioCheckOut_cvs() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}
	order := ecpay.Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: ecpay.FormatDatetime(time.Now()),
		PaymentType:       ecpay.PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          ecpay.FormatItemName([]string{"商品1", "商品2"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ecpay.ChoosePaymentTypeCVS,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: ecpay.NeedExtraPaidInfoTypeYes,
		InvoiceMark:       ecpay.InvoiceMarkTypeNo,
		CVSBarcode: &ecpay.CVSOrBarcodeParam{
			StoreExpireDate: 300,
			PaymentInfoURL:  "https://www.ecpay.com.tw/payment_info_url.php",
		},
	}

	html, err := client.AioCheckOut(order)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)

}

func ExampleClient_AioCheckOut_creditCard() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}
	order := ecpay.Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: ecpay.FormatDatetime(time.Now()),
		PaymentType:       ecpay.PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          ecpay.FormatItemName([]string{"商品1", "商品2"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ecpay.ChoosePaymentTypeCredit,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: ecpay.NeedExtraPaidInfoTypeYes,
		InvoiceMark:       ecpay.InvoiceMarkTypeNo,
		Credit: &ecpay.CreditParam{
			BindingCard: ecpay.BindingCardTypeNo,
			Redeem:      ecpay.RedeemTypeNo,
			UnionPay:    ecpay.UnionPayTypeSelect,
		},
	}

	html, err := client.AioCheckOut(order)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)

}

func ExampleClient_AioCheckOut_creditInstallment() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}
	order := ecpay.Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: ecpay.FormatDatetime(time.Now()),
		PaymentType:       ecpay.PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          ecpay.FormatItemName([]string{"商品1", "商品2"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ecpay.ChoosePaymentTypeCredit,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: ecpay.NeedExtraPaidInfoTypeYes,
		InvoiceMark:       ecpay.InvoiceMarkTypeNo,
		Credit: &ecpay.CreditParam{
			BindingCard: ecpay.BindingCardTypeNo,
			CreditInstallment: ecpay.FormatCreditInstallmentOption(
				ecpay.CreditInstallmentOption{
					Month3:  true,
					Month6:  true,
					Month12: true,
					Month18: true,
					Month24: true,
				},
			),
		},
	}

	html, err := client.AioCheckOut(order)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)

}

func ExampleClient_AioCheckOut_creditCardPeriod() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}
	order := ecpay.Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: ecpay.FormatDatetime(time.Now()),
		PaymentType:       ecpay.PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          ecpay.FormatItemName([]string{"商品1", "商品2"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ecpay.ChoosePaymentTypeCredit,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: ecpay.NeedExtraPaidInfoTypeYes,
		InvoiceMark:       ecpay.InvoiceMarkTypeNo,
		Credit: &ecpay.CreditParam{
			BindingCard:     ecpay.BindingCardTypeNo,
			PeriodAmount:    2000,
			PeriodType:      ecpay.PeriodTypeMonth,
			Frequency:       1,
			ExecTimes:       2,
			PeriodReturnURL: "https://www.ecpay.com.tw/receive.php",
		},
	}

	html, err := client.AioCheckOut(order)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)

}

func ExampleClient_AioCheckOut_googlePay() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}
	order := ecpay.Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: ecpay.FormatDatetime(time.Now()),
		PaymentType:       ecpay.PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          ecpay.FormatItemName([]string{"商品1...(*)", "商品2!!"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ecpay.ChoosePaymentTypeGooglePay,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: ecpay.NeedExtraPaidInfoTypeYes,
		InvoiceMark:       ecpay.InvoiceMarkTypeNo,
	}

	html, err := client.AioCheckOut(order)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)

}

func ExampleClient_AioCheckOut_webATM() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}
	order := ecpay.Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: ecpay.FormatDatetime(time.Now()),
		PaymentType:       ecpay.PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          ecpay.FormatItemName([]string{"商品1", "商品2"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ecpay.ChoosePaymentTypeWebATM,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: ecpay.NeedExtraPaidInfoTypeYes,
		InvoiceMark:       ecpay.InvoiceMarkTypeNo,
	}

	html, err := client.AioCheckOut(order)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)

}

func ExampleClient_AioCheckOut_invoice() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}
	order := ecpay.Order{
		MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: ecpay.FormatDatetime(time.Now()),
		PaymentType:       ecpay.PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          ecpay.FormatItemName([]string{"商品1", "商品2"}),
		ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
		ChoosePayment:     ecpay.ChoosePaymentTypeATM,
		ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
		ItemURL:           "https://www.ecpay.com.tw/item_url.php",
		Remark:            "交易備註",
		OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
		NeedExtraPaidInfo: ecpay.NeedExtraPaidInfoTypeYes,
		InvoiceMark:       ecpay.InvoiceMarkTypeYes,
		Invoice: &ecpay.InvoiceParam{
			RelateNumber:       "Tea9527",
			CustomerID:         "TEA_0000001",
			CustomerIdentifier: "53348111",
			CustomerName:       "客戶名稱",
			CustomerAddr:       "客戶地址",
			CustomerPhone:      "0912345678",
			CustomerEmail:      "abc@ecpay.com.tw",
			ClearanceMark:      ecpay.ClearanceMarkTypeCustoms,
			TaxType:            ecpay.TaxTypeZero,
			Donation:           ecpay.DonationTypeNo,
			Print:              ecpay.PrintTypeYes,
			InvoiceItemName:    ecpay.FormatInvoiceItem([]string{"測試商品1", "測試商品2"}),
			InvoiceItemCount:   ecpay.FormatInvoiceItem([]string{"2", "3"}),
			InvoiceItemWord:    ecpay.FormatInvoiceItem([]string{"個", "包"}),
			InvoiceItemPrice:   ecpay.FormatInvoiceItem([]string{"35", "10"}),
			InvoiceItemTaxType: ecpay.FormatInvoiceItem([]string{string(ecpay.TaxTypeFree), string(ecpay.TaxTypeFree)}),
			InvoiceRemark:      ecpay.FormatInvoiceItem([]string{"測試商品1的說明", "測試商品2的說明"}),
			DelayDay:           1,
		},
	}

	html, err := client.AioCheckOut(order)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)

}

func ExampleClient_QueryTradeInfo() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}
	info := ecpay.TradeInfo{
		MerchantTradeNo: "kncs20180804103309",
		TimeStamp:       string(time.Now().Unix()),
	}
	resp, err := client.QueryTradeInfo(info)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func ExampleClient_QueryTrade() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}
	trade := ecpay.Trade{
		CreditRefundID:  "10123456",
		CreditAmount:    100,
		CreditCheckCode: "59997889",
	}

	resp, err := client.QueryTrade(trade)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func ExampleClient_QueryPaymentInfo() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}
	info := ecpay.PaymentInfo{
		MerchantTradeNo: "kncs20180804103309",
		TimeStamp:       time.Now().Unix(),
	}

	resp, err := client.QueryPaymentInfo(info)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func ExampleClient_QueryCreditCardPeriodInfo() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}
	info := ecpay.CreditCardPeriodInfo{
		MerchantTradeNo: "kncs20180804103309",
		TimeStamp:       time.Now().Unix(),
	}

	resp, err := client.QueryCreditCardPeriodInfo(info)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func ExampleClient_DoAction() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}
	a := ecpay.CreditCardAction{
		MerchantTradeNo: "2000132",
		TradeNo:         "NO123",
		Action:          ecpay.ActionTypeC,
		TotalAmount:     100,
	}

	resp, err := client.DoAction(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func ExampleClient_FundingReconDetail() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}
	statement := ecpay.Statement{
		MerchantTradeNo: "NO123",
		PayDateType:     ecpay.PayDateTypeClose,
		BeginDate:       "2018-02-12",
		EndDate:         "2018-02-12",
		MediaFormated:   "1",
	}

	resp, err := client.FundingReconDetail(statement)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func ExampleClient_TradeNoAio() {
	client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}
	statement := ecpay.CreditCardStatement{
		DateType:        ecpay.DateTypeOrder,
		BeginDate:       "2018-02-12",
		EndDate:         "2018-02-12",
		PaymentType:     ecpay.MerchantPaymentTypeCreditCard,
		MerchantTradeNo: "NO123",
		MediaFormated:   ecpay.MediaFormatedTypeNew,
		AllocateStatus:  ecpay.AllocateStatusTypeDone,
		PaymentStatus:   ecpay.PaymentStatusTypePaid,
	}

	resp, err := client.TradeNoAio(statement)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
