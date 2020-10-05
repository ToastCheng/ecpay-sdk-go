package ecpay

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var order = Order{
	MerchantTradeNo:   fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()%10000),
	StoreID:           "",
	MerchantTradeDate: FormatDatetime(time.Now()),
	PaymentType:       PaymentTypeAIO,
	TotalAmount:       2000,
	TradeDesc:         "訂單測試",
	ItemName:          FormatItemName([]string{"商品1*", "商品2"}),
	ReturnURL:         "https://www.ecpay.com.tw/return_url.php",
	ChoosePayment:     ChoosePaymentTypeAll,
	ClientBackURL:     "https://www.ecpay.com.tw/client_back_url.php",
	ItemURL:           "https://www.ecpay.com.tw/item_url.php",
	Remark:            "交易備註",
	OrderResultURL:    "https://www.ecpay.com.tw/order_result_url.php",
	NeedExtraPaidInfo: NeedExtraPaidInfoTypeYes,
	InvoiceMark:       InvoiceMarkTypeNo,
	IgnorePayment: FormatIgnorePayment(IgnorePaymentOption{
		CVS: true,
	}),
	Credit: &CreditParam{
		BindingCard: BindingCardTypeNo,
	},
}

var inv = InvoiceParam{
	Print:         PrintTypeYes,
	CustomerName:  "customer",
	CustomerAddr:  "addr",
	CustomerEmail: "abc@gmail.com",
}

var atm = ATMParam{
	ExpireDate:     7,
	PaymentInfoURL: "https://www.ecpay.com.tw/payment_info_url.php",
}

var cvs = CVSOrBarcodeParam{
	StoreExpireDate: 15,
	PaymentInfoURL:  "https://www.ecpay.com.tw/payment_info_url.php",
}

func TestOrderValidate(t *testing.T) {
	tmp := order
	tmp.MerchantTradeNo = ""
	_, err := tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "MerchantTradeNo should not be empty")
	}

	tmp = order
	tmp.MerchantTradeDate = ""
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "MerchantTradeDate should not be empty")
	}

	tmp = order
	tmp.ChoosePayment = ""
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "ChoosePayment should not be empty")
	}

	tmp = order
	tmp.TotalAmount = 0
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "TotalAmount should not be empty")
	}

	tmp = order
	tmp.PaymentType = ""
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "PaymentType should not be empty")
	}

	tmp = order
	tmp.ItemName = ""
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "ItemName should not be empty and within 400 words")
	}

	tmp = order
	tmp.TradeDesc = ""
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "TradeDesc should not be empty")
	}

	tmp = order
	tmp.ReturnURL = ""
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "ReturnURL should not be empty")
	}

	tmp = order
	tmp.StoreID = "123456789012345678901"
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "StoreID should not exceed 20")
	}

	tmp = order
	tmp.TradeDesc = "804304443246847464624451052946404372779043697628605666592305562534148629940304465617179333954795551990724254310299180171772710477455596778622192832905674187436264322793058067343244698610870679538077248"
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "TradeDesc should not exceed 200")
	}

	tmp = order
	tmp.ItemName = "804304443246847464624451052946404372779043697628605666592305562534148629940304465617179333954795551990724254310299180171772710477455596778622192832905674187436264322793058067343244698610870679538077248"
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "ItemName should not exceed 200")
	}

	tmp = order
	tmp.ReturnURL = "804304443246847464624451052946404372779043697628605666592305562534148629940304465617179333954795551990724254310299180171772710477455596778622192832905674187436264322793058067343244698610870679538077248"
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "ReturnURL should not exceed 200")
	}

	tmp = order
	tmp.ClientBackURL = "804304443246847464624451052946404372779043697628605666592305562534148629940304465617179333954795551990724254310299180171772710477455596778622192832905674187436264322793058067343244698610870679538077248"
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "ClientBackURL should not exceed 200")
	}

	tmp = order
	tmp.ItemURL = "804304443246847464624451052946404372779043697628605666592305562534148629940304465617179333954795551990724254310299180171772710477455596778622192832905674187436264322793058067343244698610870679538077248"
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "ItemURL should not exceed 200")
	}

	tmp = order
	tmp.OrderResultURL = "804304443246847464624451052946404372779043697628605666592305562534148629940304465617179333954795551990724254310299180171772710477455596778622192832905674187436264322793058067343244698610870679538077248"
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "OrderResultURL should not exceed 200")
	}

	tmp = order
	tmp.CustomField1 = "804304443246847464624451052946404372779043697628605"
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "CustomField should not exceed 50")
	}

	tmp = order
	tmpAtm := atm
	tmp.ATM = &tmpAtm
	tmp.ATM.ExpireDate = 61
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "ExpireDate should be in the range of 1-60")
	}

	tmp = order
	tmpCvs := cvs
	tmp.CVSBarcode = &tmpCvs
	tmp.CVSBarcode.Desc1 = "804304443246847464624"
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Desc should not exceed 20")
	}

	tmp = order
	tmpInv := inv
	tmp.Invoice = &tmpInv
	tmp.Invoice.CustomerIdentifier = "1"
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "CustomerIdentifier has to fill fixed length of 8 digits")
	}

	tmp = order
	tmpInv = inv
	tmp.Invoice = &tmpInv
	tmp.Invoice.Print = PrintTypeNo
	tmp.Invoice.CustomerIdentifier = "12345678"
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Print has to be true, when CustomerIdentifier have value")
	}

	tmp = order
	tmpInv = inv
	tmp.Invoice = &tmpInv
	tmp.Invoice.Donation = DonationTypeYes
	tmp.Invoice.CustomerIdentifier = "12345678"
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Donation has to be false, when CustomerIdentifier have value")
	}

	tmp = order
	tmpInv = inv
	tmp.Invoice = &tmpInv
	tmp.Invoice.CustomerName = ""
	tmp.Invoice.Print = PrintTypeYes
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "CustomerName should not be empty if Print is true")
	}

	tmp = order
	tmpInv = inv
	tmp.Invoice = &tmpInv
	tmp.Invoice.Print = PrintTypeYes
	tmp.Invoice.CustomerAddr = ""
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "CustomerAddr should not be empty if Print is true")
	}

	tmp = order
	tmpInv = inv
	tmp.Invoice = &tmpInv
	tmp.Invoice.CustomerEmail = ""
	tmp.Invoice.CustomerPhone = ""
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "CustomerPhone should not be empty if CustomerEmail is empty")
	}

	tmp = order
	tmpInv = inv
	tmp.Invoice = &tmpInv
	tmp.Invoice.Print = PrintTypeYes
	tmp.Invoice.Donation = DonationTypeYes
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Print should be false if Donation is set to true")
	}

	tmp = order
	tmpInv = inv
	tmp.Invoice = &tmpInv
	tmp.Invoice.LoveCode = ""
	tmp.Invoice.Donation = DonationTypeYes
	tmp.Invoice.Print = PrintTypeNo
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "LoveCode should not be empty if Donation is set to true")
	}

	tmp = order
	tmpInv = inv
	tmp.Invoice = &tmpInv
	tmp.Invoice.LoveCode = "1"
	tmp.Invoice.Donation = DonationTypeYes
	tmp.Invoice.Print = PrintTypeNo
	_, err = tmp.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "LoveCode should be a 3-7 digit number")
	}
}
