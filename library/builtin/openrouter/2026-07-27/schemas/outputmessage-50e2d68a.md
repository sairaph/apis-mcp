---
title: OutputMessage
page_id: schema-outputmessage-50e2d68a
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputMessage

```yaml
{"example": {"content": [{"text": "Hello! How can I help you today?", "type": "output_text"}], "id": "msg-abc123", "role": "assistant", "status": "completed", "type": "message"}, "properties": {"content": {"items": {"anyOf": [{"$ref": "#/components/schemas/ResponseOutputText"}, {"$ref": "#/components/schemas/OpenAIResponsesRefusalContent"}]}, "type": "array"}, "id": {"type": "string"}, "phase": {"anyOf": [{"enum": ["commentary"], "type": "string"}, {"enum": ["final_answer"], "type": "string"}, {"type": "null"}], "description": "The phase of an assistant message. Use `commentary` for an intermediate assistant message and `final_answer` for the final assistant message. For follow-up requests with models like `gpt-5.3-codex` and later, preserve and resend phase on all assistant messages. Omitting it can degrade performance. Not used for user messages."}, "role": {"enum": ["assistant"], "type": "string"}, "status": {"anyOf": [{"enum": ["completed"], "type": "string"}, {"enum": ["incomplete"], "type": "string"}, {"enum": ["in_progress"], "type": "string"}]}, "type": {"enum": ["message"], "type": "string"}}, "required": ["id", "role", "type", "content"], "type": "object"}
```
