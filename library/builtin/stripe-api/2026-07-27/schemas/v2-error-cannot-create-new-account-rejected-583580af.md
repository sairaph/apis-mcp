---
title: v2.error.cannot_create_new_account_rejected
page_id: schema-v2-error-cannot-create-new-account-rejected-583580af
path: schemas
description: Platform is in a rejected state and cannot create connected accounts.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.cannot_create_new_account_rejected

Platform is in a rejected state and cannot create connected accounts.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["cannot_create_new_account_rejected"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Platform is in a rejected state and cannot create connected accounts."}
```
