resource "commercelayer_external_tax_calculator" "labd_external_tax_calculator" {
  attributes {
    name               = "labd External Tax Calculator"
    tax_calculator_url = "https://example.com"
    metadata = {
      foo : "bar"
    }
  }
}
