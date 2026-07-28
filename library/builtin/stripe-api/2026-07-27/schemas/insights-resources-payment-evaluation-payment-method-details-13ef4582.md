---
title: insights_resources_payment_evaluation_payment_method_details
page_id: schema-insights-resources-payment-evaluation-payment-method-details-13ef4582
path: schemas
description: Payment method details attached to this payment evaluation.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_payment_method_details

Payment method details attached to this payment evaluation.

```yaml
{"title": "InsightsResourcesPaymentEvaluationPaymentMethodDetails", "required": ["payment_method"], "type": "object", "properties": {"billing_details": {"description": "Billing information associated with the payment evaluation.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/insights_resources_payment_evaluation_billing_details"}]}, "payment_method": {"description": "The payment method used in this payment evaluation.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/payment_method"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/payment_method"}]}}}, "description": "Payment method details attached to this payment evaluation.", "x-expandableFields": ["billing_details", "payment_method"]}
```
