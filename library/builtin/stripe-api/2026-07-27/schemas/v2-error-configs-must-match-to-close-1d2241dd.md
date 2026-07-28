---
title: v2.error.configs_must_match_to_close
page_id: schema-v2-error-configs-must-match-to-close-1d2241dd
path: schemas
description: Account cannot be closed without specifying the right configurations.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.configs_must_match_to_close

Account cannot be closed without specifying the right configurations.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["applied_configurations", "code", "message", "user_message"], "type": "object", "properties": {"applied_configurations": {"type": "array", "description": "The applied configurations that should be specified.", "items": {"type": "string"}}, "code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["configs_must_match_to_close"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "Account cannot be closed without specifying the right configurations."}
```
