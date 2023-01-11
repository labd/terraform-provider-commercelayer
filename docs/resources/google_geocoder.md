---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "commercelayer_google_geocoder Resource - terraform-provider-commercelayer"
subcategory: ""
description: |-
  Geocoders lets you convert an address in text form into geographic coordinates (like latitude and longitude). A geocoder can be optionally associated with a market. By connecting a geocoder to a market, all the shipping and billing addresses belonging to that market orders will be geocoded
---

# commercelayer_google_geocoder (Resource)

Geocoders lets you convert an address in text form into geographic coordinates (like latitude and longitude). A geocoder can be optionally associated with a market. By connecting a geocoder to a market, all the shipping and billing addresses belonging to that market orders will be geocoded

## Example Usage

```terraform
resource "commercelayer_google_geocoder" "incentro_google_geocoder" {
  attributes {
    name    = "Incentro Google Geocoder"
    api_key = "Google Geocoder API Key"
    metadata = {
      foo : "bar"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `attributes` (Block List, Min: 1, Max: 1) Resource attributes (see [below for nested schema](#nestedblock--attributes))

### Read-Only

- `id` (String) The google geocoder unique identifier
- `type` (String) The resource type

<a id="nestedblock--attributes"></a>
### Nested Schema for `attributes`

Required:

- `api_key` (String) The Google Map API key
- `name` (String) The geocoder's internal name.

Optional:

- `metadata` (Map of String) Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format
- `reference` (String) A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
- `reference_origin` (String) Any identifier of the third party system that defines the reference code

