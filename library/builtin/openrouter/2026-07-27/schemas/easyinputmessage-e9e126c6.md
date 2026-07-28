---
title: EasyInputMessage
page_id: schema-easyinputmessage-e9e126c6
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# EasyInputMessage

```yaml
{"example": {"content": "What is the weather today?", "role": "user"}, "properties": {"content": {"anyOf": [{"items": {"oneOf": [{"$ref": "#/components/schemas/InputText"}, {"allOf": [{"$ref": "#/components/schemas/InputImage"}, {"properties": {}, "type": "object"}], "description": "Image input content item", "example": {"detail": "auto", "image_url": "https://example.com/image.jpg", "type": "input_image"}}, {"$ref": "#/components/schemas/InputFile"}, {"$ref": "#/components/schemas/InputAudio"}, {"$ref": "#/components/schemas/InputVideo"}]}, "type": "array"}, {"type": "string"}, {"type": "null"}]}, "phase": {"anyOf": [{"enum": ["commentary"], "type": "string"}, {"enum": ["final_answer"], "type": "string"}, {"type": "null"}], "description": "The phase of an assistant message. Use `commentary` for an intermediate assistant message and `final_answer` for the final assistant message. For follow-up requests with models like `gpt-5.3-codex` and later, preserve and resend phase on all assistant messages. Omitting it can degrade performance. Not used for user messages.", "example": "final_answer"}, "role": {"anyOf": [{"enum": ["user"], "type": "string"}, {"enum": ["system"], "type": "string"}, {"enum": ["assistant"], "type": "string"}, {"enum": ["developer"], "type": "string"}]}, "type": {"enum": ["message"], "type": "string"}}, "required": ["role"], "type": "object"}
```
