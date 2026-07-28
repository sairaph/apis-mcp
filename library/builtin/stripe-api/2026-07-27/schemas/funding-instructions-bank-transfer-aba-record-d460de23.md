---
title: funding_instructions_bank_transfer_aba_record
page_id: schema-funding-instructions-bank-transfer-aba-record-d460de23
path: schemas
description: ABA Records contain U.S. bank account details per the ABA format.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# funding_instructions_bank_transfer_aba_record

ABA Records contain U.S. bank account details per the ABA format.

```yaml
{"title": "FundingInstructionsBankTransferABARecord", "required": ["account_holder_address", "account_holder_name", "account_number", "account_type", "bank_address", "bank_name", "routing_number"], "type": "object", "properties": {"account_holder_address": {"$ref": "#/components/schemas/address"}, "account_holder_name": {"maxLength": 5000, "type": "string", "description": "The account holder name"}, "account_number": {"maxLength": 5000, "type": "string", "description": "The ABA account number"}, "account_type": {"maxLength": 5000, "type": "string", "description": "The account type"}, "bank_address": {"$ref": "#/components/schemas/address"}, "bank_name": {"maxLength": 5000, "type": "string", "description": "The bank name"}, "routing_number": {"maxLength": 5000, "type": "string", "description": "The ABA routing number"}}, "description": "ABA Records contain U.S. bank account details per the ABA format.", "x-expandableFields": ["account_holder_address", "bank_address"]}
```
