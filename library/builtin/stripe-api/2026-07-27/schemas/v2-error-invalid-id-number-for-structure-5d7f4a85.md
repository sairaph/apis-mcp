---
title: v2.error.invalid_id_number_for_structure
page_id: schema-v2-error-invalid-id-number-for-structure-5d7f4a85
path: schemas
description: ID number is provided that is not permitted for the Identity's entity type and business structure.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.invalid_id_number_for_structure

ID number is provided that is not permitted for the Identity's entity type and business structure.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["invalid_id_number_for_structure"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "ID number is provided that is not permitted for the Identity's entity type and business structure."}
```
