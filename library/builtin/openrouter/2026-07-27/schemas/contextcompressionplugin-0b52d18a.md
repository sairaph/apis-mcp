---
title: ContextCompressionPlugin
page_id: schema-contextcompressionplugin-0b52d18a
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ContextCompressionPlugin

```yaml
{"example": {"enabled": true, "engine": "middle-out", "id": "context-compression"}, "properties": {"enabled": {"description": "Set to false to disable the context-compression plugin for this request. Defaults to true.", "type": "boolean"}, "engine": {"$ref": "#/components/schemas/ContextCompressionEngine"}, "id": {"enum": ["context-compression"], "type": "string"}}, "required": ["id"], "type": "object"}
```
