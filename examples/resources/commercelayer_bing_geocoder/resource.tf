resource "commercelayer_bing_geocoder" "labd_bing_geocoder" {
  attributes {
    name = "labd Bing Geocoder"
    key  = "Bing Virtualearth Key"
    metadata = {
      foo : "bar"
    }
  }
}
