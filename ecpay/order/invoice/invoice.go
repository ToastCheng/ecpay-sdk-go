package invoice

// Invoice defines the struct of invoice.
type Invoice string

const (
	// General 一般稅額
	General Invoice = "07"
	// Special 特種稅額
	Special Invoice = "08"
)
