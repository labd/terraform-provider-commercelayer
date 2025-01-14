package commercelayer

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	commercelayer "github.com/labd/go-commercelayer-sdk/api"
)

func testAccCheckManualTaxCalculatorDestroy(s *terraform.State) error {
	client := testAccProviderCommercelayer.Meta().(*commercelayer.APIClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "commercelayer_manual_tax_calculator" {
			_, resp, err := client.ManualTaxCalculatorsApi.
				GETManualTaxCalculatorsManualTaxCalculatorId(context.Background(), rs.Primary.ID).Execute()
			if resp.StatusCode == 404 {
				fmt.Printf("commercelayer_manual_tax_calculator with id %s has been removed\n", rs.Primary.ID)
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

func (s *AcceptanceSuite) TestAccManualTaxCalculator_basic() {
	resourceName := "commercelayer_manual_tax_calculator.labd_manual_tax_calculator"

	resource.Test(s.T(), resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(s)
		},
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckManualTaxCalculatorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManualTaxCalculatorCreate(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "type", manualTaxCalculatorsType),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.name", "labd Manual Tax Calculator"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.metadata.foo", "bar"),
				),
			},
			{
				Config: testAccManualTaxCalculatorUpdate(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attributes.0.name", "labd Manual Tax Calculator Changed"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.metadata.bar", "foo"),
				),
			},
		},
	})
}

func testAccManualTaxCalculatorCreate(testName string) string {
	return hclTemplate(`
		resource "commercelayer_manual_tax_calculator" "labd_manual_tax_calculator" {
           attributes {
			name                   = "labd Manual Tax Calculator"
			metadata = {
				foo: "bar"
				testName: "{{.testName}}"
    		}
  		}
	}
`, map[string]any{"testName": testName})
}

func testAccManualTaxCalculatorUpdate(testName string) string {
	return hclTemplate(`
		resource "commercelayer_manual_tax_calculator" "labd_manual_tax_calculator" {
           attributes {
			name                   = "labd Manual Tax Calculator Changed"
			metadata = {
				bar: "foo"
				testName: "{{.testName}}"
    		}
  		}
	}
`, map[string]any{"testName": testName})
}
