---
title: payment_flows_private_payment_methods_paypal_amount_details_line_item_payment_method_options
page_id: schema-payment-flows-private-payment-methods-paypal-amount-details-line-item-pa-2184cffe
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_flows_private_payment_methods_paypal_amount_details_line_item_payment_method_options

```yaml
{"title": "PaymentFlowsPrivatePaymentMethodsPaypalAmountDetailsLineItemPaymentMethodOptions", "type": "object", "properties": {"category": {"type": "string", "description": "Type of the line item.", "enum": ["digital_goods", "donation", "physical_goods"]}, "description": {"maxLength": 5000, "type": "string", "description": "Description of the line item."}, "sold_by": {"maxLength": 5000, "type": "string", "description": "The Stripe account ID of the connected account that sells the item. This is only needed when using [Separate Charges and Transfers](https://docs.stripe.com/connect/separate-charges-and-transfers)."}}, "description": "", "x-expandableFields": []}
```
