---
title: insights_resources_payment_evaluation_billing_details
page_id: schema-insights-resources-payment-evaluation-billing-details-234ff11a
path: schemas
description: Billing details attached to this payment evaluation.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_billing_details

Billing details attached to this payment evaluation.

```yaml
{"title": "InsightsResourcesPaymentEvaluationBillingDetails", "required": ["address"], "type": "object", "properties": {"address": {"$ref": "#/components/schemas/insights_resources_payment_evaluation_address"}, "email": {"maxLength": 5000, "type": "string", "description": "Email address.", "nullable": true}, "name": {"maxLength": 5000, "type": "string", "description": "Full name.", "nullable": true}, "phone": {"maxLength": 5000, "type": "string", "description": "Billing phone number (including extension).", "nullable": true}}, "description": "Billing details attached to this payment evaluation.", "x-expandableFields": ["address"]}
```
