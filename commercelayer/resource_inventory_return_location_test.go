package commercelayer

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	commercelayer "github.com/labd/go-commercelayer-sdk/api"
	"net/http"
	"strings"
)

func testAccCheckInventoryReturnLocationDestroy(s *terraform.State) error {
	client := testAccProviderCommercelayer.Meta().(*commercelayer.APIClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "commercelayer_inventory_return_location" {
			err := retryRemoval(10, func() (*http.Response, error) {
				_, resp, err := client.InventoryReturnLocationsApi.
					GETInventoryReturnLocationsInventoryReturnLocationId(context.Background(), rs.Primary.ID).
					Execute()
				return resp, err
			})
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (s *AcceptanceSuite) TestAccInventoryReturnLocation_basic() {
	resourceName := "commercelayer_inventory_return_location.labd_inventory_return_location"

	resource.Test(s.T(), resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(s)
		},
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckInventoryReturnLocationDestroy,
		Steps: []resource.TestStep{
			{
				Config: strings.Join([]string{
					testAccAddressCreate(resourceName),
					testAccInventoryModelCreate(resourceName),
					testAccStockLocationCreate(resourceName),
					testAccInventoryReturnLocationCreate(resourceName),
				}, "\n"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attributes.0.priority", "1"),
				),
			},
			{
				Config: strings.Join([]string{
					testAccAddressCreate(resourceName),
					testAccInventoryModelCreate(resourceName),
					testAccStockLocationCreate(resourceName),
					testAccInventoryReturnLocationUpdate(resourceName),
				}, "\n"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attributes.0.priority", "2"),
				),
			},
		},
	})
}

func testAccInventoryReturnLocationCreate(testName string) string {
	return hclTemplate(`
		resource "commercelayer_inventory_return_location" "labd_inventory_return_location" {
		  attributes {
			priority = 1
			metadata = {
			  testName: "{{.testName}}"
			}
		  }

		  relationships {
			inventory_model_id = commercelayer_inventory_model.labd_inventory_model.id
			stock_location_id  = commercelayer_stock_location.labd_stock_location.id
		  }
		}
	`, map[string]any{"testName": testName})
}

func testAccInventoryReturnLocationUpdate(testName string) string {
	return hclTemplate(`
		resource "commercelayer_inventory_return_location" "labd_inventory_return_location" {
		  attributes {
			priority = 2
			metadata = {
			  testName: "{{.testName}}"
			}
		  }

		  relationships {
			inventory_model_id = commercelayer_inventory_model.labd_inventory_model.id
			stock_location_id  = commercelayer_stock_location.labd_stock_location.id
		  }
		}
	`, map[string]any{"testName": testName})
}
