---
title: v2.error.structure_incompatible_for_entity_type_country
page_id: schema-v2-error-structure-incompatible-for-entity-type-country-c2877e95
path: schemas
description: The `business_details.structure` value is not valid for `identity.country` and `identity.entity_type`.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.structure_incompatible_for_entity_type_country

The `business_details.structure` value is not valid for `identity.country` and `identity.entity_type`.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["structure_incompatible_for_entity_type_country"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "The `business_details.structure` value is not valid for `identity.country` and `identity.entity_type`."}
```
