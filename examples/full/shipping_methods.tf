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
    metadata = {
      foo : "bar"
    }
  }

  relationships {
    market_id            = commercelayer_market.labd_market.id
    shipping_zone_id     = commercelayer_shipping_zone.labd_shipping_zone.id
    shipping_category_id = commercelayer_shipping_category.labd_shipping_category.id
    stock_location_id    = commercelayer_stock_location.labd_warehouse_location.id
  }
}
