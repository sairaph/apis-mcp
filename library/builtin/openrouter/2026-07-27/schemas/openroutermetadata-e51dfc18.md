---
title: OpenRouterMetadata
page_id: schema-openroutermetadata-e51dfc18
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OpenRouterMetadata

```yaml
{"example": {"attempt": 1, "endpoints": {"available": [{"model": "openai/gpt-4o", "provider": "OpenAI", "selected": true}], "total": 1}, "is_byok": false, "region": "iad", "requested": "openai/gpt-4o", "strategy": "direct", "summary": "available=1, selected=OpenAI"}, "properties": {"attempt": {"type": "integer"}, "attempts": {"items": {"$ref": "#/components/schemas/RouterAttempt"}, "type": "array"}, "endpoints": {"$ref": "#/components/schemas/EndpointsMetadata"}, "is_byok": {"type": "boolean"}, "params": {"$ref": "#/components/schemas/RouterParams"}, "pipeline": {"items": {"$ref": "#/components/schemas/PipelineStage"}, "type": "array"}, "region": {"type": ["string", "null"]}, "requested": {"type": "string"}, "strategy": {"$ref": "#/components/schemas/RoutingStrategy"}, "summary": {"type": "string"}}, "required": ["requested", "strategy", "region", "summary", "attempt", "is_byok", "endpoints"], "type": "object"}
```
