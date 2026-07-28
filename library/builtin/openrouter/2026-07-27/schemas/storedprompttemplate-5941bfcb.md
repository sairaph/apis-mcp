---
title: StoredPromptTemplate
page_id: schema-storedprompttemplate-5941bfcb
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# StoredPromptTemplate

```yaml
{"example": {"id": "prompt-abc123", "variables": {"name": "John"}}, "properties": {"id": {"type": "string"}, "variables": {"additionalProperties": {"anyOf": [{"type": "string"}, {"$ref": "#/components/schemas/InputText"}, {"$ref": "#/components/schemas/InputImage"}, {"$ref": "#/components/schemas/InputFile"}]}, "type": ["object", "null"]}}, "required": ["id"], "type": ["object", "null"]}
```
