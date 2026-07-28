---
title: OpenAIResponseInputMessageItem
page_id: schema-openairesponseinputmessageitem-453df3e0
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OpenAIResponseInputMessageItem

```yaml
{"example": {"content": [{"text": "Hello, how are you?", "type": "input_text"}], "id": "msg-abc123", "role": "user", "type": "message"}, "properties": {"content": {"items": {"discriminator": {"mapping": {"input_audio": "#/components/schemas/InputAudio", "input_file": "#/components/schemas/InputFile", "input_image": "#/components/schemas/InputImage", "input_text": "#/components/schemas/InputText"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/InputText"}, {"$ref": "#/components/schemas/InputImage"}, {"$ref": "#/components/schemas/InputFile"}, {"$ref": "#/components/schemas/InputAudio"}]}, "type": "array"}, "id": {"type": "string"}, "role": {"anyOf": [{"enum": ["user"], "type": "string"}, {"enum": ["system"], "type": "string"}, {"enum": ["developer"], "type": "string"}]}, "type": {"enum": ["message"], "type": "string"}}, "required": ["id", "role", "content"], "type": "object"}
```
