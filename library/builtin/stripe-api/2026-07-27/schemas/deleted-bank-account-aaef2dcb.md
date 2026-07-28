---
title: deleted_bank_account
page_id: schema-deleted-bank-account-aaef2dcb
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# deleted_bank_account

```yaml
{"title": "DeletedBankAccount", "required": ["deleted", "id", "object"], "type": "object", "properties": {"currency": {"maxLength": 5000, "type": "string", "description": "Three-letter [ISO code for the currency](https://stripe.com/docs/payouts) paid out to the bank account.", "nullable": true}, "deleted": {"type": "boolean", "description": "Always true for a deleted object", "enum": [true]}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["bank_account"]}}, "description": "", "x-expandableFields": []}
```
