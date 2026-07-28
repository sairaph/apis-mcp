---
title: treasury_financial_accounts_resource_aba_record
page_id: schema-treasury-financial-accounts-resource-aba-record-af952b90
path: schemas
description: ABA Records contain U.S. bank account details per the ABA format.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_financial_accounts_resource_aba_record

ABA Records contain U.S. bank account details per the ABA format.

```yaml
{"title": "TreasuryFinancialAccountsResourceABARecord", "required": ["account_holder_name", "account_number_last4", "bank_name", "routing_number"], "type": "object", "properties": {"account_holder_name": {"maxLength": 5000, "type": "string", "description": "The name of the person or business that owns the bank account."}, "account_number": {"maxLength": 5000, "type": "string", "description": "The account number.", "nullable": true}, "account_number_last4": {"maxLength": 5000, "type": "string", "description": "The last four characters of the account number."}, "bank_name": {"maxLength": 5000, "type": "string", "description": "Name of the bank."}, "routing_number": {"maxLength": 5000, "type": "string", "description": "Routing number for the account."}}, "description": "ABA Records contain U.S. bank account details per the ABA format.", "x-expandableFields": []}
```
