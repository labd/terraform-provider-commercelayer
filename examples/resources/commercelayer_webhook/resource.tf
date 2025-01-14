resource "commercelayer_webhook" "labd_webhook" {
  attributes {
    name         = "labd Webhook"
    topic        = "orders.create"
    callback_url = "http://example.url"
    include_resources = [
      "customer",
      "line_items"
    ]
  }
}
