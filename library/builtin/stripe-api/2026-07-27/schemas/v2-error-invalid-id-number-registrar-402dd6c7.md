---
title: v2.error.invalid_id_number_registrar
page_id: schema-v2-error-invalid-id-number-registrar-402dd6c7
path: schemas
description: The `identity.business_details.id_numbers.registrar` value is an invalid DE registrar.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.invalid_id_number_registrar

The `identity.business_details.id_numbers.registrar` value is an invalid DE registrar.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "doc_url", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["invalid_id_number_registrar"]}, "doc_url": {"type": "string", "description": "A URL to more information about the error reported"}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "The `identity.business_details.id_numbers.registrar` value is an invalid DE registrar."}
```
