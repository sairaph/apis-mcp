---
title: insights_resources_payment_evaluation_money_movement_details
page_id: schema-insights-resources-payment-evaluation-money-movement-details-e8962dfa
path: schemas
description: Money Movement details attached to this payment.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_money_movement_details

Money Movement details attached to this payment.

```yaml
{"title": "InsightsResourcesPaymentEvaluationMoneyMovementDetails", "required": ["money_movement_type"], "type": "object", "properties": {"card": {"description": "Describes card money movement details for the payment evaluation.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/insights_resources_payment_evaluation_money_movement_card"}]}, "money_movement_type": {"type": "string", "description": "Describes the type of money movement. Currently only `card` is supported.", "enum": ["card"]}}, "description": "Money Movement details attached to this payment.", "x-expandableFields": ["card"]}
```
