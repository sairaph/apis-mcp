---
title: ImageGenTextChunkEvent
page_id: schema-imagegentextchunkevent-9421ed5b
path: schemas
description: Emitted when a text chunk becomes available during streaming generation of text-based formats (e.g. SVG)
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImageGenTextChunkEvent

Emitted when a text chunk becomes available during streaming generation of text-based formats (e.g. SVG)

```yaml
{"description": "Emitted when a text chunk becomes available during streaming generation of text-based formats (e.g. SVG)", "example": {"phase": "content", "text": "<svg xmlns=\"http://www.w3.org/2000/svg\">", "type": "image_generation.text_chunk"}, "properties": {"phase": {"description": "The generation phase this chunk belongs to. `content` is the renderable output; `reasoning` and `draft` are intermediate provider phases.", "enum": ["content", "reasoning", "draft"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "text": {"description": "A text fragment of the image being generated (e.g. partial SVG markup)", "type": "string"}, "type": {"description": "The event type", "enum": ["image_generation.text_chunk"], "type": "string"}}, "required": ["type", "text", "phase"], "type": "object"}
```
