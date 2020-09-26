# ECPay SDK for Go

[![license](https://img.shields.io/badge/license-MIT-blue)](https://github.com/toastcheng/ecpay/blob/master/LICENSE.md)
[![GoDoc](https://img.shields.io/badge/go-doc-blue)](https://pkg.go.dev/github.com/toastcheng/ecpay-sdk-go/ecpay)
[![Coverage Status](https://coveralls.io/repos/github/ToastCheng/ecpay-sdk-go/badge.svg)](https://coveralls.io/github/ToastCheng/ecpay-sdk-go)
[![GitHub Actions](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Ftoastcheng%2Fecpay-sdk-go%2Fbadge&style=flat-square)](https://actions-badge.atrox.dev/toastcheng/ecpay-sdk-go/goto)


## Introduction
ECPay SDK for Golang.

ECPay is a third party payment service, providing lots of payment options, such as ATM, web ATM, credit card, GooglePay, convenience payment code, etc.
For developer who wants to enable your app to have payment service, it would be a good choice :)

## Documentation
* pkg.go.dev : https://pkg.go.dev/github.com/toastcheng/ecpay-sdk-go
* ECPay API : https://www.ecpay.com.tw/Service/API_Dwnld

## Getting Started
### Installation
```
go get github.com/toastcheng/ecpay-sdk-go/ecpay
```

### Requirements
* Go 1.12+

### Quick Examples
#### Create a client

for dev (sandbox)
```go
client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", ecpay.WithSandbox)
```
for prod
```go
client, err := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>")
```

#### create an order
the order object is quite complex, make sure you check out the official document.
```go
order := ecpay.Order{
    MerchantTradeNo: fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
    StoreID:           "",
    MerchantTradeDate: ecpay.FormatDatetime(time.Now()),
    PaymentType:       ecpay.PaymentTypeAIO,
    TotalAmount:       2000,
    TradeDesc:         "訂單測試",
    ItemName:          ecpay.FormatItemName([]string{"商品1", "商品2"}),
    ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
    ChoosePayment:     ecpay.ChoosePaymentTypeAll,
    ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
    ItemURL:           "https://www.ecpay.com.tw/item_url.php'",
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
```
Since the ECPay API requires post form and needs to redirect the client to the ECPay host, which works more like PHP style.
To make this work in Go, `AioCheckOut` returns a html text that triggers redirect on client side.
```go
html, err := client.AioCheckOut(order)
```
You as a server, will need to pass this HTML to your frontend, make sure to add `Content-Type: text/html; charset=utf-8`. `text/html` makes the browser executes the html response instead of just displaying it; `charset=utf-8` makes the `CheckMacValue` validation in ECPay service pass. 

```go
mux.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
    ...
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprint(w, html)
}
```
The response will look something like this:
```html
<form id="5d0eb44f-e36f-45e4-80c6-579c7feba847"
    action="https://payment-stage.ecpay.com.tw/Cashier/AioCheckOut/V5"
    method="post"
>
    <input type="hidden" name="MerchantTradeDate" id="MerchantTradeDate" value="2020/09/26 23:16:49" />
    <input type="hidden" name="StoreExpireDate" id="StoreExpireDate" value="15" />
    <input type="hidden" name="Remark" id="Remark" value="交易備註" />
    <input type="hidden" name="MerchantTradeNo" id="MerchantTradeNo" value="0174e3409" />
    <input type="hidden" name="BindingCard" id="BindingCard" value="0" />
    <input type="hidden" name="OrderResultURL" id="OrderResultURL" value="https://www.ecpay.com.tw/order_result_url.php" />
    <input type="hidden" name="ItemName" id="ItemName" value="商品1*#商品2" />
    <input type="hidden" name="ReturnURL" id="ReturnURL" value="https://www.ecpay.com.tw/return_url.php" />
    <input type="hidden" name="ItemURL" id="ItemURL" value="https://www.ecpay.com.tw/item_url.php" />
    <input type="hidden" name="ChoosePayment" id="ChoosePayment" value="ALL" />
    <input type="hidden" name="NeedExtraPaidInfo" id="NeedExtraPaidInfo" value="Y" />
    <input type="hidden" name="ExpireDate" id="ExpireDate" value="7" />
    <input type="hidden" name="TotalAmount" id="TotalAmount" value="2000" />
    <input type="hidden" name="IgnorePayment" id="IgnorePayment" value="CVS" />
    <input type="hidden" name="PaymentType" id="PaymentType" value="aio" />
    <input type="hidden" name="EncryptType" id="EncryptType" value="1" />
    <input type="hidden" name="MerchantID" id="MerchantID" value="2000132" />
    <input type="hidden" name="CheckMacValue" id="CheckMacValue" value="8D48DA612C0C70B9C453D73D3C7513EC9D3600389D822DF2A00EB0639940DAFD" />
    <input type="hidden" name="PaymentInfoURL" id="PaymentInfoURL" value="https://www.ecpay.com.tw/payment_info_url.php" />
    <input type="hidden" name="TradeDesc" id="TradeDesc" value="訂單測試" />
    <input type="hidden" name="InvoiceMark" id="InvoiceMark" value="N" />
    <input type="hidden" name="ClientBackURL" id="ClientBackURL" value="https://www.ecpay.com.tw/client_back_url.php" />
    <script type="text/javascript">document.getElementById("5d0eb44f-e36f-45e4-80c6-579c7feba847").submit();</script>
</form>
```
This will redirect the client to ECPay payment page.
![](https://i.imgur.com/5QDFYdC.png)

There are more examples in `ecpay/example_test.go`, `example/main.go` and [pkg.go.dev](https://pkg.go.dev/github.com/toastcheng/ecpay-sdk-go). 

## Contributing

This project is in development, any contributions, issues and feature requests are welcome!
Please check out the [issues page](https://github.com/toastcheng/ecpay-sdk-go/issues).

## License

`ecpay-sdk-go` is available under the [MIT](https://github.com/toastcheng/ecpay-sdk-go/blob/master/LICENSE.md) license.
