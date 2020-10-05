package ecpay

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
)

// FormatItemName joins an array of item name with '#'.
func FormatItemName(items []string) string {
	return strings.Join(items, "#")
}

// IgnorePaymentOption defines the struct of IgnorePayment option.
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

// CreditInstallmentOption defines the struct of CreditInstallment option.
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

// FormatDatetime convert time.Time into yyyy/MM/dd HH:mm:ss format.
func FormatDatetime(time time.Time) string {
	return time.Format("2006/01/02 15:04:05")
}

// FormatDate convert time.Time into yyyy-MM=dd format.
func FormatDate(time time.Time) string {
	return time.Format("2006-01-02")
}

// ParseQueryString parse the query string to map.
func ParseQueryString(q string) (map[string]interface{}, error) {
	uv, err := url.ParseQuery(q)
	if err != nil {
		return nil, err
	}
	mp := map[string]interface{}{}
	for k, v := range uv {
		mp[k] = v[0]
	}
	return mp, err
}

// GeneratePostForm generates html post form for redirecting client to payment site.
func GeneratePostForm(action string, form url.Values) string {
	id := uuid.New().String()

	html := "<form id=\"" + id + "\" action=\"" + action + "\" method=\"post\">"
	for k, v := range form {
		html += "<input type=\"hidden\" name=\"" + k + "\" "
		html += "id=\"" + k + "\" "
		html += "value=\"" + v[0] + "\" />"
	}
	html += "<script type=\"text/javascript\">document.getElementById(\"" + id + "\").submit();</script>"
	html += "</form>"

	return html
}

func setUrlValues(req url.Values, mp map[string]interface{}) url.Values {
	for k, v := range mp {
		switch t := v.(type) {
		case float32, float64:
			req.Set(k, fmt.Sprintf("%.0f", t))
		case string:
			req.Set(k, t)
		}
	}
	return req
}
