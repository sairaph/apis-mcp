---
title: funding_instructions_bank_transfer_iban_record
page_id: schema-funding-instructions-bank-transfer-iban-record-798ad344
path: schemas
description: Iban Records contain E.U. bank account details per the SEPA format.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# funding_instructions_bank_transfer_iban_record

Iban Records contain E.U. bank account details per the SEPA format.

```yaml
{"title": "FundingInstructionsBankTransferIbanRecord", "required": ["account_holder_address", "account_holder_name", "bank_address", "bic", "country", "iban"], "type": "object", "properties": {"account_holder_address": {"$ref": "#/components/schemas/address"}, "account_holder_name": {"maxLength": 5000, "type": "string", "description": "The name of the person or business that owns the bank account"}, "bank_address": {"$ref": "#/components/schemas/address"}, "bic": {"maxLength": 5000, "type": "string", "description": "The BIC/SWIFT code of the account."}, "country": {"maxLength": 5000, "type": "string", "description": "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2))."}, "iban": {"maxLength": 5000, "type": "string", "description": "The IBAN of the account."}}, "description": "Iban Records contain E.U. bank account details per the SEPA format.", "x-expandableFields": ["account_holder_address", "bank_address"]}
```
