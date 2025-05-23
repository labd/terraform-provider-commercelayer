---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "commercelayer_payment_method Resource - terraform-provider-commercelayer"
subcategory: ""
description: |-
  Payment methods represent the type of payment sources (e.g., Credit Card, PayPal, or Apple Pay) offered in a market. They can have a price and must be present before placing an order.
---

# commercelayer_payment_method (Resource)

Payment methods represent the type of payment sources (e.g., Credit Card, PayPal, or Apple Pay) offered in a market. They can have a price and must be present before placing an order.

## Example Usage

```terraform
resource "commercelayer_adyen_gateway" "labd_adyen_gateway" {
  attributes {
    name             = "labd Adyen Gateway"
    merchant_account = "xxxx-yyyy-zzzz"
    api_key          = "xxxx-yyyy-zzzz"
    public_key       = "xxxx-yyyy-zzzz"
    live_url_prefix  = "1797a841fbb37ca7-AdyenDemo"
  }
}

resource "commercelayer_payment_method" "labd_payment_method" {
  attributes {
    payment_source_type = "AdyenPayment"
    currency_code       = "EUR"
    price_amount_cents  = 0
  }

  relationships {
    payment_gateway_id = commercelayer_adyen_gateway.labd_adyen_gateway.id
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `attributes` (Block List, Min: 1, Max: 1) Resource attributes (see [below for nested schema](#nestedblock--attributes))
- `relationships` (Block List, Min: 1, Max: 1) Resource relationships (see [below for nested schema](#nestedblock--relationships))

### Read-Only

- `id` (String) The payment method unique identifier
- `type` (String) The resource type

<a id="nestedblock--attributes"></a>
### Nested Schema for `attributes`

Required:

- `currency_code` (String) The international 3-letter currency code as defined by the ISO 4217 standard. Required, unless inherited by market
- `payment_source_type` (String) The payment source type, can be one of: 'adyen_payments', 'axerve_payments', 'braintree_payments', 'checkout_com_payments', 'credit_cards', 'external_payments', 'klarna_payments', 'paypal_payments', 'satispay_payments', 'stripe_payments', or 'wire_transfers'
- `price_amount_cents` (Number) The payment method's price, in cents.

Optional:

- `metadata` (Map of String) Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format
- `moto` (Boolean) Send this attribute if you want to mark the payment as MOTO, must be supported by payment gateway.
- `reference` (String) A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
- `reference_origin` (String) Any identifier of the third party system that defines the reference code


<a id="nestedblock--relationships"></a>
### Nested Schema for `relationships`

Required:

- `payment_gateway_id` (String) The associated payment gateway.

Optional:

- `market_id` (String) The associated market.
