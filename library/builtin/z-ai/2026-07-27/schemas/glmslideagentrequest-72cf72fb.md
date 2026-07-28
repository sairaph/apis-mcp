---
title: GlmSlideAgentRequest
page_id: schema-glmslideagentrequest-72cf72fb
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# GlmSlideAgentRequest

```yaml
{"type": "object", "properties": {"agent_id": {"type": "string", "description": "Agent ID.", "enum": ["slides_glm_agent"]}, "stream": {"type": "boolean", "default": true, "description": "False for sync calls (default). True for streaming."}, "conversation_id": {"type": "string", "description": "Conversation Id."}, "request_id": {"type": "string", "description": "User-defined unique ID; used to distinguish requests. Auto-generated if omitted."}, "messages": {"type": "array", "description": "Message body.", "items": {"type": "object", "properties": {"role": {"type": "string", "description": "User input role: `user`", "example": "user", "default": "user", "enum": ["user"]}, "content": {"type": "array", "description": "Content list.", "items": {"oneOf": [{"title": "text", "type": "object", "properties": {"type": {"type": "string", "description": "Specifies that this content is text.", "enum": ["text"]}, "text": {"type": "string", "description": "User text input."}}, "required": ["type", "text"]}]}}}, "required": ["role", "content"]}}}, "required": ["agent_id", "messages"]}
```
