resource "commercelayer_google_geocoder" "labd_google_geocoder" {
  attributes {
    name    = "labd Google Geocoder"
    api_key = "Google Geocoder API Key"
    metadata = {
      foo : "bar"
    }
  }
}
