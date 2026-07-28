---
title: v2.error.v1_token_invalid_in_v2
page_id: schema-v2-error-v1-token-invalid-in-v2-1d1181a6
path: schemas
description: A v1 token ID is passed in v2 APIs.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.v1_token_invalid_in_v2

A v1 token ID is passed in v2 APIs.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["v1_token_invalid_in_v2"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "A v1 token ID is passed in v2 APIs."}
```
