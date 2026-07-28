---
title: v2.error.invalid_cancel_configuration
page_id: schema-v2-error-invalid-cancel-configuration-f7f07f77
path: schemas
description: The adjustment configuration is invalid for the adjustment type.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.invalid_cancel_configuration

The adjustment configuration is invalid for the adjustment type.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["invalid_cancel_configuration"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "The adjustment configuration is invalid for the adjustment type."}
```
