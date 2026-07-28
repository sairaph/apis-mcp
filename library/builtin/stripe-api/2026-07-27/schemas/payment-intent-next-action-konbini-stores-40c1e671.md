---
title: payment_intent_next_action_konbini_stores
page_id: schema-payment-intent-next-action-konbini-stores-40c1e671
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_next_action_konbini_stores

```yaml
{"title": "payment_intent_next_action_konbini_stores", "type": "object", "properties": {"familymart": {"description": "FamilyMart instruction details.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_intent_next_action_konbini_familymart"}]}, "lawson": {"description": "Lawson instruction details.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_intent_next_action_konbini_lawson"}]}, "ministop": {"description": "Ministop instruction details.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_intent_next_action_konbini_ministop"}]}, "seicomart": {"description": "Seicomart instruction details.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_intent_next_action_konbini_seicomart"}]}}, "description": "", "x-expandableFields": ["familymart", "lawson", "ministop", "seicomart"]}
```
