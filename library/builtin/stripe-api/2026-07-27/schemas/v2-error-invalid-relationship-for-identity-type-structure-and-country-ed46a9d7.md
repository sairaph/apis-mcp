---
title: v2.error.invalid_relationship_for_identity_type_structure_and_country
page_id: schema-v2-error-invalid-relationship-for-identity-type-structure-and-country-ed46a9d7
path: schemas
description: Some relationships are specific to type, structure, and country.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.invalid_relationship_for_identity_type_structure_and_country

Some relationships are specific to type, structure, and country.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["invalid_relationship_for_identity_type_structure_and_country"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "Some relationships are specific to type, structure, and country."}
```
