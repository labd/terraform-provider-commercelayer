package commercelayer

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	commercelayer "github.com/labd/go-commercelayer-sdk/api"
)

func testAccCheckExternalTaxCalculatorDestroy(s *terraform.State) error {
	client := testAccProviderCommercelayer.Meta().(*commercelayer.APIClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "commercelayer_external_tax_calculator" {
			_, resp, err := client.ExternalTaxCalculatorsApi.GETExternalTaxCalculatorsExternalTaxCalculatorId(context.Background(), rs.Primary.ID).Execute()
			if resp.StatusCode == 404 {
				fmt.Printf("commercelayer_external_tax_calculator with id %s has been removed\n", rs.Primary.ID)
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

func (s *AcceptanceSuite) TestAccExternalTaxCalculator_basic() {
	resourceName := "commercelayer_external_tax_calculator.labd_external_tax_calculator"

	resource.Test(s.T(), resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(s)
		},
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckExternalTaxCalculatorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccExternalTaxCalculatorCreate(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "type", externalTaxCalculatorType),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.name", "labd_external_tax_calculator"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.metadata.foo", "bar"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.tax_calculator_url", "https://example.com"),
				),
			},
			{
				Config: testAccExternalTaxCalculatorUpdate(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attributes.0.name", "labd_external_tax_calculator_changed"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.metadata.bar", "foo"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.tax_calculator_url", "https://foo.com"),
				),
			},
		},
	})
}

func testAccExternalTaxCalculatorCreate(testName string) string {
	return hclTemplate(`
	resource "commercelayer_external_tax_calculator" "labd_external_tax_calculator" {
	  attributes {
		name          = "labd_external_tax_calculator"
		tax_calculator_url = "https://example.com"
		metadata = {
		  foo : "bar"
		  testName: "{{.testName}}"
		}
	  }
	}
	`, map[string]any{"testName": testName})
}

func testAccExternalTaxCalculatorUpdate(testName string) string {
	return hclTemplate(`
		resource "commercelayer_external_tax_calculator" "labd_external_tax_calculator" {
		  attributes {
			name          = "labd_external_tax_calculator_changed"
			tax_calculator_url = "https://foo.com"
			metadata = {
			  bar : "foo"
		 	  testName: "{{.testName}}"
			}
		  }
		}
	`, map[string]any{"testName": testName})
}
