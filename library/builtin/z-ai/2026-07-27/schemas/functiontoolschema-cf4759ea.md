---
title: FunctionToolSchema
page_id: schema-functiontoolschema-cf4759ea
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# FunctionToolSchema

```yaml
{"type": "object", "title": "Function Call", "properties": {"type": {"type": "string", "default": "function", "enum": ["function"]}, "function": {"$ref": "#/components/schemas/FunctionObject"}}, "required": ["type", "function"], "additionalProperties": false}
```
