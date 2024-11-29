package commercelayer

import (
	diag2 "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrencyCodeValidationErr(t *testing.T) {
	diag := currencyCodeValidation("FOOBAR", nil)
	assert.True(t, diag.HasError())
}

func TestCurrencyCodeValidationOK(t *testing.T) {
	diag := currencyCodeValidation("EUR", nil)
	assert.False(t, diag.HasError())
}

func TestPaymentSourceValidationWarning(t *testing.T) {
	diag := paymentSourceValidation("AdyenPayment", nil)
	assert.Len(t, diag, 1)
	assert.Equal(t, diag[0].Severity, diag2.Warning)
}

func TestPaymentSourceValidationOK(t *testing.T) {
	diag := paymentSourceValidation("braintree_payments", nil)
	assert.False(t, diag.HasError())
}
