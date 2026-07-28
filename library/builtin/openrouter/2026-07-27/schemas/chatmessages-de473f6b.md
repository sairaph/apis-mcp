---
title: ChatMessages
page_id: schema-chatmessages-de473f6b
path: schemas
description: Chat completion message with role-based discrimination
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatMessages

Chat completion message with role-based discrimination

```yaml
{"description": "Chat completion message with role-based discrimination", "discriminator": {"mapping": {"assistant": "#/components/schemas/ChatAssistantMessage", "developer": "#/components/schemas/ChatDeveloperMessage", "system": "#/components/schemas/ChatSystemMessage", "tool": "#/components/schemas/ChatToolMessage", "user": "#/components/schemas/ChatUserMessage"}, "propertyName": "role"}, "example": {"content": "What is the capital of France?", "role": "user"}, "oneOf": [{"$ref": "#/components/schemas/ChatSystemMessage"}, {"$ref": "#/components/schemas/ChatUserMessage"}, {"$ref": "#/components/schemas/ChatDeveloperMessage"}, {"$ref": "#/components/schemas/ChatAssistantMessage"}, {"$ref": "#/components/schemas/ChatToolMessage"}]}
```
