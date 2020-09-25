package ecpay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatItemName(t *testing.T) {
	items := []string{"item1", "item2"}
	assert.Equal(t, FormatItemName(items), "item1#item2")
}

func TestIgnorePayment(t *testing.T) {
	option := IgnorePaymentOption{
		Credit:  true,
		ATM:     true,
		Barcode: true,
	}
	assert.Equal(t, FormatIgnorePayment(option), "Credit#ATM#BARCODE")

}

func TestFormatCreditInstallmentOption(t *testing.T) {
	option := CreditInstallmentOption{
		Month3:  true,
		Month12: true,
	}
	assert.Equal(t, FormatCreditInstallmentOption(option), "3,12")

}
