---
title: cash_balance
page_id: schema-cash-balance-23bb84c6
path: schemas
description: A customer's `Cash balance` represents real funds. Customers can add funds to their cash balance by sending a bank transfer. These funds can be used for payment and can eventually be paid out to your bank account.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# cash_balance

A customer's `Cash balance` represents real funds. Customers can add funds to their cash balance by sending a bank transfer. These funds can be used for payment and can eventually be paid out to your bank account.

```yaml
{"title": "cash_balance", "required": ["customer", "livemode", "object", "settings"], "type": "object", "properties": {"available": {"type": "object", "additionalProperties": {"type": "integer"}, "description": "A hash of all cash balances available to this customer. You cannot delete a customer with any cash balances, even if the balance is 0. Amounts are represented in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).", "nullable": true}, "customer": {"maxLength": 5000, "type": "string", "description": "The ID of the customer whose cash balance this object represents."}, "customer_account": {"maxLength": 5000, "type": "string", "description": "The ID of an Account representing a customer whose cash balance this object represents.", "nullable": true}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["cash_balance"]}, "settings": {"$ref": "#/components/schemas/customer_balance_customer_balance_settings"}}, "description": "A customer's `Cash balance` represents real funds. Customers can add funds to their cash balance by sending a bank transfer. These funds can be used for payment and can eventually be paid out to your bank account.", "x-expandableFields": ["settings"], "x-resourceId": "cash_balance"}
```
