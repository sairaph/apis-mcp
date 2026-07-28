---
title: v2.error.bgn_bank_accounts_unsupported
page_id: schema-v2-error-bgn-bank-accounts-unsupported-6c960f5c
path: schemas
description: Creating accounts with the BGN currency is no longer supported, as Bulgaria is now using the Euro as of 2026-01-01.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.bgn_bank_accounts_unsupported

Creating accounts with the BGN currency is no longer supported, as Bulgaria is now using the Euro as of 2026-01-01.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "doc_url", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["bgn_bank_accounts_unsupported"]}, "doc_url": {"type": "string", "description": "A URL to more information about the error reported"}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "Creating accounts with the BGN currency is no longer supported, as Bulgaria is now using the Euro as of 2026-01-01."}
```
