package ecpay

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type PeriodType string

const (
	YEAR  PeriodType = "Y"
	MONTH PeriodType = "M"
	DAY   PeriodType = "D"
)

type CarrierType string

const (
	NONE      CarrierType = ""
	MEMBER    CarrierType = "1"
	CITIZEN   CarrierType = "2"
	CELLPHONE CarrierType = "3"
)

type PaymentType string

const (
	AIO               PaymentType = "aio"
	WebATM_TAISHIN    PaymentType = "WebATM_TAISHIN"
	WebATM_BOT        PaymentType = "WebATM_BOT"
	WebATM_CHINATRUST PaymentType = "WebATM_CHINATRUST"
	WebATM_CATHAY     PaymentType = "WebATM_CATHAY"
	WebATM_LAND       PaymentType = "WebATM_LAND"
	WebATM_SINOPAC    PaymentType = "WebATM_SINOPAC"
	ATM_ESUN          PaymentType = "ATM_ESUN"
	ATM_FUBON         PaymentType = "ATM_FUBON"
	ATM_FIRST         PaymentType = "ATM_FIRST"
	ATM_CATHAY        PaymentType = "ATM_CATHAY"
	CVS_CVS           PaymentType = "CVS_CVS"
	CVS_FAMILY        PaymentType = "CVS_FAMILY"
	CVS_IBON          PaymentType = "CVS_IBON"
	Credit_CreditCard PaymentType = "Credit_CreditCard"
)

type ChoosePayment string

const (
	ALL       ChoosePayment = "ALL"
	Credit    ChoosePayment = "Credit"
	WebATM    ChoosePayment = "WebATM"
	ATM       ChoosePayment = "ATM"
	CVS       ChoosePayment = "CVS"
	BARCODE   ChoosePayment = "BARCODE"
	GooglePay ChoosePayment = "GooglePay"
)

// Order .
type Order struct {
	MerchantTradeNo   string
	StoreID           string
	MerchantTradeDate string
	PaymentType       PaymentType
	TotalAmount       int
	TradeDesc         string
	ItemNames         []string
	ReturnURL         string
	ChoosePayment     ChoosePayment
	ClientBackURL     string
	ItemURL           string
	Remark            string
	ChooseSubPayment  string
	OrderResultURL    string
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

	// invoice
	RelateNumber       string
	CustomerID         string
	CustomerIdentifier string
	CustomerName       string
	CustomerAddr       string
	CustomerPhone      string
	CustomerEmail      string
	ClearanceMark      string
	TaxType            string
	CarruerType        string
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

type ExtendParams1 struct {
	ExpireDate        string
	PaymentInfoURL    string
	ClientRedirectURL string
}

type ExtendParams2 struct {
	StoreExpireDate   string
	Desc1             string
	Desc2             string
	Desc3             string
	Desc4             string
	PaymentInfoURL    string
	ClientRedirectURL string
}

type ExtendParams3 struct {
	BindingCard      string
	MerchantMemberID string
}

type ExtendParams4 struct {
	Redeem   string
	UnionPay string
}

// Validate validate if the struct is valid.
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

	if ci := o.CustomerIdentifier; ci != "" {
		if len(ci) != 8 {
			return false, errors.New("CustomerIdentifier has to fill fixed length of 8 digits")
		}
		if o.CarruerType == "" {
			return false, errors.New("CarruerType does not fill any value, when CustomerIdentifier have value")
		}
		if !o.Print {
			return false, errors.New("Print has to be true, when CustomerIdentifier have value")
		}
		if o.Donation {
			return false, errors.New("Donation has to be false, when CustomerIdentifier have value")
		}
	}

	if o.Print {
		if o.CustomerName == "" {
			return false, errors.New("CustomerName should not be empty if Print is true")
		}
		if o.CustomerAddr == "" {
			return false, errors.New("CustomerAddr should not be empty if Print is true")
		}
		if o.CarruerType == "" {
			return false, errors.New("CarruerType should not be empty if Print is true")
		}
	}

	if o.CustomerEmail == "" && o.CustomerPhone == "" {
		return false, errors.New("CustomerPhone should not be empty if CustomerEmail is empty")
	}

	if o.Donation {
		if o.Print {
			return false, errors.New("Print should be false if Donation is set to true")
		}
		if lc := o.LoveCode; lc == "" {
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
	str := "HashKey=5294y06JbISpM5x9&"
	if req["ChoosePayment"][0] != "" {
		str += "ChoosePayment=" + req["ChoosePayment"][0] + "&"
	}
	if req["EncryptType"][0] != "" {
		str += "EncryptType=" + req["EncryptType"][0] + "&"
	}
	if req["ItemName"][0] != "" {
		str += "ItemName=" + req["ItemName"][0] + "&"
	}
	if req["MerchantID"][0] != "" {
		str += "MerchantID=" + req["MerchantID"][0] + "&"
	}
	if req["MerchantTradeDate"][0] != "" {
		str += "MerchantTradeDate=" + req["MerchantTradeDate"][0] + "&"
	}
	if req["MerchantTradeNo"][0] != "" {
		str += "MerchantTradeNo=" + req["MerchantTradeNo"][0] + "&"
	}
	if req["NeedExtraPaidInfo"][0] != "" {
		str += "NeedExtraPaidInfo=" + req["NeedExtraPaidInfo"][0] + "&"
	}
	if req["PaymentType"][0] != "" {
		str += "PaymentType=" + req["PaymentType"][0] + "&"
	}
	if req["ReturnURL"][0] != "" {
		str += "ReturnURL=" + req["ReturnURL"][0] + "&"
	}
	if req["TotalAmount"][0] != "" {
		str += "TotalAmount=" + req["TotalAmount"][0] + "&"
	}
	if req["TradeDesc"][0] != "" {
		str += "TradeDesc=" + req["TradeDesc"][0] + "&"
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
