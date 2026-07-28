---
title: funding_instructions_bank_transfer_sort_code_record
page_id: schema-funding-instructions-bank-transfer-sort-code-record-9bd7953f
path: schemas
description: Sort Code Records contain U.K. bank account details per the sort code format.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# funding_instructions_bank_transfer_sort_code_record

Sort Code Records contain U.K. bank account details per the sort code format.

```yaml
{"title": "FundingInstructionsBankTransferSortCodeRecord", "required": ["account_holder_address", "account_holder_name", "account_number", "bank_address", "sort_code"], "type": "object", "properties": {"account_holder_address": {"$ref": "#/components/schemas/address"}, "account_holder_name": {"maxLength": 5000, "type": "string", "description": "The name of the person or business that owns the bank account"}, "account_number": {"maxLength": 5000, "type": "string", "description": "The account number"}, "bank_address": {"$ref": "#/components/schemas/address"}, "sort_code": {"maxLength": 5000, "type": "string", "description": "The six-digit sort code"}}, "description": "Sort Code Records contain U.K. bank account details per the sort code format.", "x-expandableFields": ["account_holder_address", "bank_address"]}
```
