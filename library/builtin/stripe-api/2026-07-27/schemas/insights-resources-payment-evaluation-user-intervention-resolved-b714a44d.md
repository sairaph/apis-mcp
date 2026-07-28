---
title: insights_resources_payment_evaluation_user_intervention_resolved
page_id: schema-insights-resources-payment-evaluation-user-intervention-resolved-b714a44d
path: schemas
description: User Intervention Resolved Event details attached to this payment evaluation
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_user_intervention_resolved

User Intervention Resolved Event details attached to this payment evaluation

```yaml
{"title": "InsightsResourcesPaymentEvaluationUserInterventionResolved", "required": ["key"], "type": "object", "properties": {"key": {"maxLength": 5000, "type": "string", "description": "Unique ID of this intervention. Use this to provide the result."}, "outcome": {"type": "string", "description": "Result of the intervention if it has been completed.", "nullable": true, "enum": ["abandoned", "failed", "passed"]}}, "description": "User Intervention Resolved Event details attached to this payment evaluation", "x-expandableFields": []}
```
