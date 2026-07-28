---
title: OutputItemImageGenerationCall
page_id: schema-outputitemimagegenerationcall-135c2d92
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputItemImageGenerationCall

```yaml
{"example": {"id": "imagegen-abc123", "result": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg==", "status": "completed", "type": "image_generation_call"}, "properties": {"id": {"type": "string"}, "result": {"default": null, "type": ["string", "null"]}, "status": {"$ref": "#/components/schemas/ImageGenerationStatus"}, "type": {"enum": ["image_generation_call"], "type": "string"}}, "required": ["type", "id", "status"], "type": "object"}
```
