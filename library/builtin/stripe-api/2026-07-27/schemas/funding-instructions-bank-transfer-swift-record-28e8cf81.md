---
title: funding_instructions_bank_transfer_swift_record
page_id: schema-funding-instructions-bank-transfer-swift-record-28e8cf81
path: schemas
description: SWIFT Records contain U.S. bank account details per the SWIFT format.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# funding_instructions_bank_transfer_swift_record

SWIFT Records contain U.S. bank account details per the SWIFT format.

```yaml
{"title": "FundingInstructionsBankTransferSwiftRecord", "required": ["account_holder_address", "account_holder_name", "account_number", "account_type", "bank_address", "bank_name", "swift_code"], "type": "object", "properties": {"account_holder_address": {"$ref": "#/components/schemas/address"}, "account_holder_name": {"maxLength": 5000, "type": "string", "description": "The account holder name"}, "account_number": {"maxLength": 5000, "type": "string", "description": "The account number"}, "account_type": {"maxLength": 5000, "type": "string", "description": "The account type"}, "bank_address": {"$ref": "#/components/schemas/address"}, "bank_name": {"maxLength": 5000, "type": "string", "description": "The bank name"}, "swift_code": {"maxLength": 5000, "type": "string", "description": "The SWIFT code"}}, "description": "SWIFT Records contain U.S. bank account details per the SWIFT format.", "x-expandableFields": ["account_holder_address", "bank_address"]}
```
