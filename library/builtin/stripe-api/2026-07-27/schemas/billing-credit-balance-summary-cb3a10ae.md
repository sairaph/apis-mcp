---
title: billing.credit_balance_summary
page_id: schema-billing-credit-balance-summary-cb3a10ae
path: schemas
description: Indicates the billing credit balance for billing credits granted to a customer.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing.credit_balance_summary

Indicates the billing credit balance for billing credits granted to a customer.

```yaml
{"title": "CreditBalanceSummary", "required": ["balances", "customer", "livemode", "object"], "type": "object", "properties": {"balances": {"type": "array", "description": "The billing credit balances. One entry per credit grant currency. If a customer only has credit grants in a single currency, then this will have a single balance entry.", "items": {"$ref": "#/components/schemas/credit_balance"}}, "customer": {"description": "The customer the balance is for.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/customer"}, {"$ref": "#/components/schemas/deleted_customer"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/customer"}, {"$ref": "#/components/schemas/deleted_customer"}]}}, "customer_account": {"maxLength": 5000, "type": "string", "description": "The account the balance is for.", "nullable": true}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["billing.credit_balance_summary"]}}, "description": "Indicates the billing credit balance for billing credits granted to a customer.", "x-expandableFields": ["balances", "customer"], "x-resourceId": "billing.credit_balance_summary"}
```
