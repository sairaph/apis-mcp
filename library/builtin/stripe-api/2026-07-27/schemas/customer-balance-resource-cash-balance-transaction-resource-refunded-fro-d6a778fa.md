---
title: customer_balance_resource_cash_balance_transaction_resource_refunded_from_payment_transaction
page_id: schema-customer-balance-resource-cash-balance-transaction-resource-refunded-fro-d6a778fa
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# customer_balance_resource_cash_balance_transaction_resource_refunded_from_payment_transaction

```yaml
{"title": "CustomerBalanceResourceCashBalanceTransactionResourceRefundedFromPaymentTransaction", "required": ["refund"], "type": "object", "properties": {"refund": {"description": "The [Refund](https://docs.stripe.com/api/refunds/object) that moved these funds into the customer's cash balance.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/refund"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/refund"}]}}}, "description": "", "x-expandableFields": ["refund"]}
```
