package ecpay

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
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
	NeedExtraPaidInfo NeedExtraPaidInfoType `json:"NeedExtraPaidInfo,omitempty"`
	// DeviceSource (裝置來源).
	DeviceSource string `json:"DeviceSource,omitempty"`
	// IgnorePayment (隱藏付款).
	IgnorePayment string `json:"IgnorePayment,omitempty"`
	// PlatformID (特約合作平台商代號).
	PlatformID string `json:"PlatformID,omitempty"`
	// InvoiceMark (電子發票開立註記).
	InvoiceMark InvoiceMarkType `json:"InvoiceMark,omitempty"`
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
	Language LanguageType `json:"Language,omitempty"`

	ATM        *ATMParam          `json:"ATM,omitempty"`
	CVSBarcode *CVSOrBarcodeParam `json:"CVSBarcode,omitempty"`
	Credit     *CreditParam       `json:"Credit,omitempty"`
	Invoice    *InvoiceParam      `json:"Invoice,omitempty"`
}

// ATMParam defines the parameters tailored for ATM transaction.
type ATMParam struct {
	// ExpireDate (允許繳費有效天數).
	ExpireDate int `json:"ExpireDate,omitempty"`
	// PaymentInfoURL (Server端回傳付款相關資訊).
	PaymentInfoURL string `json:"PaymentInfoURL,omitempty"`
	// ClientRedirectURL (Client端回傳付款相關資訊).
	ClientRedirectURL string `json:"ClientRedirectURL,omitempty"`
}

// CVSOrBarcodeParam defines the parameters tailored for CVS or bar code transaction.
type CVSOrBarcodeParam struct {
	// StoreExpireDate (超商繳費截止時間).
	// unit: CVS (min), Barcode (day)
	StoreExpireDate int `json:"StoreExpireDate,omitempty"`
	// Desc1 (交易描述1).
	Desc1 string `json:"Desc1,omitempty"`
	// Desc2 (交易描述2).
	Desc2 string `json:"Desc2,omitempty"`
	// Desc3 (交易描述3).
	Desc3 string `json:"Desc3,omitempty"`
	// Desc4 (交易描述4).
	Desc4 string `json:"Desc4,omitempty"`
	// PaymentInfoURL (Server端回傳付款相關資訊).
	PaymentInfoURL string `json:"PaymentInfoURL,omitempty"`
	// ClientRedirectURL (Client端回傳付款方式相關資訊).
	ClientRedirectURL string `json:"ClientRedirectURL,omitempty"`
}

// CreditParam defines the parameters tailored for credit card transaction.
type CreditParam struct {
	// BindingCard (記憶卡號).
	BindingCard BindingCardType `json:"BindingCard,omitempty"`
	// MerchantMemberID (記憶卡號識別碼).
	MerchantMemberID string `json:"MerchantMemberID,omitempty"`
	// Language (語系設定).
	Language string `json:"Language,omitempty"`

	// 一次付清
	// Redeem (信用卡是否使用紅利折抵).
	Redeem RedeemType `json:"Redeem,omitempty"`
	// UnionPay (銀聯卡交易選項).
	UnionPay UnionPayType `json:"UnionPay,omitempty"`

	// 分期付款
	// CreditInstallment (刷卡分期期數).
	CreditInstallment string `json:"CreditInstallment,omitempty"`

	// 定期定額
	PeriodAmount int `json:"PeriodAmount,omitempty"`
	// PeriodType (週期種類).
	PeriodType PeriodType `json:"PeriodType,omitempty"`
	// Frequency (執行頻率).
	Frequency int `json:"Frequency,omitempty"`
	// ExecTimes (執行次數).
	ExecTimes int `json:"ExecTimes,omitempty"`
	// PeriodReturnURL (定期定額的執行結果回應URL).
	PeriodReturnURL string `json:"PeriodReturnURL,omitempty"`
}

