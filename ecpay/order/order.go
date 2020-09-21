package order

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/toastcheng/ecpay-sdk-go/ecpay/order/carrier"
	"github.com/toastcheng/ecpay-sdk-go/ecpay/order/payment"
	"github.com/toastcheng/ecpay-sdk-go/ecpay/order/period"
	"github.com/toastcheng/ecpay-sdk-go/ecpay/order/subpayment"

	"github.com/google/uuid"
)

// Order defines the structure of an order.
type Order struct {
	MerchantTradeNo   string
	StoreID           string
	MerchantTradeDate string
	TotalAmount       int
	TradeDesc         string
	ItemNames         []string
	ItemURL           string
	ReturnURL         string
	ClientBackURL     string
	OrderResultURL    string
	PaymentType       payment.PaymentType
	ChoosePayment     payment.ChoosePaymentType
	ChooseSubPayment  subpayment.ChooseSubpaymentType
	Remark            string
	NeedExtraPaidInfo bool
	DeviceSource      string
	IgnorePayment     string
	PlatformID        string
	InvoiceMark       bool
	CustomField1      string
	CustomField2      string
	CustomField3      string
	CustomField4      string
	EncryptType       string

	ATM        *ATMParam
	CVSBarcode *CVSOrBarcodeParam
	Credit     *CreditParam
	Invoice    *InvoiceParam
}

// ATMParam defines the parameters tailored for ATM transaction.
type ATMParam struct {
	ExpireDate        string
	PaymentInfoURL    string
	ClientRedirectURL string
}

// CVSOrBarcodeParam defines the parameters tailored for CVS or bar code transaction.
type CVSOrBarcodeParam struct {
	StoreExpireDate   string
	Desc1             string
	Desc2             string
	Desc3             string
	Desc4             string
	PaymentInfoURL    string
	ClientRedirectURL string
}

// CreditParam defines the parameters tailored for credit card transaction.
type CreditParam struct {
	BindingCard      string
	MerchantMemberID string
	Language         string

	Redeem            string
	UnionPay          string
	CreditInstallment string

	PeriodAmount    int
	PeriodType      period.Period
	Frequency       int
	ExecTimes       int
	PeriodReturnURL string
}

// InvoiceParam defines the parameters for invoice specific settings.
type InvoiceParam struct {
	RelateNumber       string
	CustomerID         string
	CustomerIdentifier string
	CustomerName       string
	CustomerAddr       string
	CustomerPhone      string
	CustomerEmail      string
	ClearanceMark      string
	TaxType            string
	CarruerType        carrier.CarrierType
	CarruerNum         string
	Donation           bool
	LoveCode           string
	Print              bool
	InvoiceItemName    string
	InvoiceItemCount   string
	InvoiceItemWord    string
	InvoiceItemPrice   string
	InvoiceItemTaxType string
	InvoiceRemark      string
	DelayDay           string
	InvType            string
}

