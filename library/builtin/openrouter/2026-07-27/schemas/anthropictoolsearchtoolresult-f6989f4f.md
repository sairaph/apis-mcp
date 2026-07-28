---
title: AnthropicToolSearchToolResult
page_id: schema-anthropictoolsearchtoolresult-f6989f4f
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicToolSearchToolResult

```yaml
{"example": {"content": {"tool_references": [{"tool_name": "my_tool", "type": "tool_reference"}], "type": "tool_search_tool_search_result"}, "tool_use_id": "srvtoolu_01abc", "type": "tool_search_tool_result"}, "properties": {"content": {"$ref": "#/components/schemas/AnthropicToolSearchContent"}, "tool_use_id": {"type": "string"}, "type": {"enum": ["tool_search_tool_result"], "type": "string"}}, "required": ["type", "content", "tool_use_id"], "type": "object"}
```
