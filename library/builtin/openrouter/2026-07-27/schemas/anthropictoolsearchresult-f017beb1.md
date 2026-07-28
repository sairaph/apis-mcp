---
title: AnthropicToolSearchResult
page_id: schema-anthropictoolsearchresult-f017beb1
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicToolSearchResult

```yaml
{"example": {"tool_references": [{"tool_name": "my_tool", "type": "tool_reference"}], "type": "tool_search_tool_search_result"}, "properties": {"tool_references": {"items": {"$ref": "#/components/schemas/AnthropicToolReference"}, "type": "array"}, "type": {"enum": ["tool_search_tool_search_result"], "type": "string"}}, "required": ["tool_references", "type"], "type": "object"}
```
