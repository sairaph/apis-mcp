---
title: FormatJsonSchemaConfig
page_id: schema-formatjsonschemaconfig-d6d7ca04
path: schemas
description: JSON schema constrained response format
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FormatJsonSchemaConfig

JSON schema constrained response format

```yaml
{"description": "JSON schema constrained response format", "example": {"description": "User information schema", "name": "user_info", "schema": {"properties": {"age": {"type": "number"}, "name": {"type": "string"}}, "required": ["name"], "type": "object"}, "type": "json_schema"}, "properties": {"description": {"type": "string"}, "name": {"type": "string"}, "schema": {"additionalProperties": {}, "type": "object"}, "strict": {"type": ["boolean", "null"]}, "type": {"enum": ["json_schema"], "type": "string"}}, "required": ["type", "name", "schema"], "type": "object"}
```
