---
title: AnthropicWebFetchToolResultError
page_id: schema-anthropicwebfetchtoolresulterror-7583d2cc
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicWebFetchToolResultError

```yaml
{"example": {"error_code": "unavailable", "type": "web_fetch_tool_result_error"}, "properties": {"error_code": {"enum": ["invalid_tool_input", "url_too_long", "url_not_allowed", "url_not_accessible", "unsupported_content_type", "too_many_requests", "max_uses_exceeded", "unavailable"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "type": {"enum": ["web_fetch_tool_result_error"], "type": "string"}}, "required": ["type", "error_code"], "type": "object"}
```
