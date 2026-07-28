---
title: funding_instructions_bank_transfer_zengin_record
page_id: schema-funding-instructions-bank-transfer-zengin-record-a5772e4e
path: schemas
description: Zengin Records contain Japan bank account details per the Zengin format.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# funding_instructions_bank_transfer_zengin_record

Zengin Records contain Japan bank account details per the Zengin format.

```yaml
{"title": "FundingInstructionsBankTransferZenginRecord", "required": ["account_holder_address", "bank_address"], "type": "object", "properties": {"account_holder_address": {"$ref": "#/components/schemas/address"}, "account_holder_name": {"maxLength": 5000, "type": "string", "description": "The account holder name", "nullable": true}, "account_number": {"maxLength": 5000, "type": "string", "description": "The account number", "nullable": true}, "account_type": {"maxLength": 5000, "type": "string", "description": "The bank account type. In Japan, this can only be `futsu` or `toza`.", "nullable": true}, "bank_address": {"$ref": "#/components/schemas/address"}, "bank_code": {"maxLength": 5000, "type": "string", "description": "The bank code of the account", "nullable": true}, "bank_name": {"maxLength": 5000, "type": "string", "description": "The bank name of the account", "nullable": true}, "branch_code": {"maxLength": 5000, "type": "string", "description": "The branch code of the account", "nullable": true}, "branch_name": {"maxLength": 5000, "type": "string", "description": "The branch name of the account", "nullable": true}}, "description": "Zengin Records contain Japan bank account details per the Zengin format.", "x-expandableFields": ["account_holder_address", "bank_address"]}
```
