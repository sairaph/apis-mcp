---
title: AnthropicToolSearchContent
page_id: schema-anthropictoolsearchcontent-e209d6b0
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicToolSearchContent

```yaml
{"discriminator": {"mapping": {"tool_search_tool_result_error": "#/components/schemas/AnthropicToolSearchResultError", "tool_search_tool_search_result": "#/components/schemas/AnthropicToolSearchResult"}, "propertyName": "type"}, "example": {"tool_references": [{"tool_name": "my_tool", "type": "tool_reference"}], "type": "tool_search_tool_search_result"}, "oneOf": [{"$ref": "#/components/schemas/AnthropicToolSearchResultError"}, {"$ref": "#/components/schemas/AnthropicToolSearchResult"}]}
```
