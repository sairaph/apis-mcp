---
title: ChatToolCall
page_id: schema-chattoolcall-aa661491
path: schemas
description: Tool call made by the assistant
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatToolCall

Tool call made by the assistant

```yaml
{"description": "Tool call made by the assistant", "example": {"function": {"arguments": "{\"location\": \"Boston, MA\"}", "name": "get_current_weather"}, "id": "call_abc123", "type": "function"}, "properties": {"function": {"properties": {"arguments": {"description": "Function arguments as JSON string", "type": "string"}, "name": {"description": "Function name to call", "type": "string"}}, "required": ["name", "arguments"], "type": "object"}, "id": {"description": "Tool call identifier", "type": "string"}, "type": {"enum": ["function"], "type": "string"}}, "required": ["id", "type", "function"], "type": "object"}
```
