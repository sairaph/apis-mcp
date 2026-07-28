---
title: InputMessageItem
page_id: schema-inputmessageitem-2b4a321c
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# InputMessageItem

```yaml
{"example": {"content": [{"text": "Hello, how are you?", "type": "input_text"}], "id": "msg-abc123", "role": "user", "type": "message"}, "properties": {"content": {"items": {"oneOf": [{"$ref": "#/components/schemas/InputText"}, {"allOf": [{"$ref": "#/components/schemas/InputImage"}, {"properties": {}, "type": "object"}], "description": "Image input content item", "example": {"detail": "auto", "image_url": "https://example.com/image.jpg", "type": "input_image"}}, {"$ref": "#/components/schemas/InputFile"}, {"$ref": "#/components/schemas/InputAudio"}, {"$ref": "#/components/schemas/InputVideo"}]}, "type": ["array", "null"]}, "id": {"type": "string"}, "role": {"anyOf": [{"enum": ["user"], "type": "string"}, {"enum": ["system"], "type": "string"}, {"enum": ["developer"], "type": "string"}]}, "type": {"enum": ["message"], "type": "string"}}, "required": ["role"], "type": "object"}
```
