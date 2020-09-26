package ecpay

// ActionType defines the struct of credit card action.
type ActionType string

const (
	// ActionTypeC (關帳).
	ActionTypeC ActionType = "C"
	// ActionTypeR (退刷).
	ActionTypeR ActionType = "R"
	// ActionTypeE (取消).
	ActionTypeE ActionType = "E"
	// ActionTypeN (放棄).
	ActionTypeN ActionType = "N"
)

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
	// UnionPayTypeSelect customer choose whether use UnionPay or not in webpage (消費者於交易頁面可選擇是否使用銀聯交易).
	UnionPayTypeSelect UnionPayType = "0"
	// UnionPayTypeOnly use UnionPay (只使用銀聯卡交易,且綠界會將交易頁面直接導到銀聯網站).
	UnionPayTypeOnly UnionPayType = "1"
	// UnionPayTypeHidden do not use UnionPay (不可使用銀聯卡,綠界會將交易頁面隱藏銀聯選項).
	UnionPayTypeHidden UnionPayType = "2"
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
	// ChooseSubpaymentTypeTaishin Taishin Bank (台新銀行).
	ChooseSubpaymentTypeTaishin ChooseSubpaymentType = "TAISHIN"
	// ChooseSubpaymentTypeESUN E.SUN Bank (玉山銀行).
	ChooseSubpaymentTypeESUN ChooseSubpaymentType = "ESUN"
	// ChooseSubpaymentTypeBOT Bank of Taiwan (台灣銀行).
	ChooseSubpaymentTypeBOT ChooseSubpaymentType = "BOT"
	// ChooseSubpaymentTypeFubon Fubon (富邦銀行).
	ChooseSubpaymentTypeFubon ChooseSubpaymentType = "FUBON"
	// ChooseSubpaymentTypeChinaTrust ChinaTrust Bank (中國信託).
	ChooseSubpaymentTypeChinaTrust ChooseSubpaymentType = "CHINATRUST"
	// ChooseSubpaymentTypeFirst First Commercial Bank (第一銀行).
	ChooseSubpaymentTypeFirst ChooseSubpaymentType = "FIRST"
	// ChooseSubpaymentTypeCathay Cathay Bank (花旗銀行).
	ChooseSubpaymentTypeCathay ChooseSubpaymentType = "CATHAY"
	// ChooseSubpaymentTypeMega Mega International Commercial Bank (兆豐銀行).
	ChooseSubpaymentTypeMega ChooseSubpaymentType = "MEGA"
	// ChooseSubpaymentTypeLand Land Bank (土地銀行).
	ChooseSubpaymentTypeLand ChooseSubpaymentType = "LAND"
	// ChooseSubpaymentTypeTachong (大眾銀行).
	ChooseSubpaymentTypeTachong ChooseSubpaymentType = "TACHONG"
	// ChooseSubpaymentTypeSinoPac SinoPac Bank (永豐銀行).
	ChooseSubpaymentTypeSinoPac ChooseSubpaymentType = "SINOPAC"
	// ChooseSubpaymentTypeCVS convenience store payment code (超商代碼繳款).
	ChooseSubpaymentTypeCVS ChooseSubpaymentType = "CVS"
	// ChooseSubpaymentTypeOK OK Mart.
	ChooseSubpaymentTypeOK ChooseSubpaymentType = "OK"
	// ChooseSubpaymentTypeFamily Family mart (全家).
	ChooseSubpaymentTypeFamily ChooseSubpaymentType = "FAMILY"
	// ChooseSubpaymentTypeHiLife HiFife.
	ChooseSubpaymentTypeHiLife ChooseSubpaymentType = "HILIFE"
	// ChooseSubpaymentTypeIBon 7-11 ibon.
	ChooseSubpaymentTypeIBon ChooseSubpaymentType = "IBON"
	// ChooseSubpaymentTypeBarcode barcode.
	ChooseSubpaymentTypeBarcode ChooseSubpaymentType = "BARCODE"
)

// NeedExtraPaidInfoType defines the struct of extra payment info options (是否需要額外的付款資訊).
type NeedExtraPaidInfoType string

const (
	NeedExtraPaidInfoTypeYes NeedExtraPaidInfoType = "Y"
	NeedExtraPaidInfoTypeNo  NeedExtraPaidInfoType = "N"
)

