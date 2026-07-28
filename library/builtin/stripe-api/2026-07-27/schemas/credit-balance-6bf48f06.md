---
title: credit_balance
page_id: schema-credit-balance-6bf48f06
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# credit_balance

```yaml
{"title": "CreditBalance", "required": ["available_balance", "ledger_balance"], "type": "object", "properties": {"available_balance": {"$ref": "#/components/schemas/billing_credit_grants_resource_amount"}, "ledger_balance": {"$ref": "#/components/schemas/billing_credit_grants_resource_amount"}}, "description": "", "x-expandableFields": ["available_balance", "ledger_balance"]}
```
