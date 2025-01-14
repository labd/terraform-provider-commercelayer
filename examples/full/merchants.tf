resource "commercelayer_merchant" "labd_merchant" {
  attributes {
    name = "labd Merchant"
    metadata = {
      foo : "bar"
    }
  }

  relationships {
    address_id = commercelayer_address.labd_address.id
  }
}
