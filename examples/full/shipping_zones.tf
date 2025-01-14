resource "commercelayer_shipping_zone" "labd_shipping_zone" {
  attributes {
    name                   = "labd Shipping Zone"
    country_code_regex     = ".*"
    not_country_code_regex = "[^i*&2@]"
    state_code_regex       = "^dog"
    not_state_code_regex   = "//[^\r\n]*[\r\n]"
    zip_code_regex         = "[a-zA-Z]{2,4}"
    not_zip_code_regex     = ".+"
    metadata = {
      foo : "bar"
    }
  }
}
