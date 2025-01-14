resource "commercelayer_shipping_method" "labd_shipping_method" {
  attributes {
    name                   = "labd Shipping Method"
    scheme                 = "flat"
    currency_code          = "EUR"
    price_amount_cents     = 1000
    free_over_amount_cents = 10000
    min_weight             = 0.50
    max_weight             = 10
    unit_of_weight         = "kg"
  }
}
