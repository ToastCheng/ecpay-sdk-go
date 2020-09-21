package tax

// Tax tax type.
type Tax string

const (
	// Dutiable 應稅
	Dutiable Tax = "1"
	// Zero 零稅率
	Zero Tax = "2"
	// Free 免稅
	Free Tax = "3"
	// Mix 若為混合應稅與免稅或零稅率時(限收 銀機發票無法分辨時使用，且需通過申 請核可)
	Mix Tax = "9"
)
