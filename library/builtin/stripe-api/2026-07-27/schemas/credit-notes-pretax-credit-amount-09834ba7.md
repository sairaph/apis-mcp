---
title: credit_notes_pretax_credit_amount
page_id: schema-credit-notes-pretax-credit-amount-09834ba7
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# credit_notes_pretax_credit_amount

```yaml
{"title": "CreditNotesPretaxCreditAmount", "required": ["amount", "type"], "type": "object", "properties": {"amount": {"type": "integer", "description": "The amount, in cents (or local equivalent), of the pretax credit amount."}, "credit_balance_transaction": {"description": "The credit balance transaction that was applied to get this pretax credit amount.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/billing.credit_balance_transaction"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/billing.credit_balance_transaction"}]}}, "discount": {"description": "The discount that was applied to get this pretax credit amount.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/discount"}, {"$ref": "#/components/schemas/deleted_discount"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/discount"}, {"$ref": "#/components/schemas/deleted_discount"}]}}, "type": {"type": "string", "description": "Type of the pretax credit amount referenced.", "enum": ["credit_balance_transaction", "discount"]}}, "description": "", "x-expandableFields": ["credit_balance_transaction", "discount"]}
```
