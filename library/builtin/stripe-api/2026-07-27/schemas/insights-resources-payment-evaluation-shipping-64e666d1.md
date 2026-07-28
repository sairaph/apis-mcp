---
title: insights_resources_payment_evaluation_shipping
page_id: schema-insights-resources-payment-evaluation-shipping-64e666d1
path: schemas
description: Shipping details attached to this payment.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_shipping

Shipping details attached to this payment.

```yaml
{"title": "InsightsResourcesPaymentEvaluationShipping", "required": ["address"], "type": "object", "properties": {"address": {"$ref": "#/components/schemas/insights_resources_payment_evaluation_address"}, "name": {"maxLength": 5000, "type": "string", "description": "Shipping name.", "nullable": true}, "phone": {"maxLength": 5000, "type": "string", "description": "Shipping phone number.", "nullable": true}}, "description": "Shipping details attached to this payment.", "x-expandableFields": ["address"]}
```
