---
title: AnthropicToolSearchResultError
page_id: schema-anthropictoolsearchresulterror-ca007204
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicToolSearchResultError

```yaml
{"example": {"error_code": "unavailable", "error_message": null, "type": "tool_search_tool_result_error"}, "properties": {"error_code": {"$ref": "#/components/schemas/AnthropicServerToolErrorCode"}, "error_message": {"type": ["string", "null"]}, "type": {"enum": ["tool_search_tool_result_error"], "type": "string"}}, "required": ["error_code", "error_message", "type"], "type": "object"}
```
