package order

// CarrierType defines the struct of carrier options(載具類別).
type CarrierType string

const (
	// CarrierTypeNone none (無載具).
	CarrierTypeNone CarrierType = ""
	// CarrierTypeMember member (特店載具).
	CarrierTypeMember CarrierType = "1"
	// CarrierTypeCitizen citizen (買受人之自然人憑證號碼).
	CarrierTypeCitizen CarrierType = "2"
	// CarrierTypeCellphone phone (買受人之手機條碼資料).
	CarrierTypeCellphone CarrierType = "3"
)

// DonationType defines the struct of donation options (捐贈).
type DonationType string

const (
	// DonationTypeYes yes (捐贈).
	DonationTypeYes DonationType = "1"
	// DonationTypeNo no (不捐贈或統一編號CustomerIdentifier有值).
	DonationTypeNo DonationType = "2"
)

// InvoiceType defines the struct of invoice options (發票).
type InvoiceType string

const (
	// InvoiceTypeGeneral 一般稅額
	InvoiceTypeGeneral InvoiceType = "07"
	// InvoiceTypeSpecial 特種稅額
	InvoiceTypeSpecial InvoiceType = "08"
)

// PeriodType defines the struct of period options (週期種類).
type PeriodType string

const (
	// PeriodTypeYear year.
	PeriodTypeYear PeriodType = "Y"
	// PeriodTypeMonth month.
	PeriodTypeMonth PeriodType = "M"
	// PeriodTypeDay day.
	PeriodTypeDay PeriodType = "D"
)

// TaxType defines the struct of tax type options (課稅類別).
type TaxType string

const (
	// TaxTypeDutiable (應稅).
	TaxTypeDutiable TaxType = "1"
	// TaxTypeZero (零稅率).
	TaxTypeZero TaxType = "2"
	// TaxTypeFree (免稅).
	TaxTypeFree TaxType = "3"
	// TaxTypeMix (若為混合應稅與免稅或零稅率時,限收銀機發票無法分辨時使用,且需通過申請核可).
	TaxTypeMix TaxType = "9"
)

// UnionPayType defines the struct of UnionPay options (銀聯卡交易選項).
type UnionPayType string

const (
	// UnionPayTypeChooseInWebPage customer choose whether use UnionPay or not in webpage (消費者於交易頁面可選擇是否使用銀聯交易).
	UnionPayTypeChooseInWebPage UnionPayType = "0"
	// UnionPayTypeYes use UnionPay (只使用銀聯卡交易,且綠界會將交易頁面直接導到銀聯網站).
	UnionPayTypeYes UnionPayType = "1"
	// UnionPayTypeNo do not use UnionPay (不可使用銀聯卡,綠界會將交易頁面隱藏銀聯選項).
	UnionPayTypeNo UnionPayType = "2"
)

// PaymentType defines the struct of payment type options (銀聯卡交易選項).
type PaymentType string

const (
	// PaymentTypeAIO all in one.
	PaymentTypeAIO PaymentType = "aio"
	// PaymentTypeWebATMTaishin Taishin Bank (台新銀行).
	PaymentTypeWebATMTaishin PaymentType = "WebATM_TAISHIN"
	// PaymentTypeWebATMBOT Bank of Taiwan (台灣銀行).
	PaymentTypeWebATMBOT PaymentType = "WebATM_BOT"
	// PaymentTypeWebATMChinaTrust ChinaTrust Bank (中國信託).
	PaymentTypeWebATMChinaTrust PaymentType = "WebATM_CHINATRUST"
	// PaymentTypeWebATMCathay Cathay Bank (花旗銀行).
	PaymentTypeWebATMCathay PaymentType = "WebATM_CATHAY"
	// PaymentTypeWebATMLand Land Bank (土地銀行).
	PaymentTypeWebATMLand PaymentType = "WebATM_LAND"
	// PaymentTypeWebATMSinoPac SinoPac Bank (永豐銀行).
	PaymentTypeWebATMSinoPac PaymentType = "WebATM_SINOPAC"
	// PaymentTypeATMESUN E.SUN Bank (玉山銀行).
	PaymentTypeATMESUN PaymentType = "ATM_ESUN"
	// PaymentTypeATMFubon Fubon (富邦銀行).
	PaymentTypeATMFubon PaymentType = "ATM_FUBON"
	// PaymentTypeATMFirst First Commercial Bank (第一銀行).
	PaymentTypeATMFirst PaymentType = "ATM_FIRST"
	// PaymentTypeATMCathay Cathay Bank (花旗銀行).
	PaymentTypeATMCathay PaymentType = "ATM_CATHAY"
	// PaymentTypeCVSCVS convenience store payment code (超商代碼繳款).
	PaymentTypeCVSCVS PaymentType = "CVS_CVS"
	// PaymentTypeCVSFamily Family mart (全家).
	PaymentTypeCVSFamily PaymentType = "CVS_FAMILY"
	// PaymentTypeCVSIbon 7-11 ibon.
	PaymentTypeCVSIbon PaymentType = "CVS_IBON"
	// PaymentTypeCreditCreditCard credit card (信用卡).
	PaymentTypeCreditCreditCard PaymentType = "Credit_CreditCard"
)

