---
title: v2.error.account_not_yet_compatible_with_v2
page_id: schema-v2-error-account-not-yet-compatible-with-v2-82584b82
path: schemas
description: Account is not yet compatible with V2 APIs.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.account_not_yet_compatible_with_v2

Account is not yet compatible with V2 APIs.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["account_not_yet_compatible_with_v2"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Account is not yet compatible with V2 APIs."}
```
