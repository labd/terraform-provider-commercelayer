resource "commercelayer_inventory_return_location" "labd_return_location" {
  attributes {
    priority = 1
  }

  relationships {
    inventory_model_id = commercelayer_inventory_model.labd_inventory_model.id
    stock_location_id  = commercelayer_stock_location.labd_warehouse_location.id
  }
}
