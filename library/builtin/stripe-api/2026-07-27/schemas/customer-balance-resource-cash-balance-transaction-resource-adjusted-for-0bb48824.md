---
title: customer_balance_resource_cash_balance_transaction_resource_adjusted_for_overdraft
page_id: schema-customer-balance-resource-cash-balance-transaction-resource-adjusted-for-0bb48824
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# customer_balance_resource_cash_balance_transaction_resource_adjusted_for_overdraft

```yaml
{"title": "CustomerBalanceResourceCashBalanceTransactionResourceAdjustedForOverdraft", "required": ["balance_transaction", "linked_transaction"], "type": "object", "properties": {"balance_transaction": {"description": "The [Balance Transaction](https://docs.stripe.com/api/balance_transactions/object) that corresponds to funds taken out of your Stripe balance.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/balance_transaction"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/balance_transaction"}]}}, "linked_transaction": {"description": "The [Cash Balance Transaction](https://docs.stripe.com/api/cash_balance_transactions/object) that brought the customer balance negative, triggering the clawback of funds.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/customer_cash_balance_transaction"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/customer_cash_balance_transaction"}]}}}, "description": "", "x-expandableFields": ["balance_transaction", "linked_transaction"]}
```
