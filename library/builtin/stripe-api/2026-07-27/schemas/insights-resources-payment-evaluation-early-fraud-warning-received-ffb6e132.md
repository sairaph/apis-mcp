---
title: insights_resources_payment_evaluation_early_fraud_warning_received
page_id: schema-insights-resources-payment-evaluation-early-fraud-warning-received-ffb6e132
path: schemas
description: Early Fraud Warning Received event details attached to this payment evaluation.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_early_fraud_warning_received

Early Fraud Warning Received event details attached to this payment evaluation.

```yaml
{"title": "InsightsResourcesPaymentEvaluationEarlyFraudWarningReceived", "required": ["fraud_type"], "type": "object", "properties": {"fraud_type": {"type": "string", "description": "The type of fraud labeled by the issuer.", "enum": ["made_with_lost_card", "made_with_stolen_card", "other", "unauthorized_use_of_card"]}}, "description": "Early Fraud Warning Received event details attached to this payment evaluation.", "x-expandableFields": []}
```
