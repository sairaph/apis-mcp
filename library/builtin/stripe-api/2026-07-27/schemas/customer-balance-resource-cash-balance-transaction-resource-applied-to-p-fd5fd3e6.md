---
title: customer_balance_resource_cash_balance_transaction_resource_applied_to_payment_transaction
page_id: schema-customer-balance-resource-cash-balance-transaction-resource-applied-to-p-fd5fd3e6
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# customer_balance_resource_cash_balance_transaction_resource_applied_to_payment_transaction

```yaml
{"title": "CustomerBalanceResourceCashBalanceTransactionResourceAppliedToPaymentTransaction", "required": ["payment_intent"], "type": "object", "properties": {"payment_intent": {"description": "The [Payment Intent](https://docs.stripe.com/api/payment_intents/object) that funds were applied to.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/payment_intent"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/payment_intent"}]}}}, "description": "", "x-expandableFields": ["payment_intent"]}
```
