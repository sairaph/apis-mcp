---
title: ChatDeveloperMessage
page_id: schema-chatdevelopermessage-2454fa7b
path: schemas
description: Developer message
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatDeveloperMessage

Developer message

```yaml
{"description": "Developer message", "example": {"content": "This is a message from the developer.", "role": "developer"}, "properties": {"content": {"anyOf": [{"type": "string"}, {"items": {"$ref": "#/components/schemas/ChatContentText"}, "type": "array"}], "description": "Developer message content", "example": "This is a message from the developer."}, "name": {"description": "Optional name for the developer message", "example": "Developer", "type": "string"}, "role": {"enum": ["developer"], "type": "string"}}, "required": ["role", "content"], "type": "object"}
```