// Validate validate if the order struct is valid.
func (o *Order) Validate() (bool, error) {
	// check null.
	if o.MerchantTradeNo == "" {
		return false, errors.New("MerchantTradeNo should not be empty")
	}
	if o.MerchantTradeDate == "" {
		return false, errors.New("MerchantTradeDate should not be empty")
	}
	if o.PlatformID == "" {
		return false, errors.New("PlatformID should not be empty")
	}
	if o.ChoosePayment == "" {
		return false, errors.New("ChoosePayment should not be empty")
	}
	if o.TotalAmount == 0 {
		return false, errors.New("TotalAmount should not be empty")
	}
	if o.PaymentType == "" {
		return false, errors.New("PaymentType should not be empty")
	}
	if len(o.ItemNames) == 0 {
		return false, errors.New("ItemNames should not be empty")
	}
	if o.TradeDesc == "" {
		return false, errors.New("TradeDesc should not be empty")
	}
	if o.ReturnURL == "" {
		return false, errors.New("ReturnURL should not be empty")
	}

	// check length.
	if len(o.MerchantTradeNo) > 10 {
		return false, errors.New("MerchantTradeNo should not exceed 10")
	}
	if len(o.StoreID) > 20 {
		return false, errors.New("StoreID should not exceed 20")
	}
	if len(o.TradeDesc) > 200 {
		return false, errors.New("TradeDesc should not exceed 200")
	}
	if len(o.ItemNames) > 200 {
		return false, errors.New("ItemName should not exceed 200")
	}
	if len(o.ReturnURL) > 200 {
		return false, errors.New("ReturnURL should not exceed 200")
	}
	if len(o.ClientBackURL) > 200 {
		return false, errors.New("ClientBackURL should not exceed 200")
	}
	if len(o.ItemURL) > 200 {
		return false, errors.New("ItemURL should not exceed 200")
	}
	if len(o.OrderResultURL) > 200 {
		return false, errors.New("OrderResultURL should not exceed 200")
	}

	if ci := o.Invoice.CustomerIdentifier; ci != "" {
		if len(ci) != 8 {
			return false, errors.New("CustomerIdentifier has to fill fixed length of 8 digits")
		}
		if o.Invoice.CarruerType == "" {
			return false, errors.New("CarruerType does not fill any value, when CustomerIdentifier have value")
		}
		if !o.Invoice.Print {
			return false, errors.New("Print has to be true, when CustomerIdentifier have value")
		}
		if o.Invoice.Donation {
			return false, errors.New("Donation has to be false, when CustomerIdentifier have value")
		}
	}

	if o.Invoice.Print {
		if o.Invoice.CustomerName == "" {
			return false, errors.New("CustomerName should not be empty if Print is true")
		}
		if o.Invoice.CustomerAddr == "" {
			return false, errors.New("CustomerAddr should not be empty if Print is true")
		}
		if o.Invoice.CarruerType == "" {
			return false, errors.New("CarruerType should not be empty if Print is true")
		}
	}

	if o.Invoice.CustomerEmail == "" && o.Invoice.CustomerPhone == "" {
		return false, errors.New("CustomerPhone should not be empty if CustomerEmail is empty")
	}

	if o.Invoice.Donation {
		if o.Invoice.Print {
			return false, errors.New("Print should be false if Donation is set to true")
		}
		if lc := o.Invoice.LoveCode; lc == "" {
			return false, errors.New("LoveCode should not be empty if Donation is set to true")
		} else if len(lc) < 3 || len(lc) > 7 {
			return false, errors.New("LoveCode should be a 3-7 digit number")
		}
	}

	return true, nil
}

// ToFormData transform the Order struct to url.Values
func (o *Order) ToFormData(merchantID string) url.Values {
	ecpayReq := map[string][]string{}
	ecpayReq["MerchantID"] = []string{merchantID}
	ecpayReq["ChoosePayment"] = []string{string(o.ChoosePayment)}
	ecpayReq["EncryptType"] = []string{"1"}
	ecpayReq["MerchantTradeNo"] = []string{
		fmt.Sprintf("%s%d", uuid.New().String()[:5], time.Now().Unix()),
	}
	ecpayReq["MerchantTradeDate"] = []string{
		time.Now().Format("2006/01/02 15:04:05"),
	}
	ecpayReq["PaymentType"] = []string{string(o.PaymentType)}
	ecpayReq["TotalAmount"] = []string{strconv.Itoa(o.TotalAmount)}
	ecpayReq["TradeDesc"] = []string{o.TradeDesc}
	ecpayReq["ItemName"] = []string{strings.Join(o.ItemNames, "#")}
	ecpayReq["ReturnURL"] = []string{o.ReturnURL}
	if o.NeedExtraPaidInfo {
		ecpayReq["NeedExtraPaidInfo"] = []string{"Y"}
	} else {
		ecpayReq["NeedExtraPaidInfo"] = []string{"N"}
	}

	ecpayReq["CheckMacValue"] = []string{
		getCheckMacValue(ecpayReq),
	}
	return ecpayReq
}

func getCheckMacValue(req url.Values) string {
	keys := []string{}
	for k := range req {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	str := "HashKey=5294y06JbISpM5x9&"
	for _, k := range keys {
		if req[k][0] != "" {
			str += k + "=" + req[k][0] + "&"
		}
	}

	str += "HashIV=v77hoKGq4kWxNNIS"
	str = url.QueryEscape(str)
	str = strings.ReplaceAll(str, "%2A", "*")
	str = strings.ReplaceAll(str, "%28", "(")
	str = strings.ReplaceAll(str, "%29", ")")
	str = strings.ReplaceAll(str, "%21", "!")
	str = strings.ToLower(str)
	str = fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
	str = strings.ToUpper(str)

	return str
}
