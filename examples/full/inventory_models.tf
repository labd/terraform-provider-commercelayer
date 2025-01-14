resource "commercelayer_inventory_model" "labd_inventory_model" {
  attributes {
    name                   = "labd Inventory Model"
    stock_locations_cutoff = 2
    strategy               = "split_shipments"
  }
}
