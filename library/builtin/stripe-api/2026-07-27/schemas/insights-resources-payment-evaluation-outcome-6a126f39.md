---
title: insights_resources_payment_evaluation_outcome
page_id: schema-insights-resources-payment-evaluation-outcome-6a126f39
path: schemas
description: Outcome details for this payment evaluation.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_outcome

Outcome details for this payment evaluation.

```yaml
{"title": "InsightsResourcesPaymentEvaluationOutcome", "required": ["type"], "type": "object", "properties": {"merchant_blocked": {"$ref": "#/components/schemas/insights_resources_payment_evaluation_merchant_blocked"}, "payment_intent_id": {"maxLength": 5000, "type": "string", "description": "The PaymentIntent ID associated with the payment evaluation."}, "rejected": {"$ref": "#/components/schemas/insights_resources_payment_evaluation_rejected"}, "succeeded": {"$ref": "#/components/schemas/insights_resources_payment_evaluation_succeeded"}, "type": {"type": "string", "description": "Indicates the outcome of the payment evaluation.", "enum": ["failed", "merchant_blocked", "rejected", "succeeded"]}}, "description": "Outcome details for this payment evaluation.", "x-expandableFields": ["merchant_blocked", "rejected", "succeeded"]}
```
