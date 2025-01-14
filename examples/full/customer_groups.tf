resource "commercelayer_customer_group" "labd_customer_group" {
  attributes {
    name = "labd Customer Group"
    metadata = {
      foo : "bar"
    }
  }
}
