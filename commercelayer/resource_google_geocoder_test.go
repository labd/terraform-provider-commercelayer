package commercelayer

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	commercelayer "github.com/labd/go-commercelayer-sdk/api"
	"strings"
)

func testAccCheckGoogleGeocoderDestroy(s *terraform.State) error {
	client := testAccProviderCommercelayer.Meta().(*commercelayer.APIClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "commercelayer_google_geocoder" {
			_, resp, err := client.GoogleGeocodersApi.GETGoogleGeocodersGoogleGeocoderId(context.Background(), rs.Primary.ID).Execute()
			if resp.StatusCode == 404 {
				fmt.Printf("commercelayer_google_geocoder with id %s has been removed\n", rs.Primary.ID)
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

func (s *AcceptanceSuite) TestAccGoogleGeocoder_basic() {
	resourceName := "commercelayer_google_geocoder.labd_google_geocoder"

	resource.Test(s.T(), resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(s)
		},
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckGoogleGeocoderDestroy,
		Steps: []resource.TestStep{
			{
				Config: strings.Join([]string{testAccGoogleGeocoderCreate(resourceName)}, "\n"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "type", googleGeocodersType),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.name", "labd Google Geocoder"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.metadata.foo", "bar"),
				),
			},
			{
				Config: strings.Join([]string{testAccGoogleGeocoderUpdate(resourceName)}, "\n"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attributes.0.name", "labd Updated Google Geocoder"),
					resource.TestCheckResourceAttr(resourceName, "attributes.0.metadata.bar", "foo"),
				),
			},
		},
	})
}

func testAccGoogleGeocoderCreate(testName string) string {
	return hclTemplate(`
		resource "commercelayer_google_geocoder" "labd_google_geocoder" {
  			attributes {
    			name                   = "labd Google Geocoder"
    			api_key                = "Google Geocoder API Key"
				metadata = {
			  		foo : "bar"
		 	 		testName: "{{.testName}}"
				}
  			}
	}`, map[string]any{"testName": testName})
}

func testAccGoogleGeocoderUpdate(testName string) string {
	return hclTemplate(`
		resource "commercelayer_google_geocoder" "labd_google_geocoder" {
  			attributes {
    			name                   = "labd Updated Google Geocoder"
    			api_key                = "Google Geocoder API Key"
				metadata = {
			  		bar : "foo"
		 	 		testName: "{{.testName}}"
				}
  			}
	}`, map[string]any{"testName": testName})
}
