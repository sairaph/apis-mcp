---
title: ImageGenPartialImageEvent
page_id: schema-imagegenpartialimageevent-10e91677
path: schemas
description: Emitted when a partial image becomes available during streaming generation
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImageGenPartialImageEvent

Emitted when a partial image becomes available during streaming generation

```yaml
{"description": "Emitted when a partial image becomes available during streaming generation", "example": {"b64_json": "<base64-encoded-partial-image>", "partial_image_index": 0, "type": "image_generation.partial_image"}, "properties": {"b64_json": {"description": "Base64-encoded partial image data", "type": "string"}, "partial_image_index": {"description": "0-based index indicating which partial image this is in the sequence", "type": "integer"}, "type": {"description": "The event type", "enum": ["image_generation.partial_image"], "type": "string"}}, "required": ["type", "partial_image_index", "b64_json"], "type": "object"}
```
