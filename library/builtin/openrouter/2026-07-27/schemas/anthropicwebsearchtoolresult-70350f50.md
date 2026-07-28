---
title: AnthropicWebSearchToolResult
page_id: schema-anthropicwebsearchtoolresult-70350f50
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicWebSearchToolResult

```yaml
{"example": {"caller": {"type": "direct"}, "content": [], "tool_use_id": "srvtoolu_01abc", "type": "web_search_tool_result"}, "properties": {"caller": {"$ref": "#/components/schemas/AnthropicCaller"}, "content": {"anyOf": [{"items": {"$ref": "#/components/schemas/AnthropicWebSearchResult"}, "type": "array"}, {"$ref": "#/components/schemas/AnthropicWebSearchToolResultError"}]}, "tool_use_id": {"type": "string"}, "type": {"enum": ["web_search_tool_result"], "type": "string"}}, "required": ["type", "caller", "tool_use_id", "content"], "type": "object"}
```
