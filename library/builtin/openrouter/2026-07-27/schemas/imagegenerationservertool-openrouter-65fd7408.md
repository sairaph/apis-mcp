---
title: ImageGenerationServerTool_OpenRouter
page_id: schema-imagegenerationservertool-openrouter-65fd7408
path: schemas
description: 'OpenRouter built-in server tool: generates images from text prompts using an image generation model'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImageGenerationServerTool_OpenRouter

OpenRouter built-in server tool: generates images from text prompts using an image generation model

```yaml
{"description": "OpenRouter built-in server tool: generates images from text prompts using an image generation model", "example": {"parameters": {"model": "openai/gpt-5-image", "quality": "high", "size": "1024x1024"}, "type": "openrouter:image_generation"}, "properties": {"parameters": {"$ref": "#/components/schemas/ImageGenerationServerToolConfig"}, "type": {"enum": ["openrouter:image_generation"], "type": "string"}}, "required": ["type"], "type": "object"}
```
