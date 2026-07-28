---
title: payment_intent_next_action_boleto
page_id: schema-payment-intent-next-action-boleto-25f1710d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_next_action_boleto

```yaml
{"title": "payment_intent_next_action_boleto", "type": "object", "properties": {"expires_at": {"type": "integer", "description": "The timestamp after which the boleto expires.", "format": "unix-time", "nullable": true}, "hosted_voucher_url": {"maxLength": 5000, "type": "string", "description": "The URL to the hosted boleto voucher page, which allows customers to view the boleto voucher.", "nullable": true}, "number": {"maxLength": 5000, "type": "string", "description": "The boleto number.", "nullable": true}, "pdf": {"maxLength": 5000, "type": "string", "description": "The URL to the downloadable boleto voucher PDF.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
