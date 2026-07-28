---
title: v2.error.person_percent_ownership_invalid
page_id: schema-v2-error-person-percent-ownership-invalid-0fc891a9
path: schemas
description: Error returned when relationship.owner is set to true but the ownership percentage is set to 0%.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.person_percent_ownership_invalid

Error returned when relationship.owner is set to true but the ownership percentage is set to 0%.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["person_percent_ownership_invalid"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "Error returned when relationship.owner is set to true but the ownership percentage is set to 0%."}
```
