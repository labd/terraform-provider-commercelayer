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
