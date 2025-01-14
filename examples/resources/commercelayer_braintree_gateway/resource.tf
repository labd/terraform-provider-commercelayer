resource "commercelayer_braintree_gateway" "labd_braintree_gateway" {
  attributes {
    name                = "labd Braintree Gateway"
    merchant_account_id = "xxxx-yyyy-zzzz"
    merchant_id         = "xxxx-yyyy-zzzz"
    public_key          = "xxxx-yyyy-zzzz"
    private_key         = "xxxx-yyyy-zzzz"
  }
}
