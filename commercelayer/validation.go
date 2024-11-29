package commercelayer

import (
	"fmt"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/ladydascalie/currency"
	"strings"
)

var currencyCodeValidation = func(i interface{}, path cty.Path) diag.Diagnostics {
	_, err := currency.Get(i.(string))
	return diagErr(err)
}

func getInventoryModelStrategies() []string {
	return []string{
		"no_split",
		"split_shipments",
		"split_by_line_items",
		"ship_from_primary",
		"ship_from_first_available_or_primary",
	}
}

var inventoryModelStrategyValidation = func(i interface{}, path cty.Path) diag.Diagnostics {
	for _, s := range getInventoryModelStrategies() {
		if s == i.(string) {
			return nil
		}
	}
	return diag.Errorf("Invalid inventory model strategy provided: %s. Must be one of %s",
		i.(string), strings.Join(getInventoryModelStrategies(), ", "))
}

var paymentSourceValidation = func(i interface{}, path cty.Path) diag.Diagnostics {
	var paymentMap = map[string]string{
		"AdyenPayment":       "adyen_payments",
		"BraintreePayment":   "braintree_payments",
		"CheckoutComPayment": "checkout_com_payments",
		"CreditCard":         "credit_cards",
		"ExternalPayment":    "external_payments",
		"KlarnaPayment":      "klarna_payments",
		"PaypalPayment":      "paypal_payments",
		"SatispayPayment":    "satispay_payments",
		"StripePayment":      "stripe_payments",
		"WireTransfer":       "wire_transfers",
	}

	if val, ok := paymentMap[i.(string)]; ok {
		return []diag.Diagnostic{
			{
				Severity: diag.Warning,
				Summary:  fmt.Sprintf("Payment source type key %s has been deprecated", i.(string)),
				Detail:   fmt.Sprintf("Payment source type key `%s` has been deprecated, Please use the new key `%s` instead", i.(string), val),
			},
		}
	}

	return nil
}
