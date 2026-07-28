---
title: payment_method_us_bank_account_blocked
page_id: schema-payment-method-us-bank-account-blocked-0b4fe591
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_us_bank_account_blocked

```yaml
{"title": "payment_method_us_bank_account_blocked", "type": "object", "properties": {"network_code": {"type": "string", "description": "The ACH network code that resulted in this block.", "nullable": true, "enum": ["R02", "R03", "R04", "R05", "R07", "R08", "R10", "R11", "R16", "R20", "R29", "R31"]}, "reason": {"type": "string", "description": "The reason why this PaymentMethod's fingerprint has been blocked", "nullable": true, "enum": ["bank_account_closed", "bank_account_frozen", "bank_account_invalid_details", "bank_account_restricted", "bank_account_unusable", "debit_not_authorized", "tokenized_account_number_deactivated"]}}, "description": "", "x-expandableFields": []}
```
