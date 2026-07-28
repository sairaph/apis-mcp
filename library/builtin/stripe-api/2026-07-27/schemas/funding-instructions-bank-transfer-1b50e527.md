---
title: funding_instructions_bank_transfer
page_id: schema-funding-instructions-bank-transfer-1b50e527
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# funding_instructions_bank_transfer

```yaml
{"title": "FundingInstructionsBankTransfer", "required": ["country", "financial_addresses", "type"], "type": "object", "properties": {"country": {"maxLength": 5000, "type": "string", "description": "The country of the bank account to fund"}, "financial_addresses": {"type": "array", "description": "A list of financial addresses that can be used to fund a particular balance", "items": {"$ref": "#/components/schemas/funding_instructions_bank_transfer_financial_address"}}, "type": {"type": "string", "description": "The bank_transfer type", "enum": ["eu_bank_transfer", "jp_bank_transfer"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["financial_addresses"]}
```
