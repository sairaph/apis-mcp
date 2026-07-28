---
title: v2.error.entity_type_not_supported_in_country
page_id: schema-v2-error-entity-type-not-supported-in-country-c1dd066c
path: schemas
description: The `identity.entity_type` value is not supported in a given `identity.country`.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.entity_type_not_supported_in_country

The `identity.entity_type` value is not supported in a given `identity.country`.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "doc_url", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["entity_type_not_supported_in_country"]}, "doc_url": {"type": "string", "description": "A URL to more information about the error reported"}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "The `identity.entity_type` value is not supported in a given `identity.country`."}
```
