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
    name                          = "Adyen Payment Method"
    payment_source_type           = "AdyenPayment"
    require_capture               = true
    auto_place                    = true 
    auto_capture                  = true
    reference                     = "internal-payment-ref"
    currency_code                 = "EUR"
    price_amount_cents            = 0
    auto_capture_max_amount_cents = 10000
  }

  relationships {
    payment_gateway_id = commercelayer_adyen_gateway.labd_adyen_gateway.id
  }
}