// ChoosePaymentType defines the struct of default payment type options (選擇預設付款方式).
type ChoosePaymentType string

const (
	// ChoosePaymentTypeAll all type.
	ChoosePaymentTypeAll ChoosePaymentType = "ALL"
	// ChoosePaymentTypeCredit credit card (信用卡).
	ChoosePaymentTypeCredit ChoosePaymentType = "Credit"
	// ChoosePaymentTypeWebATM web ATM (網路ATM).
	ChoosePaymentTypeWebATM ChoosePaymentType = "WebATM"
	// ChoosePaymentTypeATM ATM.
	ChoosePaymentTypeATM ChoosePaymentType = "ATM"
	// ChoosePaymentTypeCVS convenience store payment code (超商代碼繳款)
	ChoosePaymentTypeCVS ChoosePaymentType = "CVS"
	// ChoosePaymentTypeBarCode bar code (超商條碼繳款).
	ChoosePaymentTypeBarCode ChoosePaymentType = "BARCODE"
	// ChoosePaymentTypeGooglePay Google Pay.
	ChoosePaymentTypeGooglePay ChoosePaymentType = "GooglePay"
)

// ChooseSubpaymentType defines the struct of subpayment type options (付款子項目).
type ChooseSubpaymentType string

const (
	// ChooseSubpaymentTypeTaishin .
	ChooseSubpaymentTypeTaishin ChooseSubpaymentType = "TAISHIN"
	// ChooseSubpaymentTypeESUN .
	ChooseSubpaymentTypeESUN ChooseSubpaymentType = "ESUN"
	// ChooseSubpaymentTypeBOT .
	ChooseSubpaymentTypeBOT ChooseSubpaymentType = "BOT"
	// ChooseSubpaymentTypeFubon .
	ChooseSubpaymentTypeFubon ChooseSubpaymentType = "FUBON"
	// ChooseSubpaymentTypeChinaTrust .
	ChooseSubpaymentTypeChinaTrust ChooseSubpaymentType = "CHINATRUST"
	// ChooseSubpaymentTypeFirst .
	ChooseSubpaymentTypeFirst ChooseSubpaymentType = "FIRST"
	// ChooseSubpaymentTypeCathay .
	ChooseSubpaymentTypeCathay ChooseSubpaymentType = "CATHAY"
	// ChooseSubpaymentTypeMega .
	ChooseSubpaymentTypeMega ChooseSubpaymentType = "MEGA"
	// ChooseSubpaymentTypeLand .
	ChooseSubpaymentTypeLand ChooseSubpaymentType = "LAND"
	// ChooseSubpaymentTypeTachong (大眾銀行).
	ChooseSubpaymentTypeTachong ChooseSubpaymentType = "TACHONG"
	// ChooseSubpaymentTypeSinoPac .
	ChooseSubpaymentTypeSinoPac ChooseSubpaymentType = "SINOPAC"
	// ChooseSubpaymentTypeCVS .
	ChooseSubpaymentTypeCVS ChooseSubpaymentType = "CVS"
	// ChooseSubpaymentTypeOK .
	ChooseSubpaymentTypeOK ChooseSubpaymentType = "OK"
	// ChooseSubpaymentTypeFamily .
	ChooseSubpaymentTypeFamily ChooseSubpaymentType = "FAMILY"
	// ChooseSubpaymentTypeHiLife .
	ChooseSubpaymentTypeHiLife ChooseSubpaymentType = "HILIFE"
	// ChooseSubpaymentTypeIBon .
	ChooseSubpaymentTypeIBon ChooseSubpaymentType = "IBON"
	// ChooseSubpaymentTypeBarcode .
	ChooseSubpaymentTypeBarcode ChooseSubpaymentType = "BARCODE"
)
