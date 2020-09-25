# ECPay SDK for Go

[![license](https://img.shields.io/badge/license-MIT-blue)](https://github.com/toastcheng/ecpay/blob/master/LICENSE.md)
[![GoDoc](https://img.shields.io/badge/go-doc-blue)](https://pkg.go.dev/github.com/toastcheng/ecpay-sdk-go/ecpay)
[![Go Report Card](https://goreportcard.com/badge/github.com/toastcheng/ecpay-sdk-go)](https://goreportcard.com/report/github.com/toastcheng/ecpay-sdk-go)
[![Coverage Status](https://coveralls.io/repos/github/ToastCheng/ecpay-sdk-go/badge.svg)](https://coveralls.io/github/ToastCheng/ecpay-sdk-go)
[![GitHub Actions](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Ftoastcheng%2Fecpay-sdk-go%2Fbadge&style=flat-square)](https://actions-badge.atrox.dev/toastcheng/ecpay-sdk-go/goto)


## Introduction
ECPay SDK for Golang.
ECPay is a third party payment service in Taiwan, providing lots of payment options, such as ATM, web ATM, credit card, GooglePay, convenience code, etc.
For developer who wants to enable your app to have payment service, it would be a good choice :).

## Documentation
* Go Doc : https://pkg.go.dev/github.com/toastcheng/ecpay-sdk-go/ecpay
* ECPay API : https://www.ecpay.com.tw/Service/API_Dwnld

## Getting Started
### Installation
```
go get github.com/toastcheng/ecpay-sdk-go/ecpay
```

### Requirements
* Go 1.12+

### Quick Examples
#### Create a client:

for dev (sandbox)
```go
client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
```
for prod
```go
client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>")
```

#### create an order:
the order object is quite complex, make sure you check out the official document.
```go
order := Order{
    MerchantTradeNo: fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
    StoreID:           "",
    MerchantTradeDate: FormatDatetime(time.Now()),
    PaymentType:       PaymentTypeAIO,
    TotalAmount:       2000,
    TradeDesc:         "訂單測試",
    ItemName:          FormatItemName([]string{"商品1", "商品2"}),
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
```
The response is in HTML text, just display a few line here:
```html
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<!-- InstanceBegin template="/Templates/seller_template.dwt" codeOutsideHTMLIsLocked="false" -->
<head>
    <!-- SiteMap -->
    <meta name="google-site-verification" content="g1tlroYW-dChyLSinXJxV7BeP_T8nsDP1HpFSwORDgE" />
    <meta charset="UTF-8">
    <meta http-equiv="Content-Language" content="zh-TW">
    <meta name="viewport" content="width=device-width, initial-scale=1,maximum-scale=1.0, user-scalable=yes">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="format-detection" content="telephone=no">
    <meta name="renderer" content="webkit|ie-comp|ie-stand">
    <meta name="description" content="">
    <meta name="keywords" content="">
    <title>選擇支付方式|綠界科技</title>
    ...
```
Which provides UI for customers to submit their order:
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
    <body>
        <div class="site-body">
            <div class="site-main-wrapper">
                <div class="site-main">
                    <a href="#" class="main-pic">
                        <img src="https://payment-stage.ecpay.com.tw/Content/themes/WebStylePayment/images/other/bn_950x200_02.jpg?t=20200925152844" alt="綠界科技ECPay">
                    </a>
                </div>
            </div>
            <div id="ECPay" class="site-content-wrapper">
                <div class="site-content">
                    <p class="provider">金流服務由綠界科技ECPay提供 Payment cashflow service provided by ECPay</p>
                    <h3 class="content-title">訂單資訊 Order information</h3>
                    <div class="order-table o-info-1">
                        <dl>
                            <dt>訂單編號 Order number</dt>
                                <dd>0fd681601018943</dd>
                        </dl>
                        <dl>
                            <dt>商店名稱 Merchant&#39;s name</dt>
                            <dd>綠界測試店家</dd>
                        </dl>
                    </div>
                </div>
                <simplert :use-radius="true"
                        :use-icon="true"
                        ref="simplert">
                </simplert>
            </div>
        </div>
    </body>
</html>
...

#### Full example:
```go
package ecpay

import (
    "log"

    "github.com/toastcheng/ecpay-sdk-go/ecpay"
    "github.com/toastcheng/ecpay-sdk-go/ecpay/order"
)

func main() {
    // new a client.
    client, err := ecpay.NewClient("2000132", "5294y06JbISpM5x9", "v77hoKGq4kWxNNIS", ecpay.WithSandbox)
    if err != nil {
        log.Fatalf("failed to new client: %v", err)
    }
    order := Order{
		MerchantTradeNo: fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
		StoreID:           "",
		MerchantTradeDate: FormatDatetime(time.Now()),
		PaymentType:       PaymentTypeAIO,
		TotalAmount:       2000,
		TradeDesc:         "訂單測試",
		ItemName:          FormatItemName([]string{"商品1", "商品2"}),
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
    ...
}
```

There are more examples in `client_test.go` and Godoc. 

## Contributing

Contributions, issues and feature requests are welcome,
 please check out the [issues page](https://github.com/toastcheng/ecpay-sdk-go/issues).

## License

`ecpay-sdk-go` is available under the [MIT](https://github.com/toastcheng/ecpay-sdk-go/blob/master/LICENSE.md) license.
