package order

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/toastcheng/ecpay-sdk-go/ecpay/utils"

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
	PaymentType       PaymentType
	ChoosePayment     ChoosePaymentType
	ChooseSubPayment  ChooseSubpaymentType
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

	// 一次付清
	Redeem   bool
	UnionPay UnionPayType

	// 分期付款
	CreditInstallment string

	// 定期定額
	PeriodAmount    int
	PeriodType      PeriodType
	Frequency       int
	ExecTimes       int
	PeriodReturnURL string
}

// InvoiceParam defines the parameters for invoice specific settings.
type InvoiceParam struct {
	RelateNumber     string
	TaxType          TaxType
	Donation         DonationType
	Print            bool
	InvoiceItemName  string
	InvoiceItemCount string
	DelayDay         string
	InvType          string

	CustomerID         string
	CustomerIdentifier string
	CustomerName       string
	CustomerAddr       string
	CustomerPhone      string
	CustomerEmail      string
	ClearanceMark      string
	CarrierType        CarrierType
	CarrierNum         string
	LoveCode           string
	InvoiceItemWord    string
	InvoiceItemPrice   string
	InvoiceItemTaxType string
	InvoiceRemark      string
}

// Validate validate if the order struct is valid.
func (o Order) Validate() (bool, error) {
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
		if o.Invoice.CarrierType == "" {
			return false, errors.New("CarruerType does not fill any value, when CustomerIdentifier have value")
		}
		if !o.Invoice.Print {
			return false, errors.New("Print has to be true, when CustomerIdentifier have value")
		}
		if o.Invoice.Donation == DonationTypeNo {
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
		if o.Invoice.CarrierType == "" {
			return false, errors.New("CarruerType should not be empty if Print is true")
		}
	}

	if o.Invoice.CustomerEmail == "" && o.Invoice.CustomerPhone == "" {
		return false, errors.New("CustomerPhone should not be empty if CustomerEmail is empty")
	}

	if o.Invoice.Donation == DonationTypeYes {
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
func (o Order) ToFormData(merchantID string) url.Values {
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
	ecpayReq["ReturnURL"] = []string{o.ReturnURL}
	ecpayReq["ItemName"] = []string{strings.Join(o.ItemNames, "#")}
	if o.NeedExtraPaidInfo {
		ecpayReq["NeedExtraPaidInfo"] = []string{"Y"}
	} else {
		ecpayReq["NeedExtraPaidInfo"] = []string{"N"}
	}
	if o.InvoiceMark {
		ecpayReq["InvoiceMark"] = []string{"Y"}
	} else {
		ecpayReq["InvoiceMark"] = []string{"N"}
	}

	cp := o.ChoosePayment
	if cp == ChoosePaymentTypeAll || cp == ChoosePaymentTypeCredit {
		// 一次付清
		if o.Credit.Redeem {
			if o.Credit.Redeem {
				ecpayReq["Redeem"] = []string{"Y"}
			} else {
				ecpayReq["Redeem"] = []string{"N"}
			}
			ecpayReq["UnionPay"] = []string{string(o.Credit.UnionPay)}

		} else if o.Credit.CreditInstallment != "" {
			// 分期付款

		} else if o.Credit.PeriodAmount != 0 ||
			// 定期定額
			o.Credit.PeriodType == PeriodTypeDay ||
			o.Credit.Frequency != 0 ||
			o.Credit.ExecTimes != 0 ||
			o.Credit.PeriodReturnURL != "" {
		}
	}

	if o.InvoiceMark {
		ecpayReq["RelateNumber"] = []string{o.Invoice.RelateNumber}
		ecpayReq["TaxType"] = []string{string(o.Invoice.TaxType)}
		ecpayReq["DelayDay"] = []string{o.Invoice.DelayDay}
		ecpayReq["InvType"] = []string{o.Invoice.InvType}
		ecpayReq["Donation"] = []string{string(o.Invoice.Donation)}
		if o.Invoice.Print {
			ecpayReq["Print"] = []string{"1"}
		} else {
			ecpayReq["Print"] = []string{"0"}
		}

		ecpayReq["CustomerID"] = []string{o.Invoice.CustomerID}
		ecpayReq["CustomerIdentifier"] = []string{o.Invoice.CustomerIdentifier}
		ecpayReq["CustomerName"] = []string{o.Invoice.CustomerName}
		ecpayReq["CustomerAddr"] = []string{o.Invoice.CustomerAddr}
		ecpayReq["CustomerPhone"] = []string{o.Invoice.CustomerPhone}
		ecpayReq["CustomerEmail"] = []string{o.Invoice.CustomerEmail}
		ecpayReq["ClearanceMark"] = []string{o.Invoice.ClearanceMark}
		ecpayReq["LoveCode"] = []string{o.Invoice.LoveCode}
		ecpayReq["InvoiceItemTaxType"] = []string{string(o.Invoice.InvoiceItemTaxType)}
		ecpayReq["InvoiceRemark"] = []string{o.Invoice.InvoiceRemark}

		// might be a typo in ECPay server.
		ecpayReq["CarruerType"] = []string{string(o.Invoice.CarrierType)}
		ecpayReq["CarruerNum"] = []string{o.Invoice.CarrierNum}
		// TODO: use '|' to seperate multiple items.
		ecpayReq["InvoiceItemName"] = []string{o.Invoice.InvoiceItemName}
		ecpayReq["InvoiceItemCount"] = []string{o.Invoice.InvoiceItemCount}
		ecpayReq["InvoiceItemWord"] = []string{o.Invoice.InvoiceItemWord}
		ecpayReq["InvoiceItemPrice"] = []string{o.Invoice.InvoiceItemPrice}
	}

	ecpayReq["CheckMacValue"] = []string{utils.GetCheckMacValue(ecpayReq)}

	return ecpayReq
}