// InvoiceMarkType defines the struct of invoice options (電子發票開立註記).
type InvoiceMarkType string

const (
	InvoiceMarkTypeYes InvoiceMarkType = "Y"
	InvoiceMarkTypeNo  InvoiceMarkType = "N"
)

// LanguageType defines the struct of language options (語系設定).
type LanguageType string

const (
	LanguageTypeEnglish           LanguageType = "ENG"
	LanguageTypeKorean            LanguageType = "KOR"
	LanguageTypeJapanese          LanguageType = "JPN"
	LanguageTypeSimplifiedChinese LanguageType = "CHI"
)

// BindingCardType defines the struct of binding card options (記憶卡號).
type BindingCardType string

const (
	BindingCardTypeYes BindingCardType = "1"
	BindingCardTypeNo  BindingCardType = "0"
)

// RedeemType defines the struct of redeem options (信用卡是否使用紅利折抵).
type RedeemType string

const (
	RedeemTypeYes RedeemType = "Y"
	RedeemTypeNo  RedeemType = "N"
)

// ClearanceMarkType defines the struct of redeem options (通關方式).
type ClearanceMarkType string

const (
	ClearanceMarkTypeNormal  ClearanceMarkType = "1"
	ClearanceMarkTypeCustoms ClearanceMarkType = "2"
)

// PrintType defines the struct of redeem options (列印註記).
type PrintType string

const (
	PrintTypeYes PrintType = "1"
	PrintTypeNo  PrintType = "0"
)

// PayDateType defines the struct of pay date type options (查詢日期類別).
type PayDateType string

const (
	PayDateTypeFund  PayDateType = "fund"
	PayDateTypeClose PayDateType = "close"
	PayDateTypeEnter PayDateType = "enter"
)

// DateType defines the struct of pay date type options (查詢日期類別).
type DateType string

const (
	// DateTypePayment (付款日期).
	DateTypePayment DateType = "2"
	// DateTypeAllocation (撥款日期).
	DateTypeAllocation DateType = "4"
	// DateTypeOrder (訂單日期).
	DateTypeOrder DateType = "6"
)

// MerchantPaymentType defines the struct of pay date type options (查詢日期類別).
type MerchantPaymentType string

const (
	// MerchantPaymentTypeCreditCard (付款日期).
	MerchantPaymentTypeCreditCard MerchantPaymentType = "01"
	// DateTypeApproriation (網路ATM).
	MerchantPaymentTypeWebATM MerchantPaymentType = "02"
	// MerchantPaymentTypeATM (ATM).
	MerchantPaymentTypeATM MerchantPaymentType = "03"
	// MerchantPaymentTypeCVS (超商代碼).
	MerchantPaymentTypeCVS MerchantPaymentType = "04"
	// MerchantPaymentTypeBarcode (超商條碼).
	MerchantPaymentTypeBarcode MerchantPaymentType = "05"
)

type PlatformStatusType string

const (
	// PlatformStatusTypeAll (全部).
	PlatformStatusTypeAll PlatformStatusType = "0"
	// PlatformStatusTypeNormal (一般).
	PlatformStatusTypeNormal PlatformStatusType = "1"
	// PlatformStatusTypePlatform (平台).
	PlatformStatusTypePlatform PlatformStatusType = "2"
)

type PaymentStatusType string

const (
	// PaymentStatusTypeUnpaid (未付款).
	PaymentStatusTypeUnpaid PaymentStatusType = "0"
	// PaymentStatusTypePaid (已付款).
	PaymentStatusTypePaid PaymentStatusType = "1"
	// PaymentStatusTypeFailed (訂單失敗).
	PaymentStatusTypeFailed PaymentStatusType = "2"
)

type AllocateStatusType string

const (
	// AllocateStatusTypeDone (已撥款).
	AllocateStatusTypeDone AllocateStatusType = "0"
	// AllocateStatusNotYet (未撥款).
	AllocateStatusNotYet AllocateStatusType = "1"
)

type MediaFormatedType string

const (
	// MediaFormatedTypeOld (舊版格式).
	MediaFormatedTypeOld MediaFormatedType = "0"
	// MediaFormatedTypeNew (新版格式).
	MediaFormatedTypeNew MediaFormatedType = "1"
)

type CharSetType string

const (
	CharSetTypeBig5 CharSetType = "1"
	CharSetTypeUTF8 CharSetType = "2"
)
