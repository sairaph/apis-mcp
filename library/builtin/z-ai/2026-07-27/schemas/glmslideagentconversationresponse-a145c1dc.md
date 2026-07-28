---
title: GlmSlideAgentConversationResponse
page_id: schema-glmslideagentconversationresponse-a145c1dc
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# GlmSlideAgentConversationResponse

```yaml
{"type": "object", "properties": {"conversation_id": {"type": "string", "description": "Conversation ID"}, "agent_id": {"type": "string", "description": "Agent ID"}, "choices": {"type": "array", "description": "Agent output.", "items": {"type": "object", "properties": {"message": {"type": "array", "items": {"type": "object", "properties": {"role": {"type": "string", "description": "Role: fixed as `assistant`."}, "content": {"type": "array", "description": "Content metadata", "items": {"type": "object", "properties": {"type": {"type": "string", "description": "Response Content type: file_url、image_url"}, "tag_cn": {"type": "string", "description": "CN Tag."}, "tag_en": {"type": "string", "description": "EN Tag."}, "file_url": {"type": "string", "description": "Output file_url content when type is file_url"}, "image_url": {"type": "string", "description": "Output image_url content when type is image_url"}}}}}}}}}}, "error": {"type": "object", "properties": {"code": {"type": "string", "description": "Error code."}, "message": {"type": "string", "description": "Error message."}}}}}
```
