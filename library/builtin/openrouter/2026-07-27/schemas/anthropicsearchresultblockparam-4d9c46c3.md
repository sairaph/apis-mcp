---
title: AnthropicSearchResultBlockParam
page_id: schema-anthropicsearchresultblockparam-4d9c46c3
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicSearchResultBlockParam

```yaml
{"example": {"content": [{"text": "Result content", "type": "text"}], "source": "example_source", "title": "Example Result", "type": "search_result"}, "properties": {"cache_control": {"$ref": "#/components/schemas/AnthropicCacheControlDirective"}, "citations": {"properties": {"enabled": {"type": "boolean"}}, "type": "object"}, "content": {"items": {"$ref": "#/components/schemas/AnthropicTextBlockParam"}, "type": "array"}, "source": {"type": "string"}, "title": {"type": "string"}, "type": {"enum": ["search_result"], "type": "string"}}, "required": ["type", "source", "title", "content"], "type": "object"}
```
