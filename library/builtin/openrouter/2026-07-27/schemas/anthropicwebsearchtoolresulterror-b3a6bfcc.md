---
title: AnthropicWebSearchToolResultError
page_id: schema-anthropicwebsearchtoolresulterror-b3a6bfcc
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicWebSearchToolResultError

```yaml
{"example": {"error_code": "unavailable", "type": "web_search_tool_result_error"}, "properties": {"error_code": {"enum": ["invalid_tool_input", "unavailable", "max_uses_exceeded", "too_many_requests", "query_too_long", "request_too_large"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "type": {"enum": ["web_search_tool_result_error"], "type": "string"}}, "required": ["type", "error_code"], "type": "object"}
```
