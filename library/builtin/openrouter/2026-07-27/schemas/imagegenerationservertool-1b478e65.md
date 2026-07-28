---
title: ImageGenerationServerTool
page_id: schema-imagegenerationservertool-1b478e65
path: schemas
description: Image generation tool configuration
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImageGenerationServerTool

Image generation tool configuration

```yaml
{"description": "Image generation tool configuration", "example": {"quality": "high", "type": "image_generation"}, "properties": {"background": {"enum": ["transparent", "opaque", "auto"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "input_fidelity": {"enum": ["high", "low", null], "type": ["string", "null"], "x-speakeasy-unknown-values": "allow"}, "input_image_mask": {"properties": {"file_id": {"type": "string"}, "image_url": {"type": "string"}}, "type": "object"}, "model": {"type": "string"}, "moderation": {"enum": ["auto", "low"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "output_compression": {"type": "integer"}, "output_format": {"enum": ["png", "webp", "jpeg"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "partial_images": {"type": "integer"}, "quality": {"enum": ["low", "medium", "high", "auto"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "size": {"type": "string"}, "type": {"enum": ["image_generation"], "type": "string"}}, "required": ["type"], "type": "object"}
```
