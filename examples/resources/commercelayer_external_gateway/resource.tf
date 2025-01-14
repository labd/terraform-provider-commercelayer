resource "commercelayer_external_gateway" "labd_external_gateway" {
  attributes {
    name          = "labd External Gateway"
    authorize_url = "https://example.com"
    capture_url   = "https://foo.com"
    void_url      = "https://foo.com"
    refund_url    = "https://example.com"
    token_url     = "https://example.com"
  }
}
