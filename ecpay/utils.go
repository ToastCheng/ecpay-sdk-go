package ecpay

import (
	"strings"
)

// FormatItemName joins an array of item name with '#'.
func FormatItemName(items []string) string {
	return strings.Join(items, "#")
}

type IgnorePaymentOption struct {
	Credit  bool
	WebATM  bool
	ATM     bool
	CVS     bool
	Barcode bool
}

// FormatIgnorePayment joins the payments that will be ignored with '#'.
func FormatIgnorePayment(option IgnorePaymentOption) string {
	res := make([]string, 0)
	if option.Credit {
		res = append(res, "Credit")
	}
	if option.WebATM {
		res = append(res, "WebATM")
	}
	if option.ATM {
		res = append(res, "ATM")
	}
	if option.CVS {
		res = append(res, "CVS")
	}
	if option.Barcode {
		res = append(res, "BARCODE")
	}
	return strings.Join(res, "#")
}

type CreditInstallmentOption struct {
	Month3  bool
	Month6  bool
	Month9  bool
	Month12 bool
	Month18 bool
	Month24 bool
}

// FormatCreditInstallmentOption joins the installment options with ','.
func FormatCreditInstallmentOption(option CreditInstallmentOption) string {
	res := make([]string, 0)
	if option.Month3 {
		res = append(res, "3")
	}
	if option.Month6 {
		res = append(res, "6")
	}
	if option.Month9 {
		res = append(res, "9")
	}
	if option.Month12 {
		res = append(res, "12")
	}
	if option.Month18 {
		res = append(res, "18")
	}
	if option.Month24 {
		res = append(res, "24")
	}
	return strings.Join(res, ",")
}

// FormatInvoiceItem joins the invoice item name with '|'.
// Apply to InvoiceItemName, InvoiceItemCount, InvoiceItemWord, InvoiceItemPrice and InvoiceItemTaxType.
func FormatInvoiceItem(items []string) string {
	return strings.Join(items, "|")
}
