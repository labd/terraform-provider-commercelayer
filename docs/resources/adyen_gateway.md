---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "commercelayer_adyen_gateway Resource - terraform-provider-commercelayer"
subcategory: ""
description: |-
  Configuring a Adyen payment gateway for a market lets you safely process payments through Adyen. The Adyen gateway is compliant with the PSD2 European regulation so that you can implement a payment flow that supports SCA and 3DS2 by using the Adyen's official JS SDK and libraries.To create a Adyen gateway choose a meaningful name that helps you identify it within your organization and gather all the credentials requested (like secret and publishable keys, etc. — contact Adyen's support if you are not sure about the requested data).
---

# commercelayer_adyen_gateway (Resource)

Configuring a Adyen payment gateway for a market lets you safely process payments through Adyen. The Adyen gateway is compliant with the PSD2 European regulation so that you can implement a payment flow that supports SCA and 3DS2 by using the Adyen's official JS SDK and libraries.To create a Adyen gateway choose a meaningful name that helps you identify it within your organization and gather all the credentials requested (like secret and publishable keys, etc. — contact Adyen's support if you are not sure about the requested data).

## Example Usage

```terraform
resource "commercelayer_adyen_gateway" "labd_adyen_gateway" {
  attributes {
    name                    = "labd Adyen Gateway"
    merchant_account        = "xxxx-yyyy-zzzz"
    api_key                 = "xxxx-yyyy-zzzz"
    public_key              = "xxxx-yyyy-zzzz"
    live_url_prefix         = "1797a841fbb37ca7-AdyenDemo"
    api_version             = 68
    async_api               = true
    webhook_endpoint_secret = "foobar"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `attributes` (Block List, Min: 1, Max: 1) Resource attributes (see [below for nested schema](#nestedblock--attributes))

### Read-Only

- `id` (String) The adyen payment unique identifier
- `type` (String) The resource type

<a id="nestedblock--attributes"></a>
### Nested Schema for `attributes`

Required:

- `api_key` (String) The gateway API key.
- `live_url_prefix` (String) The prefix of the endpoint used for live transactions.
- `merchant_account` (String) The gateway merchant account.
- `name` (String) The payment gateway's internal name.

Optional:

- `api_version` (String) The checkout API version, supported range is from 66 to 68, default is 68.
- `async_api` (Boolean) Indicates if the gateway will leverage on the Adyen notification webhooks.
- `metadata` (Map of String) Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format
- `public_key` (String) The public key linked to your API credential.
- `reference` (String) A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
- `reference_origin` (String) Any identifier of the third party system that defines the reference code
- `webhook_endpoint_secret` (String) The gateway webhook endpoint secret, generated by Adyen customer area.
