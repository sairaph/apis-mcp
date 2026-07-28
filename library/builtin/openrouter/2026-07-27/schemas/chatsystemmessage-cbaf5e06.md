---
title: ChatSystemMessage
page_id: schema-chatsystemmessage-cbaf5e06
path: schemas
description: System message for setting behavior
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatSystemMessage

System message for setting behavior

```yaml
{"description": "System message for setting behavior", "example": {"content": "You are a helpful assistant.", "name": "Assistant Config", "role": "system"}, "properties": {"content": {"anyOf": [{"type": "string"}, {"items": {"$ref": "#/components/schemas/ChatContentText"}, "type": "array"}], "description": "System message content", "example": "You are a helpful assistant."}, "name": {"description": "Optional name for the system message", "example": "Assistant Config", "type": "string"}, "role": {"enum": ["system"], "type": "string"}}, "required": ["role", "content"], "type": "object"}
```
