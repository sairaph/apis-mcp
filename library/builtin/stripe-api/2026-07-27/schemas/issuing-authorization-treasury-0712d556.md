---
title: issuing_authorization_treasury
page_id: schema-issuing-authorization-treasury-0712d556
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_authorization_treasury

```yaml
{"title": "IssuingAuthorizationTreasury", "required": ["received_credits", "received_debits"], "type": "object", "properties": {"received_credits": {"type": "array", "description": "The array of [ReceivedCredits](https://docs.stripe.com/api/treasury/received_credits) associated with this authorization", "items": {"maxLength": 5000, "type": "string"}}, "received_debits": {"type": "array", "description": "The array of [ReceivedDebits](https://docs.stripe.com/api/treasury/received_debits) associated with this authorization", "items": {"maxLength": 5000, "type": "string"}}, "transaction": {"maxLength": 5000, "type": "string", "description": "The Treasury [Transaction](https://docs.stripe.com/api/treasury/transactions) associated with this authorization", "nullable": true}}, "description": "", "x-expandableFields": []}
```
