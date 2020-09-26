# ECPay SDK for Go

[![license](https://img.shields.io/badge/license-MIT-blue)](https://github.com/toastcheng/ecpay/blob/master/LICENSE.md)
[![GoDoc](https://img.shields.io/badge/go-doc-blue)](https://pkg.go.dev/github.com/toastcheng/ecpay-sdk-go/ecpay)
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

html, err := client.AioCheckOut(order)
```
In your server, you will need to pass this HTML to your frontend, see `example/main.go` for example:
```go
mux.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
    ...
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprint(w, html)
}
```
This will redirect the client to ECPay payment page.

There are more examples in `ecpay/example_test.go`, `example/main.go` and Godoc. 

## Contributing

This project is in development, any contributions, issues and feature requests are welcome!
Please check out the [issues page](https://github.com/toastcheng/ecpay-sdk-go/issues).

## License

`ecpay-sdk-go` is available under the [MIT](https://github.com/toastcheng/ecpay-sdk-go/blob/master/LICENSE.md) license.
