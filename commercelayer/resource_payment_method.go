package commercelayer

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	commercelayer "github.com/labd/go-commercelayer-sdk/api"
)

func resourcePaymentMethod() *schema.Resource {
	return &schema.Resource{
		Description: "Payment methods represent the type of payment sources " +
			"(e.g., Credit Card, PayPal, or Apple Pay) offered in a market. " +
			"They can have a price and must be present before placing an order.",
		ReadContext:   resourcePaymentMethodReadFunc,
		CreateContext: resourcePaymentMethodCreateFunc,
		UpdateContext: resourcePaymentMethodUpdateFunc,
		DeleteContext: resourcePaymentMethodDeleteFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The payment method unique identifier",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "The resource type",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"attributes": {
				Description: "Resource attributes",
				Type:        schema.TypeList,
				MaxItems:    1,
				MinItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description:      "The payment method's internal name.",
							Type:             schema.TypeString,
							Optional:         true,
							ValidateDiagFunc: paymentSourceValidation,
						},
						"payment_source_type": {
							Description: "The payment source type, can be one of: 'adyen_payments', " +
								"'axerve_payments', 'braintree_payments', 'checkout_com_payments', 'credit_cards', " +
								"'external_payments', 'klarna_payments', 'paypal_payments', 'satispay_payments', " +
								"'stripe_payments', or 'wire_transfers'",
							Type:     schema.TypeString,
							Required: true,
						},
						"currency_code": {
							Description: "The international 3-letter currency code as defined by the ISO 4217 standard. " +
								"Required, unless inherited by market",
							Type:     schema.TypeString,
							Required: true,
						},
						"moto": {
							Description: "Send this attribute if you want to mark the payment as MOTO, " +
								"must be supported by payment gateway.",
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"require_capture": {
							Description: "Send this attribute if you want to require the payment capture before fulfillment.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
						},
						"auto_place": {
							Description: "Send this attribute if you want to automatically place the order upon authorization performed asynchronously.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
						},
						"auto_capture": {
							Description: "Send this attribute if you want to automatically capture the payment upon authorization.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
						},
						"price_amount_cents": {
							Description: "The payment method's price, in cents.",
							Type:        schema.TypeInt,
							Required:    true,
						},
						"auto_capture_max_amount_cents": {
							Description: "Send this attribute if you want to limit automatic capture to orders for which the total amount is equal or less than the specified value, in cents.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"_disable": {
							Description: "Send this attribute if you want to mark this resource as disabled.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"_enable": {
							Description: "Send this attribute if you want to mark this resource as enabled.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"reference": {
							Description: "A string that you can use to add any external identifier to the resource. This " +
								"can be useful for integrating the resource to an external system, like an ERP, a " +
								"marketing tool, a CRM, or whatever.",
							Type:     schema.TypeString,
							Optional: true,
						},
						"reference_origin": {
							Description: "Any identifier of the third party system that defines the reference code",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"metadata": {
							Description: "Set of key-value pairs that you can attach to the resource. This can be useful " +
								"for storing additional information about the resource in a structured format",
							Type: schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional: true,
						},
					},
				},
			},
			"relationships": {
				Description: "Resource relationships",
				Type:        schema.TypeList,
				MaxItems:    1,
				MinItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"market_id": {
							Description: "The associated market.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"payment_gateway_id": {
							Description: "The associated payment gateway.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func resourcePaymentMethodReadFunc(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	c := i.(*commercelayer.APIClient)

	resp, _, err := c.PaymentMethodsApi.GETPaymentMethodsPaymentMethodId(ctx, d.Id()).Execute()
	if err != nil {
		return diagErr(err)
	}

	address, ok := resp.GetDataOk()
	if !ok {
		d.SetId("")
		return nil
	}

	d.SetId(address.GetId().(string))

	return nil
}

func resourcePaymentMethodCreateFunc(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	c := i.(*commercelayer.APIClient)

	attributes := nestedMap(d.Get("attributes"))
	relationships := nestedMap(d.Get("relationships"))

	paymentMethodCreate := commercelayer.PaymentMethodCreate{
		Data: commercelayer.PaymentMethodCreateData{
			Type: paymentMethodType,
			Attributes: commercelayer.POSTPaymentMethods201ResponseDataAttributes{
				Name:        			   stringRef(attributes["name"]),
				PaymentSourceType:         attributes["payment_source_type"].(string),
				CurrencyCode:              stringRef(attributes["currency_code"]),
				Moto:                      boolRef(attributes["moto"]),
				PriceAmountCents:          int32(attributes["price_amount_cents"].(int)),
				RequireCapture:            boolRef(attributes["require_capture"]),
				AutoPlace:                 boolRef(attributes["auto_place"]),
				AutoCapture:               boolRef(attributes["auto_capture"]),
				AutoCaptureMaxAmountCents: int32(attributes["auto_capture_max_amount_cents"].(int)),
				Disable:                   boolRef(attributes["_disable"]),
				Enable:                    boolRef(attributes["_enable"]),
				Reference:                 stringRef(attributes["reference"]),
				ReferenceOrigin:           stringRef(attributes["reference_origin"]),
				Metadata:                  keyValueRef(attributes["metadata"]),
			},
			Relationships: &commercelayer.PaymentMethodCreateDataRelationships{
				PaymentGateway: commercelayer.PaymentMethodCreateDataRelationshipsPaymentGateway{
					Data: commercelayer.AdyenPaymentDataRelationshipsPaymentGatewayData{
						Type: stringRef(adyenGatewaysType),
						Id:   stringRef(relationships["payment_gateway_id"]),
					},
				},
			},
		},
	}

	marketId := stringRef(relationships["market_id"])
	if marketId != nil {
		paymentMethodCreate.Data.Relationships.Market = &commercelayer.BundleCreateDataRelationshipsMarket{
			Data: commercelayer.AvalaraAccountDataRelationshipsMarketsData{
				Type: stringRef(marketType),
				Id:   marketId,
			}}
	}

	paymentGatewayId := stringRef(relationships["payment_gateway_id"])
	if paymentGatewayId != nil {
		paymentMethodCreate.Data.Relationships.PaymentGateway = commercelayer.PaymentMethodCreateDataRelationshipsPaymentGateway{
			Data: commercelayer.AdyenPaymentDataRelationshipsPaymentGatewayData{
				Type: stringRef(paymentGatewayType),
				Id:   paymentGatewayId,
			},
		}
	}

	err := d.Set("type", paymentMethodType)
	if err != nil {
		return diagErr(err)
	}

	paymentMethod, _, err := c.PaymentMethodsApi.POSTPaymentMethods(ctx).PaymentMethodCreate(paymentMethodCreate).Execute()
	if err != nil {
		return diagErr(err)
	}

	d.SetId(paymentMethod.Data.GetId().(string))

	return nil
}

func resourcePaymentMethodDeleteFunc(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	c := i.(*commercelayer.APIClient)
	_, err := c.PaymentMethodsApi.DELETEPaymentMethodsPaymentMethodId(ctx, d.Id()).Execute()
	return diag.FromErr(err)
}

func resourcePaymentMethodUpdateFunc(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	c := i.(*commercelayer.APIClient)

	attributes := nestedMap(d.Get("attributes"))
	relationships := nestedMap(d.Get("relationships"))

	var paymentMethodUpdate = commercelayer.PaymentMethodUpdate{
		Data: commercelayer.PaymentMethodUpdateData{
			Type: paymentMethodType,
			Id:   d.Id(),
			Attributes: commercelayer.PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes{
				Name:        			   stringRef(attributes["name"]),
				PaymentSourceType:         stringRef(attributes["payment_source_type"]),
				CurrencyCode:              stringRef(attributes["currency_code"]),
				Moto:                      boolRef(attributes["moto"]),
				PriceAmountCents:          int32(attributes["price_amount_cents"].(int)),
				RequireCapture:            boolRef(attributes["require_capture"]),
				AutoPlace:                 boolRef(attributes["auto_place"]),
				AutoCapture:               boolRef(attributes["auto_capture"]),
				AutoCaptureMaxAmountCents: int32(attributes["auto_capture_max_amount_cents"].(int)),
				Disable:                   boolRef(attributes["_disable"]),
				Enable:                    boolRef(attributes["_enable"]),
				Reference:                 stringRef(attributes["reference"]),
				ReferenceOrigin:           stringRef(attributes["reference_origin"]),
				Metadata:                  keyValueRef(attributes["metadata"]),
			},
			Relationships: &commercelayer.PaymentMethodUpdateDataRelationships{
				PaymentGateway: &commercelayer.PaymentMethodCreateDataRelationshipsPaymentGateway{
					Data: commercelayer.AdyenPaymentDataRelationshipsPaymentGatewayData{
						Type: stringRef(adyenGatewaysType),
						Id:   stringRef(relationships["payment_gateway_id"]),
					},
				},
			},
		},
	}

	marketId := stringRef(relationships["market_id"])
	if marketId != nil {
		paymentMethodUpdate.Data.Relationships.Market = &commercelayer.BundleCreateDataRelationshipsMarket{
			Data: commercelayer.AvalaraAccountDataRelationshipsMarketsData{
				Type: stringRef(marketType),
				Id:   marketId,
			}}
	}

	paymentGatewayId := stringRef(relationships["payment_gateway_id"])
	if paymentGatewayId != nil {
		paymentMethodUpdate.Data.Relationships.PaymentGateway =
			&commercelayer.PaymentMethodCreateDataRelationshipsPaymentGateway{
				Data: commercelayer.AdyenPaymentDataRelationshipsPaymentGatewayData{
					Type: stringRef(paymentGatewayType),
					Id:   paymentGatewayId,
				},
			}
	}

	_, _, err := c.PaymentMethodsApi.PATCHPaymentMethodsPaymentMethodId(ctx, d.Id()).PaymentMethodUpdate(paymentMethodUpdate).Execute()

	return diag.FromErr(err)

}
