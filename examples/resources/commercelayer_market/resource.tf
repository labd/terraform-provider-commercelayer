resource "commercelayer_inventory_model" "labd_inventory_model" {
  attributes {
    name                   = "labd Inventory Model"
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


resource "commercelayer_merchant" "labd_merchant" {
  attributes {
    name = "labd Merchant"
  }

  relationships {
    address_id = commercelayer_address.labd_address.id
  }
}

resource "commercelayer_price_list" "labd_price_list" {
  attributes {
    name          = "labd Price List"
    currency_code = "EUR"
  }
}

resource "commercelayer_market" "labd_market" {
  attributes {
    name                          = "labd Market"
    facebook_pixel_id             = "pixel"
    external_order_validation_url = "https://www.example.com"
  }

  relationships {
    inventory_model_id = commercelayer_inventory_model.labd_inventory_model.id
    merchant_id        = commercelayer_merchant.labd_merchant.id
    price_list_id      = commercelayer_price_list.labd_price_list.id
  }
}
