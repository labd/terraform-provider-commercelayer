---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "commercelayer_webhook Resource - terraform-provider-commercelayer"
subcategory: ""
description: |-
  A webhook object is returned as part of the response body of each successful list, retrieve, create or update API call to the /api/webhooks endpoint.
---

# commercelayer_webhook (Resource)

A webhook object is returned as part of the response body of each successful list, retrieve, create or update API call to the /api/webhooks endpoint.

## Example Usage

```terraform
resource "commercelayer_webhook" "labd_webhook" {
  attributes {
    name         = "labd Webhook"
    topic        = "orders.create"
    callback_url = "http://example.url"
    include_resources = [
      "customer",
      "line_items"
    ]
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `attributes` (Block List, Min: 1, Max: 1) Resource attributes (see [below for nested schema](#nestedblock--attributes))

### Read-Only

- `id` (String) The webhook unique identifier
- `shared_secret` (String, Sensitive) The shared secret used to sign the external request payload.
- `type` (String) The resource type

<a id="nestedblock--attributes"></a>
### Nested Schema for `attributes`

Required:

- `callback_url` (String) URI where the webhook subscription should send the POST request when the event occurs.
- `topic` (String) The identifier of the resource/event that will trigger the webhook.

Optional:

- `include_resources` (List of String) List of related commercelayer_inventory_stock_location that should be included in the webhook body.
- `metadata` (Map of String) Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format
- `name` (String) Unique name for the webhook.
- `reference` (String) A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
- `reference_origin` (String) Any identifier of the third party system that defines the reference code
