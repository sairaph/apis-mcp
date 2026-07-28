---
title: ChatFormatJsonSchemaConfig
page_id: schema-chatformatjsonschemaconfig-34f01c48
path: schemas
description: JSON Schema response format for structured outputs
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatFormatJsonSchemaConfig

JSON Schema response format for structured outputs

```yaml
{"description": "JSON Schema response format for structured outputs", "example": {"json_schema": {"name": "math_response", "schema": {"properties": {"answer": {"type": "number"}}, "required": ["answer"], "type": "object"}}, "type": "json_schema"}, "properties": {"json_schema": {"$ref": "#/components/schemas/ChatJsonSchemaConfig"}, "type": {"enum": ["json_schema"], "type": "string"}}, "required": ["type", "json_schema"], "type": "object"}
```
