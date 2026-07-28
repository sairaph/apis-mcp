---
title: v2.error.token_must_be_created_with_publishable_key
page_id: schema-v2-error-token-must-be-created-with-publishable-key-3232f99b
path: schemas
description: Token must be created with publishable key.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.token_must_be_created_with_publishable_key

Token must be created with publishable key.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["token_must_be_created_with_publishable_key"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Token must be created with publishable key."}
```
