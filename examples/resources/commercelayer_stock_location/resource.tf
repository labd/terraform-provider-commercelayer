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


resource "commercelayer_stock_location" "labd_warehouse_location" {
  attributes {
    name         = "labd Warehouse Stock Location"
    label_format = "PNG"
    suppress_etd = true
  }

  relationships {
    address_id = commercelayer_address.labd_address.id
  }
}
