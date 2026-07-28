---
title: ImageGenerationServerToolConfig
page_id: schema-imagegenerationservertoolconfig-a78e32b8
path: schemas
description: Configuration for the openrouter:image_generation server tool. Accepts all image_config params (aspect_ratio, quality, size, background, output_format, output_compression, moderation, etc.) plus a model field.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImageGenerationServerToolConfig

Configuration for the openrouter:image_generation server tool. Accepts all image_config params (aspect_ratio, quality, size, background, output_format, output_compression, moderation, etc.) plus a model field.

```yaml
{"additionalProperties": {"anyOf": [{"type": "string"}, {"format": "double", "type": "number"}, {"items": {}, "type": "array"}]}, "description": "Configuration for the openrouter:image_generation server tool. Accepts all image_config params (aspect_ratio, quality, size, background, output_format, output_compression, moderation, etc.) plus a model field.", "example": {"aspect_ratio": "16:9", "model": "openai/gpt-5-image", "quality": "high"}, "properties": {"model": {"description": "Which image generation model to use (e.g. \"openai/gpt-5-image\"). Defaults to \"openai/gpt-5-image\".", "example": "openai/gpt-5-image", "type": "string"}}, "type": "object"}
```
