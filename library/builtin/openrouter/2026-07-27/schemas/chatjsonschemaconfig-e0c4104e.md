---
title: ChatJsonSchemaConfig
page_id: schema-chatjsonschemaconfig-e0c4104e
path: schemas
description: JSON Schema configuration object
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatJsonSchemaConfig

JSON Schema configuration object

```yaml
{"description": "JSON Schema configuration object", "example": {"description": "A mathematical response", "name": "math_response", "schema": {"properties": {"answer": {"type": "number"}}, "required": ["answer"], "type": "object"}, "strict": true}, "properties": {"description": {"description": "Schema description for the model", "example": "A mathematical response", "type": "string"}, "name": {"description": "Schema name (a-z, A-Z, 0-9, underscores, dashes, max 64 chars)", "example": "math_response", "maxLength": 64, "type": "string"}, "schema": {"additionalProperties": {}, "description": "JSON Schema object", "example": {"properties": {"answer": {"type": "number"}}, "required": ["answer"], "type": "object"}, "type": "object"}, "strict": {"description": "Enable strict schema adherence", "example": false, "type": ["boolean", "null"]}}, "required": ["name"], "type": "object"}
```
