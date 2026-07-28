---
title: OutputImageGenerationServerToolItem
page_id: schema-outputimagegenerationservertoolitem-6bc236bc
path: schemas
description: An openrouter:image_generation server tool output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputImageGenerationServerToolItem

An openrouter:image_generation server tool output item

```yaml
{"description": "An openrouter:image_generation server tool output item", "example": {"id": "ig_tmp_abc123", "imageUrl": "https://example.com/image.png", "result": "https://example.com/image.png", "status": "completed", "type": "openrouter:image_generation"}, "properties": {"id": {"type": "string"}, "imageB64": {"type": "string"}, "imageUrl": {"type": "string"}, "prompt": {"description": "The prompt (possibly rewritten) that the image was generated from.", "type": "string"}, "result": {"description": "The generated image as a base64-encoded string or URL, matching OpenAI image_generation_call format", "type": ["string", "null"]}, "revisedPrompt": {"type": "string"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "type": {"enum": ["openrouter:image_generation"], "type": "string"}}, "required": ["status", "type"], "type": "object"}
```
