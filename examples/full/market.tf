resource "commercelayer_market" "labd_market" {
  attributes {
    name = "labd Market"
    //TODO: check why these are considered invalid
    #    checkout_url        = "http://example.url"
    #    external_prices_url = "http://example.url"
    facebook_pixel_id             = "pixel"
    external_order_validation_url = "https://www.example.com"
  }

  relationships {
    inventory_model_id = commercelayer_inventory_model.labd_inventory_model.id
    merchant_id        = commercelayer_merchant.labd_merchant.id
    price_list_id      = commercelayer_price_list.labd_price_list.id
    customer_group_id  = commercelayer_customer_group.labd_customer_group.id
    tax_calculator_id  = commercelayer_external_tax_calculator.labd_external_tax_calculator.id
  }
}
