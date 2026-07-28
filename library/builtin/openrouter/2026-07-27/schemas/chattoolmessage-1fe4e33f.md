---
title: ChatToolMessage
page_id: schema-chattoolmessage-1fe4e33f
path: schemas
description: Tool response message
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatToolMessage

Tool response message

```yaml
{"description": "Tool response message", "example": {"content": "The weather in San Francisco is 72°F and sunny.", "role": "tool", "tool_call_id": "call_abc123"}, "properties": {"content": {"anyOf": [{"type": "string"}, {"items": {"$ref": "#/components/schemas/ChatContentItems"}, "type": "array"}], "description": "Tool response content", "example": "The weather in San Francisco is 72°F and sunny."}, "role": {"enum": ["tool"], "type": "string"}, "tool_call_id": {"description": "ID of the assistant message tool call this message responds to", "example": "call_abc123", "type": "string"}}, "required": ["role", "content", "tool_call_id"], "type": "object"}
```
