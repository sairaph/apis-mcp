---
title: customer_balance_resource_cash_balance_transaction_resource_funded_transaction_resource_bank_transfer_resource_us_bank_transfer
page_id: schema-customer-balance-resource-cash-balance-transaction-resource-funded-trans-499dc4e4
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# customer_balance_resource_cash_balance_transaction_resource_funded_transaction_resource_bank_transfer_resource_us_bank_transfer

```yaml
{"title": "CustomerBalanceResourceCashBalanceTransactionResourceFundedTransactionResourceBankTransferResourceUsBankTransfer", "type": "object", "properties": {"network": {"type": "string", "description": "The banking network used for this funding.", "enum": ["ach", "domestic_wire_us", "swift"]}, "sender_name": {"maxLength": 5000, "type": "string", "description": "The full name of the sender, as supplied by the sending bank.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
