---
title: FunctionObject
page_id: schema-functionobject-55c6052c
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# FunctionObject

```yaml
{"type": "object", "properties": {"name": {"type": "string", "description": "The name of the function to be called. Must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum length of 64.", "minLength": 1, "maxLength": 64, "pattern": "^[a-zA-Z0-9_-]+$"}, "description": {"type": "string", "description": "A description of what the function does, used by the model to choose when and how to call the function."}, "parameters": {"$ref": "#/components/schemas/FunctionParameters"}}, "required": ["name", "description", "parameters"]}
```
