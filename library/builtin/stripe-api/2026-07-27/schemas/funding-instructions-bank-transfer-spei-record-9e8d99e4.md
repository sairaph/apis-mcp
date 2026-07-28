---
title: funding_instructions_bank_transfer_spei_record
page_id: schema-funding-instructions-bank-transfer-spei-record-9e8d99e4
path: schemas
description: SPEI Records contain Mexico bank account details per the SPEI format.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# funding_instructions_bank_transfer_spei_record

SPEI Records contain Mexico bank account details per the SPEI format.

```yaml
{"title": "FundingInstructionsBankTransferSpeiRecord", "required": ["account_holder_address", "account_holder_name", "bank_address", "bank_code", "bank_name", "clabe"], "type": "object", "properties": {"account_holder_address": {"$ref": "#/components/schemas/address"}, "account_holder_name": {"maxLength": 5000, "type": "string", "description": "The account holder name"}, "bank_address": {"$ref": "#/components/schemas/address"}, "bank_code": {"maxLength": 5000, "type": "string", "description": "The three-digit bank code"}, "bank_name": {"maxLength": 5000, "type": "string", "description": "The short banking institution name"}, "clabe": {"maxLength": 5000, "type": "string", "description": "The CLABE number"}}, "description": "SPEI Records contain Mexico bank account details per the SPEI format.", "x-expandableFields": ["account_holder_address", "bank_address"]}
```
