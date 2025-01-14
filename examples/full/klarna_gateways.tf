resource "commercelayer_klarna_gateway" "labd_klarna_gateway" {
  attributes {
    name         = "labd Klarna Gateway"
    country_code = "EU"
    api_key      = "xxxx-yyyy-zzzz"
    api_secret   = "xxxx-yyyy-zzzz"
  }
}
