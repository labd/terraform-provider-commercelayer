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
    payment_source_type           = "AdyenPayment"
    currency_code                 = "EUR"
    price_amount_cents            = 0
    auto_capture_max_amount_cents = 100
  }

  relationships {
    payment_gateway_id = commercelayer_adyen_gateway.labd_adyen_gateway.id
  }
}
