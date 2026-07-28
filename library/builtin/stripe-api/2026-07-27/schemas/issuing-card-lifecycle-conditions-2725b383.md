---
title: issuing_card_lifecycle_conditions
page_id: schema-issuing-card-lifecycle-conditions-2725b383
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_card_lifecycle_conditions

```yaml
{"title": "IssuingCardLifecycleConditions", "required": ["payment_count"], "type": "object", "properties": {"payment_count": {"type": "integer", "description": "The card is automatically cancelled when it makes this number of non-zero payment authorizations and transactions. The count includes penny authorizations, but doesn't include non-payment actions, such as authorization advice."}}, "description": "", "x-expandableFields": []}
```
