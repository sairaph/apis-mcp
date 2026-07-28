---
title: payment_method_details_nz_bank_account
page_id: schema-payment-method-details-nz-bank-account-2ededf60
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_nz_bank_account

```yaml
{"title": "payment_method_details_nz_bank_account", "required": ["bank_code", "bank_name", "branch_code", "last4"], "type": "object", "properties": {"account_holder_name": {"maxLength": 5000, "type": "string", "description": "The name on the bank account. Only present if the account holder name is different from the name of the authorized signatory collected in the PaymentMethod’s billing details.", "nullable": true}, "bank_code": {"maxLength": 5000, "type": "string", "description": "The numeric code for the bank account's bank."}, "bank_name": {"maxLength": 5000, "type": "string", "description": "The name of the bank."}, "branch_code": {"maxLength": 5000, "type": "string", "description": "The numeric code for the bank account's bank branch."}, "expected_debit_date": {"maxLength": 5000, "type": "string", "description": "Estimated date to debit the customer's bank account. A date string in YYYY-MM-DD format."}, "last4": {"maxLength": 5000, "type": "string", "description": "Last four digits of the bank account number."}, "suffix": {"maxLength": 5000, "type": "string", "description": "The suffix of the bank account number.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
