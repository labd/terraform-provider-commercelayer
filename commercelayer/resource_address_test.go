package commercelayer

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	commercelayer "github.com/labd/go-commercelayer-sdk/api"
)

func testAccCheckAddressDestroy(s *terraform.State) error {
	client := testAccProviderCommercelayer.Meta().(*commercelayer.APIClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "commercelayer_address" {
			_, resp, err := client.AddressesApi.GETAddressesAddressId(context.Background(), rs.Primary.ID).Execute()
			if resp.StatusCode == 404 {
				fmt.Printf("commercelayer_address with id %s has been removed\n", rs.Primary.ID)
				continue
			}
			if err != nil {
				return err
			}

			return fmt.Errorf("received response code with status %d", resp.StatusCode)
		}

	}
	return nil
}

func (s *AcceptanceSuite) TestAccAddress_basic() {
	resourceName := "commercelayer_address.labd_address"

	resource.Test(s.T(), resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(s)
		},
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAddressCreate(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "type", addressType),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.business", "true"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.company", "labd"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.line_1", "Van Nelleweg 1"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.zip_code", "3044 BC"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.country_code", "NL"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.city", "Rotterdam"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.phone", "+31(0)10 20 20 544"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.state_code", "ZH"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.metadata.foo", "bar"),
				),
			},
			{
				Config: testAccAddressUpdate(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attributes.0.line_1", "Moermanskkade 113"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.zip_code", "1013 BC"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.country_code", "NL"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.city", "Amsterdam"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.phone", "020 409 0444"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.state_code", "NH"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.metadata.bar", "foo"),
				),
			},
		},
	})
}

func testAccAddressCreate(testName string) string {
	return hclTemplate(`
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
			metadata = {
			  foo: "bar"
			  testName: "{{.testName}}"
			}
		  }
		}
	`, map[string]any{"testName": testName})
}

func testAccAddressUpdate(testName string) string {
	return hclTemplate(`
		resource "commercelayer_address" "labd_address" {
		  attributes {
			business     = true
			company      = "labd"
			line_1       = "Moermanskkade 113"
			zip_code     = "1013 BC"
			country_code = "NL"
			city         = "Amsterdam"
			phone        = "020 409 0444"
			state_code   = "NH"
			metadata = {
			  bar: "foo"
			  testName: "{{.testName}}"
			}
		  }
		}
	`, map[string]any{"testName": testName})
}
