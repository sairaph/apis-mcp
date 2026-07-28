---
title: insights_resources_payment_evaluation_signal_v2
page_id: schema-insights-resources-payment-evaluation-signal-v2-398ac9fd
path: schemas
description: A payment evaluation signal with evaluated_at, risk_level, and score fields.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_signal_v2

A payment evaluation signal with evaluated_at, risk_level, and score fields.

```yaml
{"title": "InsightsResourcesPaymentEvaluationSignalV2", "required": ["evaluated_at", "risk_level", "score"], "type": "object", "properties": {"evaluated_at": {"type": "integer", "description": "The time when this signal was evaluated.", "format": "unix-time"}, "risk_level": {"type": "string", "description": "Risk level of this signal, based on the score.", "enum": ["elevated", "highest", "low", "normal", "not_assessed", "unknown"]}, "score": {"type": "number", "description": "Score for this signal. Possible values for evaluated payments are between 0 and 100. The value is returned with two decimal places and higher scores indicate a higher likelihood of the signal being true. A score of -1 is returned when a model evaluation was not performed, such as requests from incomplete integrations."}}, "description": "A payment evaluation signal with evaluated_at, risk_level, and score fields.", "x-expandableFields": []}
```
