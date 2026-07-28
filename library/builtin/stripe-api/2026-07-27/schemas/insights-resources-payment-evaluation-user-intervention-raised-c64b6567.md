---
title: insights_resources_payment_evaluation_user_intervention_raised
page_id: schema-insights-resources-payment-evaluation-user-intervention-raised-c64b6567
path: schemas
description: User intervention raised event details attached to this payment evaluation
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_user_intervention_raised

User intervention raised event details attached to this payment evaluation

```yaml
{"title": "InsightsResourcesPaymentEvaluationUserInterventionRaised", "required": ["key", "type"], "type": "object", "properties": {"custom": {"$ref": "#/components/schemas/insights_resources_payment_evaluation_user_intervention_raised_custom"}, "key": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the user intervention event."}, "type": {"type": "string", "description": "Type of user intervention raised.", "enum": ["3ds", "captcha", "custom"]}}, "description": "User intervention raised event details attached to this payment evaluation", "x-expandableFields": ["custom"]}
```