// InvoiceParam defines the parameters for invoice specific settings.
type InvoiceParam struct {
	// RelateNumber (特店自訂編號).
	RelateNumber string `json:"RelateNumber,omitempty"`
	// TaxType (課稅類別).
	TaxType TaxType `json:"TaxType,omitempty"`
	// Donation (捐贈註記).
	Donation DonationType `json:"Donation,omitempty"`
	// Print (列印註記).
	Print PrintType `json:"Print,omitempty"`
	// InvoiceItemName (商品名稱).
	InvoiceItemName string `json:"InvoiceItemName,omitempty"`
	// InvoiceItemCount (商品數量).
	InvoiceItemCount string `json:"InvoiceItemCount,omitempty"`
	// DelayDay (延遲天數).
	DelayDay int `json:"DelayDay,omitempty"`
	// InvType (字軌類別).
	InvType string `json:"InvType,omitempty"`

	// CustomerID (客戶編號).
	CustomerID string `json:"CustomerID,omitempty"`
	// CustomerIdentifier (統一編號).
	CustomerIdentifier string `json:"CustomerIdentifier,omitempty"`
	// CustomerName (客戶名稱).
	CustomerName string `json:"CustomerName,omitempty"`
	// CustomerAddr (客戶地址).
	CustomerAddr string `json:"CustomerAddr,omitempty"`
	// CustomerPhone (客戶手機號碼).
	CustomerPhone string `json:"CustomerPhone,omitempty"`
	// CustomerEmail (客戶電子信箱).
	CustomerEmail string `json:"CustomerEmail,omitempty"`
	// ClearanceMark (通關方式).
	ClearanceMark ClearanceMarkType `json:"ClearanceMark,omitempty"`
	// CarrierType (載具類別).
	CarrierType CarrierType `json:"CarruerType,omitempty"`
	// CarrierNum (載具編號).
	CarrierNum string `json:"CarruerNum,omitempty"`
	// LoveCode (捐贈碼).
	LoveCode string `json:"LoveCode,omitempty"`
	// InvoiceItemWord (商品單位).
	InvoiceItemWord string `json:"InvoiceItemWord,omitempty"`
	// InvoiceItemPrice (商品價格).
	InvoiceItemPrice string `json:"InvoiceItemPrice,omitempty"`
	// InvoiceItemTaxType (商品課稅別).
	InvoiceItemTaxType string `json:"InvoiceItemTaxType,omitempty"`
	// InvoiceRemark (備註).
	InvoiceRemark string `json:"InvoiceRemark,omitempty"`
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
	if o.ChoosePayment == "" {
		return false, errors.New("ChoosePayment should not be empty")
	}
	if o.TotalAmount == 0 {
		return false, errors.New("TotalAmount should not be empty")
	}
	if o.PaymentType == "" {
		return false, errors.New("PaymentType should not be empty")
	}
	if len(o.ItemName) == 0 || len(o.ItemName) > 400 {
		return false, errors.New("ItemName should not be empty and within 400 words")
	}
	if o.TradeDesc == "" {
		return false, errors.New("TradeDesc should not be empty")
	}
	if o.ReturnURL == "" {
		return false, errors.New("ReturnURL should not be empty")
	}

	// check length.
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
	if len(o.CustomField1) > 50 ||
		len(o.CustomField2) > 50 ||
		len(o.CustomField3) > 50 ||
		len(o.CustomField4) > 50 {
		return false, errors.New("CustomField should not exceed 50")
	}

	// atm.
	if atm := o.ATM; atm != nil {
		if atm.ExpireDate > 60 || atm.ExpireDate < 1 {
			return false, errors.New("ExpireDate should be in the range of 1-60")
		}
	}

	// cvs.
	if cvsbar := o.CVSBarcode; cvsbar != nil {
		if len(cvsbar.Desc1) > 20 ||
			len(cvsbar.Desc2) > 20 ||
			len(cvsbar.Desc3) > 20 ||
			len(cvsbar.Desc4) > 20 {
			return false, errors.New("Desc should not exceed 20")
		}
	}

	// invoice.
	if inv := o.Invoice; inv != nil {
		if ci := inv.CustomerIdentifier; ci != "" {
			if len(ci) != 8 {
				return false, errors.New("CustomerIdentifier has to fill fixed length of 8 digits")
			}
			if inv.Print == PrintTypeNo {
				return false, errors.New("Print has to be true, when CustomerIdentifier have value")
			}
			if inv.Donation == DonationTypeYes {
				return false, errors.New("Donation has to be false, when CustomerIdentifier have value")
			}
		}

		if inv.Print == PrintTypeYes {
			if inv.CustomerName == "" {
				return false, errors.New("CustomerName should not be empty if Print is true")
			}
			if inv.CustomerAddr == "" {
				return false, errors.New("CustomerAddr should not be empty if Print is true")
			}
		}

		if inv.CustomerEmail == "" && inv.CustomerPhone == "" {
			return false, errors.New("CustomerPhone should not be empty if CustomerEmail is empty")
		}

		if inv.Donation == DonationTypeYes {
			if inv.Print == PrintTypeYes {
				return false, errors.New("Print should be false if Donation is set to true")
			}
			if lc := inv.LoveCode; lc == "" {
				return false, errors.New("LoveCode should not be empty if Donation is set to true")
			} else if len(lc) < 3 || len(lc) > 7 {
				return false, errors.New("LoveCode should be a 3-7 digit number")
			}
		}
	}
	return true, nil
}

// ToFormData transform the Order struct to url.Values
func (o Order) ToFormData() url.Values {
	req := url.Values{}
	mp := map[string]interface{}{}
	databytes, _ := json.Marshal(o)
	json.Unmarshal(databytes, &mp)
	for k, v := range mp {
		if k == "Invoice" ||
			k == "ATM" ||
			k == "CVSBarcode" ||
			k == "Credit" {
			for kk, vv := range v.(map[string]interface{}) {
				switch t := vv.(type) {
				case float32, float64:
					req.Set(kk, fmt.Sprintf("%.0f", t))
				case string:
					req.Set(kk, t)
				}
			}
		} else {
			switch t := v.(type) {
			case float32, float64:
				req.Set(k, fmt.Sprintf("%.0f", t))
			case string:
				req.Set(k, t)
			}
		}
	}
	req.Set("EncryptType", "1")
	return req
}
