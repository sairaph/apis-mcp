---
title: EndpointsMetadata
page_id: schema-endpointsmetadata-c7df20b8
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# EndpointsMetadata

```yaml
{"example": {"available": [{"model": "openai/gpt-4o", "provider": "OpenAI", "selected": true}], "total": 3}, "properties": {"available": {"items": {"$ref": "#/components/schemas/EndpointInfo"}, "type": "array"}, "total": {"type": "integer"}}, "required": ["total", "available"], "type": "object"}
```
