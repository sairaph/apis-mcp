---
title: customer_balance_resource_cash_balance_transaction_resource_funded_transaction_resource_bank_transfer
page_id: schema-customer-balance-resource-cash-balance-transaction-resource-funded-trans-59b53cac
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# customer_balance_resource_cash_balance_transaction_resource_funded_transaction_resource_bank_transfer

```yaml
{"title": "CustomerBalanceResourceCashBalanceTransactionResourceFundedTransactionResourceBankTransfer", "required": ["type"], "type": "object", "properties": {"eu_bank_transfer": {"$ref": "#/components/schemas/customer_balance_resource_cash_balance_transaction_resource_funded_transaction_resource_bank_transfer_resource_eu_bank_transfer"}, "gb_bank_transfer": {"$ref": "#/components/schemas/customer_balance_resource_cash_balance_transaction_resource_funded_transaction_resource_bank_transfer_resource_gb_bank_transfer"}, "jp_bank_transfer": {"$ref": "#/components/schemas/customer_balance_resource_cash_balance_transaction_resource_funded_transaction_resource_bank_transfer_resource_jp_bank_transfer"}, "reference": {"maxLength": 5000, "type": "string", "description": "The user-supplied reference field on the bank transfer.", "nullable": true}, "type": {"type": "string", "description": "The funding method type used to fund the customer balance. Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, `mx_bank_transfer`, or `us_bank_transfer`.", "enum": ["eu_bank_transfer", "gb_bank_transfer", "jp_bank_transfer", "mx_bank_transfer", "us_bank_transfer"], "x-stripeBypassValidation": true}, "us_bank_transfer": {"$ref": "#/components/schemas/customer_balance_resource_cash_balance_transaction_resource_funded_transaction_resource_bank_transfer_resource_us_bank_transfer"}}, "description": "", "x-expandableFields": ["eu_bank_transfer", "gb_bank_transfer", "jp_bank_transfer", "us_bank_transfer"]}
```
