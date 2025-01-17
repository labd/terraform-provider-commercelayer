---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "commercelayer_merchant Resource - terraform-provider-commercelayer"
subcategory: ""
description: |-
  A merchant is the fiscal representative that is selling in a specific market. Tax calculators use the merchant's address (and the shipping address) to determine the tax rate for an order.
---

# commercelayer_merchant (Resource)

A merchant is the fiscal representative that is selling in a specific market. Tax calculators use the merchant's address (and the shipping address) to determine the tax rate for an order.

## Example Usage

```terraform
resource "commercelayer_address" "labd_address" {
  attributes {
    business     = true
    company      = "labd"
    line_1       = "Van Nelleweg 1"
    zip_code     = "3044 BC"
    country_code = "NL"
    city         = "Rotterdam"
    phone        = "+31(0)10 20 20 544"
    state_code   = "ZH"
  }
}


resource "commercelayer_merchant" "labd_merchant" {
  attributes {
    name = "labd Merchant"
  }

  relationships {
    address_id = commercelayer_address.labd_address.id
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `attributes` (Block List, Min: 1, Max: 1) Resource attributes (see [below for nested schema](#nestedblock--attributes))
- `relationships` (Block List, Min: 1, Max: 1) Resource relationships (see [below for nested schema](#nestedblock--relationships))

### Read-Only

- `id` (String) The merchant unique identifier
- `type` (String) The resource type

<a id="nestedblock--attributes"></a>
### Nested Schema for `attributes`

Required:

- `name` (String) The merchant's internal name

Optional:

- `metadata` (Map of String) Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format
- `reference` (String) A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
- `reference_origin` (String) Any identifier of the third party system that defines the reference code


<a id="nestedblock--relationships"></a>
### Nested Schema for `relationships`

Required:

- `address_id` (String) The associated address id.
