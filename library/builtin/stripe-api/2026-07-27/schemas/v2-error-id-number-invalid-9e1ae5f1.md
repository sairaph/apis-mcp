---
title: v2.error.id_number_invalid
page_id: schema-v2-error-id-number-invalid-9e1ae5f1
path: schemas
description: Provided ID number is of the wrong format for the given type.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.id_number_invalid

Provided ID number is of the wrong format for the given type.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["id_number_invalid"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "Provided ID number is of the wrong format for the given type."}
```
