---
title: api-shield_parameter_schemas_definition
page_id: schema-api-shield-parameter-schemas-definition-0d91bf38
path: schemas
description: An operation schema object containing a response.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_parameter_schemas_definition

An operation schema object containing a response.

```yaml
{"description": "An operation schema object containing a response.", "type": "object", "properties": {"parameters": {"description": "An array containing the learned parameter schemas.", "example": [{"description": "Sufficient requests have been observed for this parameter to provide high confidence in this parameter schema.", "in": "path", "name": "var1", "required": true, "schema": {"type": "integer", "maximum": 10, "minimum": 1}}], "items": {"type": "object"}, "readOnly": true, "type": "array"}, "responses": {"description": "An empty response object. This field is required to yield a valid operation schema.", "nullable": true, "readOnly": true, "type": "object"}}, "example": {"parameters": [{"description": "Sufficient requests have been observed for this parameter to provide high confidence in this parameter schema.", "in": "path", "name": "var1", "required": true, "schema": {"type": "integer", "maximum": 10, "minimum": 1}}], "responses": null}, "readOnly": true}
```
