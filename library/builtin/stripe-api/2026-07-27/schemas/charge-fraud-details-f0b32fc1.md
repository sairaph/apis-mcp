---
title: charge_fraud_details
page_id: schema-charge-fraud-details-f0b32fc1
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# charge_fraud_details

```yaml
{"title": "ChargeFraudDetails", "type": "object", "properties": {"stripe_report": {"maxLength": 5000, "type": "string", "description": "Assessments from Stripe. If set, the value is `fraudulent`."}, "user_report": {"maxLength": 5000, "type": "string", "description": "Assessments reported by you. If set, possible values of are `safe` and `fraudulent`."}}, "description": "", "x-expandableFields": []}
```
