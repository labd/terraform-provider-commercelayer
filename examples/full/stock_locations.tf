resource "commercelayer_stock_location" "labd_warehouse_location" {
  attributes {
    name         = "labd Warehouse Location"
    label_format = "PNG"
    suppress_etd = true
    metadata = {
      foo : "bar"
    }
  }

  relationships {
    address_id = commercelayer_address.labd_address.id
  }
}

resource "commercelayer_stock_location" "labd_backorder_location" {
  attributes {
    name         = "labd Backorder Location"
    label_format = "PNG"
    suppress_etd = true
    metadata = {
      foo : "bar"
    }
  }

  relationships {
    address_id = commercelayer_address.labd_address.id
  }
}
