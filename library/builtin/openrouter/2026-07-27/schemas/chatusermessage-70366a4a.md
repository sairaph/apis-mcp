---
title: ChatUserMessage
page_id: schema-chatusermessage-70366a4a
path: schemas
description: User message
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatUserMessage

User message

```yaml
{"description": "User message", "example": {"content": "What is the capital of France?", "role": "user"}, "properties": {"content": {"anyOf": [{"type": "string"}, {"items": {"$ref": "#/components/schemas/ChatContentItems"}, "type": "array"}], "description": "User message content", "example": "What is the capital of France?"}, "name": {"description": "Optional name for the user", "example": "User", "type": "string"}, "role": {"enum": ["user"], "type": "string"}}, "required": ["role", "content"], "type": "object"}
```
