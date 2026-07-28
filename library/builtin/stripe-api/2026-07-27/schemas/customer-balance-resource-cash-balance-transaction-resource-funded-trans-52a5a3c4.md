---
title: customer_balance_resource_cash_balance_transaction_resource_funded_transaction_resource_bank_transfer_resource_eu_bank_transfer
page_id: schema-customer-balance-resource-cash-balance-transaction-resource-funded-trans-52a5a3c4
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# customer_balance_resource_cash_balance_transaction_resource_funded_transaction_resource_bank_transfer_resource_eu_bank_transfer

```yaml
{"title": "CustomerBalanceResourceCashBalanceTransactionResourceFundedTransactionResourceBankTransferResourceEuBankTransfer", "type": "object", "properties": {"bic": {"maxLength": 5000, "type": "string", "description": "The BIC of the bank of the sender of the funding.", "nullable": true}, "iban_last4": {"maxLength": 5000, "type": "string", "description": "The last 4 digits of the IBAN of the sender of the funding.", "nullable": true}, "sender_name": {"maxLength": 5000, "type": "string", "description": "The full name of the sender, as supplied by the sending bank.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
