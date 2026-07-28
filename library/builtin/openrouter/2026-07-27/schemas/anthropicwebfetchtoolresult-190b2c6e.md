---
title: AnthropicWebFetchToolResult
page_id: schema-anthropicwebfetchtoolresult-190b2c6e
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicWebFetchToolResult

```yaml
{"example": {"caller": {"type": "direct"}, "content": {"content": {"citations": null, "source": {"data": "", "media_type": "text/plain", "type": "text"}, "title": null, "type": "document"}, "retrieved_at": null, "type": "web_fetch_result", "url": "https://example.com"}, "tool_use_id": "srvtoolu_01abc", "type": "web_fetch_tool_result"}, "properties": {"caller": {"$ref": "#/components/schemas/AnthropicCaller"}, "content": {"$ref": "#/components/schemas/AnthropicWebFetchContent"}, "tool_use_id": {"type": "string"}, "type": {"enum": ["web_fetch_tool_result"], "type": "string"}}, "required": ["type", "caller", "content", "tool_use_id"], "type": "object"}
```
