---
title: api-shield_operation_schema_fragment
page_id: schema-api-shield-operation-schema-fragment-218537ce
path: schemas
description: An OpenAPI operation object fragment containing schema information for an operation. May include parameter definitions, request body specifications, and a component schema extension.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_operation_schema_fragment

An OpenAPI operation object fragment containing schema information for an operation. May include parameter definitions, request body specifications, and a component schema extension.

```yaml
{"description": "An OpenAPI operation object fragment containing schema information for an operation. May include parameter definitions, request body specifications, and a component schema extension.", "type": "object", "properties": {"parameters": {"description": "OpenAPI parameter objects describing path, query, header, or cookie parameters.", "items": {"additionalProperties": true, "type": "object"}, "type": "array"}, "requestBody": {"description": "OpenAPI request body object describing the expected request payload.", "additionalProperties": true, "type": "object"}}, "nullable": true}
```
