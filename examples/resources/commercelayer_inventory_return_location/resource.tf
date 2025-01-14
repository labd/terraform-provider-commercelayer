resource "commercelayer_inventory_model" "labd_inventory_model" {
  attributes {
    name                   = "labd Inventory Model Return Location"
    stock_locations_cutoff = 2
    strategy               = "split_shipments"
  }
}

resource "commercelayer_address" "labd_address" {
  attributes {
    business     = true
    company      = "labd"
    line_1       = "Van Nelleweg 1"
    zip_code     = "3044 BC"
    country_code = "NL"
    city         = "Rotterdam"
    phone        = "+31(0)10 20 20 544"
    state_code   = "ZH"
  }
}

resource "commercelayer_stock_location" "labd_stock_location" {
  attributes {
    name         = "labd Warehouse Location"
    label_format = "PNG"
    suppress_etd = true
  }

  relationships {
    address_id = commercelayer_address.labd_address.id
  }
}

resource "commercelayer_inventory_return_location" "labd_return_location" {
  attributes {
    priority = 1
  }

  relationships {
    inventory_model_id = commercelayer_inventory_model.labd_inventory_model.id
    stock_location_id  = commercelayer_stock_location.labd_stock_location.id
  }
}
