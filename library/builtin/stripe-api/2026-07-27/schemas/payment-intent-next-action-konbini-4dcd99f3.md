---
title: payment_intent_next_action_konbini
page_id: schema-payment-intent-next-action-konbini-4dcd99f3
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_next_action_konbini

```yaml
{"title": "payment_intent_next_action_konbini", "required": ["expires_at", "stores"], "type": "object", "properties": {"expires_at": {"type": "integer", "description": "The timestamp at which the pending Konbini payment expires.", "format": "unix-time"}, "hosted_voucher_url": {"maxLength": 5000, "type": "string", "description": "The URL for the Konbini payment instructions page, which allows customers to view and print a Konbini voucher.", "nullable": true}, "stores": {"$ref": "#/components/schemas/payment_intent_next_action_konbini_stores"}}, "description": "", "x-expandableFields": ["stores"]}
```
