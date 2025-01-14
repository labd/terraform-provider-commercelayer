resource "commercelayer_price_list" "labd_price_list" {
  attributes {
    name          = "labd Price List"
    currency_code = "EUR"
    metadata = {
      foo : "bar"
    }
  }
}
