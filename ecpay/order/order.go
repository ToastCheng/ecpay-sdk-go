package order

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// Order defines the structure of an order.
type Order struct {
	// MerchantTradeNo (特店交易編號).
	MerchantTradeNo string `json:"MerchantTradeNo,omitempty"`
	// StoreID (特店旗下店舖代號).
	StoreID string `json:"StoreID,omitempty"`
	// MerchantTradeDate yyyy/MM/dd HH:mm:ss (特店交易時間).
	MerchantTradeDate string `json:"MerchantTradeDate,omitempty"`
	// TotalAmount (交易金額).
	TotalAmount int `json:"TotalAmount,omitempty"`
	// TradeDesc (交易描述).
	TradeDesc string `json:"TradeDesc,omitempty"`
	// ItemName (商品名稱).
	ItemName string `json:"ItemName,omitempty"`
	// ItemURL (商品銷售網址).
	ItemURL string `json:"ItemURL,omitempty"`
	// ReturnURL (付款完成通知回傳網址).
	ReturnURL string `json:"ReturnURL,omitempty"`
	// ClientBackURL (Client端返回特店的按鈕連結).
	ClientBackURL string `json:"ClientBackURL,omitempty"`
	// OrderResultURL (Client端回傳付款結果網址).
	OrderResultURL string `json:"OrderResultURL,omitempty"`
	// PaymentType (交易類型).
	PaymentType PaymentType `json:"PaymentType,omitempty"`
	// ChoosePayment (選擇預設付款方式).
	ChoosePayment ChoosePaymentType `json:"ChoosePayment,omitempty"`
	// ChooseSubPayment (付款子項目).
	ChooseSubPayment ChooseSubpaymentType `json:"ChooseSubPayment,omitempty"`
	// Remark (備註欄位).
	Remark string `json:"Remark,omitempty"`
	// NeedExtraPaidInfo (是否需要額外的付款資訊).
	NeedExtraPaidInfo bool `json:"NeedExtraPaidInfo,omitempty"`
	// DeviceSource (裝置來源).
	deviceSource string `json:"DeviceSource"`
	// IgnorePayment (隱藏付款).
	IgnorePayment string `json:"IgnorePayment,omitempty"`
	// PlatformID (特約合作平台商代號).
	PlatformID string `json:"PlatformID,omitempty"`
	// InvoiceMark (電子發票開立註記).
	InvoiceMark bool `json:"InvoiceMark,omitempty"`
	// CustomField1 (自訂名稱欄位1).
	CustomField1 string `json:"CustomField1,omitempty"`
	// CustomField2 (自訂名稱欄位2).
	CustomField2 string `json:"CustomField2,omitempty"`
	// CustomField3 (自訂名稱欄位3).
	CustomField3 string `json:"CustomField3,omitempty"`
	// CustomField4 (自訂名稱欄位4).
	CustomField4 string `json:"CustomField4,omitempty"`
	// EncryptType (CheckMacValue 加密類型).
	EncryptType string `json:"EncryptType,omitempty"`
	// Language (語系設定).
	Language string `json:"Language,omitempty"`

	ATM        *ATMParam          `json:"ATM,omitempty"`
	CVSBarcode *CVSOrBarcodeParam `json:"CVSBarcode,omitempty"`
	Credit     *CreditParam       `json:Credit",omitempty"`
	Invoice    *InvoiceParam      `json:Invoice",omitempty"`
}

// ATMParam defines the parameters tailored for ATM transaction.
type ATMParam struct {
	// ExpireDate (允許繳費有效天數).
	ExpireDate string
	// PaymentInfoURL (Server端回傳付款相關資訊).
	PaymentInfoURL string
	// ClientRedirectURL (Client端回傳付款相關資訊).
	ClientRedirectURL string
}

// CVSOrBarcodeParam defines the parameters tailored for CVS or bar code transaction.
type CVSOrBarcodeParam struct {
	// StoreExpireDate (超商繳費截止時間).
	StoreExpireDate string
	// Desc1(交易描述1).
	Desc1 string
	// Desc2 (交易描述2).
	Desc2 string
	// Desc3 (交易描述3).
	Desc3 string
	// Desc4 (交易描述4).
	Desc4 string
	// PaymentInfoURL (Server端回傳付款相關資訊).
	PaymentInfoURL string
	// ClientRedirectURL (Client端回傳付款方式相關資訊).
	ClientRedirectURL string
}

// CreditParam defines the parameters tailored for credit card transaction.
type CreditParam struct {
	// BindingCard (記憶卡號).
	BindingCard string
	// MerchantMemberID (記憶卡號識別碼).
	MerchantMemberID string
	// Language (語系設定).
	Language string

	// 一次付清
	// Redeem (信用卡是否使用紅利折抵).
	Redeem bool
	// UnionPay (銀聯卡交易選項).
	UnionPay UnionPayType

	// 分期付款
	// CreditInstallment (刷卡分期期數).
	CreditInstallment string

	// 定期定額
	PeriodAmount int
	// PeriodType (週期種類).
	PeriodType PeriodType
	// Frequency (執行頻率).
	Frequency int
	// ExecTimes (執行次數).
	ExecTimes int
	// PeriodReturnURL (定期定額的執行結果回應URL).
	PeriodReturnURL string
}

// InvoiceParam defines the parameters for invoice specific settings.
type InvoiceParam struct {
	// RelateNumber (特店自訂編號).
	RelateNumber string
	// TaxType (課稅類別).
	TaxType TaxType
	// Donation (捐贈註記).
	Donation DonationType
	// Print (列印註記).
	Print bool
	// InvoiceItemName (商品名稱).
	InvoiceItemName string
	// InvoiceItemCount (商品數量).
	InvoiceItemCount string
	// DelayDay (延遲天數).
	DelayDay string
	// InvType (字軌類別).
	InvType string

	// CustomerID (客戶編號).
	CustomerID string
	// CustomerIdentifier (統一編號).
	CustomerIdentifier string
	// CustomerName (客戶名稱).
	CustomerName string
	// CustomerAddr (客戶地址).
	CustomerAddr string
	// CustomerPhone (客戶手機號碼).
	CustomerPhone string
	// CustomerEmail (客戶電子信箱).
	CustomerEmail string
	// ClearanceMark (通關方式).
	ClearanceMark string
	// CarrierType (載具類別).
	CarrierType CarrierType
	// CarrierNum (載具編號).
	CarrierNum string
	// LoveCode (捐贈碼).
	LoveCode string
	// InvoiceItemWord (商品單位).
	InvoiceItemWord string
	// InvoiceItemPrice (商品價格).
	InvoiceItemPrice string
	// InvoiceItemTaxType (商品課稅別).
	InvoiceItemTaxType string
	// InvoiceRemark (備註).
	InvoiceRemark string
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
	if len(o.ItemName) == 0 {
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
	if len(o.ItemName) > 200 {
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
func (o Order) ToFormData() url.Values {
	ecpayReq := map[string][]string{}
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
	ecpayReq["ItemName"] = []string{o.ItemName}
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

	return ecpayReq
}
