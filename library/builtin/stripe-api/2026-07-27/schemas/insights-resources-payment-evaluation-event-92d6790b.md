---
title: insights_resources_payment_evaluation_event
page_id: schema-insights-resources-payment-evaluation-event-92d6790b
path: schemas
description: Event reported for this payment evaluation.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_event

Event reported for this payment evaluation.

```yaml
{"title": "InsightsResourcesPaymentEvaluationEvent", "required": ["occurred_at", "type"], "type": "object", "properties": {"dispute_opened": {"$ref": "#/components/schemas/insights_resources_payment_evaluation_dispute_opened"}, "early_fraud_warning_received": {"$ref": "#/components/schemas/insights_resources_payment_evaluation_early_fraud_warning_received"}, "occurred_at": {"type": "integer", "description": "Timestamp when the event occurred.", "format": "unix-time"}, "refunded": {"$ref": "#/components/schemas/insights_resources_payment_evaluation_refunded"}, "type": {"type": "string", "description": "Indicates the type of event attached to the payment evaluation.", "enum": ["dispute_opened", "early_fraud_warning_received", "refunded", "user_intervention_raised", "user_intervention_resolved"]}, "user_intervention_raised": {"$ref": "#/components/schemas/insights_resources_payment_evaluation_user_intervention_raised"}, "user_intervention_resolved": {"$ref": "#/components/schemas/insights_resources_payment_evaluation_user_intervention_resolved"}}, "description": "Event reported for this payment evaluation.", "x-expandableFields": ["dispute_opened", "early_fraud_warning_received", "refunded", "user_intervention_raised", "user_intervention_resolved"]}
```
