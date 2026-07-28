---
title: SpecialEffectsVideosAgentRequest
page_id: schema-specialeffectsvideosagentrequest-7e6dc29a
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# SpecialEffectsVideosAgentRequest

```yaml
{"type": "object", "properties": {"agent_id": {"type": "string", "description": "Agent ID: `vidu_template_agent`.", "enum": ["vidu_template_agent"]}, "request_id": {"type": "string", "description": "User-defined unique ID; used to distinguish requests. Auto-generated if omitted."}, "messages": {"type": "array", "description": "Session message body.", "items": {"type": "object", "properties": {"role": {"type": "string", "description": "User input role: `user`", "example": "user", "default": "user", "enum": ["user"]}, "content": {"type": "array", "description": "Content list.", "items": {"oneOf": [{"title": "text", "type": "object", "properties": {"type": {"type": "string", "description": "Specifies that this content is text.", "enum": ["text"]}, "text": {"type": "string", "description": "User text input."}}, "required": ["type", "text"]}, {"title": "image_url", "type": "object", "properties": {"type": {"type": "string", "description": "Specifies that this content is an image URL.", "enum": ["image_url"]}, "image_url": {"type": "string", "format": "uri", "description": "Image URL input."}}, "required": ["type", "image_url"]}]}}}, "required": ["role", "content"]}}, "custom_variables": {"type": "object", "description": "Agent extension parameters.", "properties": {"template": {"type": "string", "description": "Effect template: `french_kiss`, `bodyshake`, or `sexy_me`.", "enum": ["french_kiss", "bodyshake", "sexy_me"]}}, "required": ["template"]}}, "required": ["agent_id", "messages"]}
```
