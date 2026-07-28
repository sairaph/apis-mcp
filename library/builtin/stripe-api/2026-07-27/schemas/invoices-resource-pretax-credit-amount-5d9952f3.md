---
title: invoices_resource_pretax_credit_amount
page_id: schema-invoices-resource-pretax-credit-amount-5d9952f3
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoices_resource_pretax_credit_amount

```yaml
{"title": "InvoicesResourcePretaxCreditAmount", "required": ["amount", "type"], "type": "object", "properties": {"amount": {"type": "integer", "description": "The amount, in cents (or local equivalent), of the pretax credit amount."}, "credit_balance_transaction": {"description": "The credit balance transaction that was applied to get this pretax credit amount.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/billing.credit_balance_transaction"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/billing.credit_balance_transaction"}]}}, "discount": {"description": "The discount that was applied to get this pretax credit amount.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/discount"}, {"$ref": "#/components/schemas/deleted_discount"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/discount"}, {"$ref": "#/components/schemas/deleted_discount"}]}}, "type": {"type": "string", "description": "Type of the pretax credit amount referenced.", "enum": ["credit_balance_transaction", "discount"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["credit_balance_transaction", "discount"]}
```
