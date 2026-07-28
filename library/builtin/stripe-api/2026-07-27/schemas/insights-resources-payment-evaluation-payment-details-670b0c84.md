---
title: insights_resources_payment_evaluation_payment_details
page_id: schema-insights-resources-payment-evaluation-payment-details-670b0c84
path: schemas
description: Payment details attached to this payment evaluation.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_payment_details

Payment details attached to this payment evaluation.

```yaml
{"title": "InsightsResourcesPaymentEvaluationPaymentDetails", "required": ["amount", "currency"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount intended to be collected by this payment. A positive integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The minimum amount is $0.50 US or [equivalent in charge currency](https://docs.stripe.com/currencies#minimum-and-maximum-charge-amounts). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99)."}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the object. Often useful for displaying to users.", "nullable": true}, "money_movement_details": {"description": "Details about the payment's customer presence and type.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/insights_resources_payment_evaluation_money_movement_details"}]}, "payment_method_details": {"description": "Details about the payment method used for the payment.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/insights_resources_payment_evaluation_payment_method_details"}]}, "shipping_details": {"description": "Shipping details for the payment evaluation.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/insights_resources_payment_evaluation_shipping"}]}, "statement_descriptor": {"maxLength": 5000, "type": "string", "description": "Payment statement descriptor.", "nullable": true}}, "description": "Payment details attached to this payment evaluation.", "x-expandableFields": ["money_movement_details", "payment_method_details", "shipping_details"]}
```
