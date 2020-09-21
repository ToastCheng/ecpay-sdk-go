package tax

// Tax tax type.
type Tax string

const (
	// Taxable 應稅
	Taxable Tax = "1"
	// ZeroTax 零稅率
	ZeroTax Tax = "2"
	// DutyFree 免稅
	DutyFree Tax = "3"
	// Mix 若為混合應稅與免稅或零稅率時(限收 銀機發票無法分辨時使用，且需通過申 請核可)
	Mix Tax = "9"
)
