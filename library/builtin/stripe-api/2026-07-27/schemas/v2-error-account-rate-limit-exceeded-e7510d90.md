---
title: v2.error.account_rate_limit_exceeded
page_id: schema-v2-error-account-rate-limit-exceeded-e7510d90
path: schemas
description: Account cannot exceed a configured concurrency rate limit on updates.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.account_rate_limit_exceeded

Account cannot exceed a configured concurrency rate limit on updates.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message", "type"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["account_rate_limit_exceeded"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "type": {"type": "string", "description": "The type of error returned", "enum": ["rate_limit"]}}, "description": "Information about the error that occurred"}}, "description": "Account cannot exceed a configured concurrency rate limit on updates."}
```
