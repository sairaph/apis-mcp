---
title: ChatStreamToolCall
page_id: schema-chatstreamtoolcall-8e927e44
path: schemas
description: Tool call delta for streaming responses
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatStreamToolCall

Tool call delta for streaming responses

```yaml
{"description": "Tool call delta for streaming responses", "example": {"function": {"arguments": "{\"location\": \"...\"}", "name": "get_weather"}, "id": "call_abc123", "index": 0, "type": "function"}, "properties": {"function": {"description": "Function call details", "properties": {"arguments": {"description": "Function arguments as JSON string", "example": "{\"location\": \"...\"}", "type": "string"}, "name": {"description": "Function name", "example": "get_weather", "type": "string"}}, "type": "object"}, "id": {"description": "Tool call identifier", "example": "call_abc123", "type": "string"}, "index": {"description": "Tool call index in the array", "example": 0, "type": "integer"}, "type": {"description": "Tool call type", "enum": ["function"], "example": "function", "type": "string"}}, "required": ["index"], "type": "object"}
```
