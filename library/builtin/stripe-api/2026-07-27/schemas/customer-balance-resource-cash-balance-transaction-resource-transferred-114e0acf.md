---
title: customer_balance_resource_cash_balance_transaction_resource_transferred_to_balance
page_id: schema-customer-balance-resource-cash-balance-transaction-resource-transferred-114e0acf
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# customer_balance_resource_cash_balance_transaction_resource_transferred_to_balance

```yaml
{"title": "CustomerBalanceResourceCashBalanceTransactionResourceTransferredToBalance", "required": ["balance_transaction"], "type": "object", "properties": {"balance_transaction": {"description": "The [Balance Transaction](https://docs.stripe.com/api/balance_transactions/object) that corresponds to funds transferred to your Stripe balance.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/balance_transaction"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/balance_transaction"}]}}}, "description": "", "x-expandableFields": ["balance_transaction"]}
```
