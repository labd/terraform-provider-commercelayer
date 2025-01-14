---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "commercelayer_checkout_com_gateway Resource - terraform-provider-commercelayer"
subcategory: ""
description: |-
  Configuring a CheckoutCom payment gateway for a market lets you safely process payments through CheckoutCom. The CheckoutCom gateway is compliant with the PSD2 European regulation so that you canimplement a payment flow that supports SCA and 3DS2 by using the CheckoutCom's official JS SDK and libraries.
---

# commercelayer_checkout_com_gateway (Resource)

Configuring a CheckoutCom payment gateway for a market lets you safely process payments through CheckoutCom. The CheckoutCom gateway is compliant with the PSD2 European regulation so that you canimplement a payment flow that supports SCA and 3DS2 by using the CheckoutCom's official JS SDK and libraries.

## Example Usage

```terraform
resource "commercelayer_checkout_com_gateway" "labd_checkout_com_gateway" {
  attributes {
    name       = "labd CheckoutCom Gateway"
    secret_key = "xxxx-yyyy-zzzz"
    public_key = "xxxx-yyyy-zzzz"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `attributes` (Block List, Min: 1, Max: 1) Resource attributes (see [below for nested schema](#nestedblock--attributes))

### Read-Only

- `id` (String) The checkout.com payment unique identifier
- `type` (String) The resource type

<a id="nestedblock--attributes"></a>
### Nested Schema for `attributes`

Required:

- `name` (String) The payment gateway's internal name.
- `public_key` (String) The gateway public key.
- `secret_key` (String) The gateway secret key.

Optional:

- `metadata` (Map of String) Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format
- `reference` (String) A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
- `reference_origin` (String) Any identifier of the third party system that defines the reference code
