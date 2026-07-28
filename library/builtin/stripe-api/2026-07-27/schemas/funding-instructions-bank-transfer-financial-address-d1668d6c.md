---
title: funding_instructions_bank_transfer_financial_address
page_id: schema-funding-instructions-bank-transfer-financial-address-d1668d6c
path: schemas
description: FinancialAddresses contain identifying information that resolves to a FinancialAccount.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# funding_instructions_bank_transfer_financial_address

FinancialAddresses contain identifying information that resolves to a FinancialAccount.

```yaml
{"title": "FundingInstructionsBankTransferFinancialAddress", "required": ["type"], "type": "object", "properties": {"aba": {"$ref": "#/components/schemas/funding_instructions_bank_transfer_aba_record"}, "iban": {"$ref": "#/components/schemas/funding_instructions_bank_transfer_iban_record"}, "sort_code": {"$ref": "#/components/schemas/funding_instructions_bank_transfer_sort_code_record"}, "spei": {"$ref": "#/components/schemas/funding_instructions_bank_transfer_spei_record"}, "supported_networks": {"type": "array", "description": "The payment networks supported by this FinancialAddress", "items": {"type": "string", "enum": ["ach", "bacs", "domestic_wire_us", "fps", "sepa", "spei", "swift", "zengin"], "x-stripeBypassValidation": true}}, "swift": {"$ref": "#/components/schemas/funding_instructions_bank_transfer_swift_record"}, "type": {"type": "string", "description": "The type of financial address", "enum": ["aba", "iban", "sort_code", "spei", "swift", "zengin"], "x-stripeBypassValidation": true}, "zengin": {"$ref": "#/components/schemas/funding_instructions_bank_transfer_zengin_record"}}, "description": "FinancialAddresses contain identifying information that resolves to a FinancialAccount.", "x-expandableFields": ["aba", "iban", "sort_code", "spei", "swift", "zengin"]}
```
