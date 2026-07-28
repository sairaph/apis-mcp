---
title: AgentMessageItem
page_id: schema-agentmessageitem-944026cd
path: schemas
description: A message routed between agents in a multi-agent session
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AgentMessageItem

A message routed between agents in a multi-agent session

```yaml
{"additionalProperties": {}, "description": "A message routed between agents in a multi-agent session", "example": {"author": "/root/worker", "content": [{"text": "Task complete.", "type": "input_text"}], "recipient": "/root", "type": "agent_message"}, "properties": {"agent": {"additionalProperties": {}, "properties": {"agent_name": {"type": "string"}}, "required": ["agent_name"], "type": ["object", "null"]}, "author": {"type": "string"}, "content": {"items": {"oneOf": [{"$ref": "#/components/schemas/InputText"}, {"allOf": [{"$ref": "#/components/schemas/InputImage"}, {"properties": {}, "type": "object"}], "description": "Image input content item", "example": {"detail": "auto", "image_url": "https://example.com/image.jpg", "type": "input_image"}}, {"additionalProperties": {}, "properties": {"encrypted_content": {"type": "string"}, "type": {"enum": ["encrypted_content"], "type": "string"}}, "required": ["type", "encrypted_content"], "type": "object"}]}, "type": "array"}, "id": {"type": ["string", "null"]}, "recipient": {"type": "string"}, "type": {"enum": ["agent_message"], "type": "string"}}, "required": ["type", "author", "recipient", "content"], "type": "object"}
```
