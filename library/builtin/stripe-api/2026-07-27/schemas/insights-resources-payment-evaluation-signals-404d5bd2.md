---
title: insights_resources_payment_evaluation_signals
page_id: schema-insights-resources-payment-evaluation-signals-404d5bd2
path: schemas
description: Collection of signals for this payment evaluation.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_signals

Collection of signals for this payment evaluation.

```yaml
{"title": "InsightsResourcesPaymentEvaluationSignals", "required": ["fraudulent_payment"], "type": "object", "properties": {"fraudulent_payment": {"$ref": "#/components/schemas/insights_resources_payment_evaluation_signal_v2"}}, "description": "Collection of signals for this payment evaluation.", "x-expandableFields": ["fraudulent_payment"]}
```
