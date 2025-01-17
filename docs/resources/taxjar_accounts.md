---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "commercelayer_taxjar_accounts Resource - terraform-provider-commercelayer"
subcategory: ""
description: |-
  Configure your TaxJar account to automatically compute tax calculations for the orders of the associated market.
---

# commercelayer_taxjar_accounts (Resource)

Configure your TaxJar account to automatically compute tax calculations for the orders of the associated market.

## Example Usage

```terraform
resource "commercelayer_taxjar_accounts" "labd_taxjar_account" {
  attributes {
    name    = "labd Taxjar Account"
    api_key = "TAXJAR_API_KEY"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `attributes` (Block List, Min: 1, Max: 1) Resource attributes (see [below for nested schema](#nestedblock--attributes))

### Read-Only

- `id` (String) The taxjar account unique identifier
- `type` (String) The resource type

<a id="nestedblock--attributes"></a>
### Nested Schema for `attributes`

Required:

- `api_key` (String) The TaxJar account API key.
- `name` (String) The tax calculator's internal name.

Optional:

- `metadata` (Map of String) Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format
- `reference` (String) A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
- `reference_origin` (String) Any identifier of the third party system that defines the reference code
