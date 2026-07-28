---
title: ChatContentImage
page_id: schema-chatcontentimage-d19b550f
path: schemas
description: Image content part for vision models
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatContentImage

Image content part for vision models

```yaml
{"description": "Image content part for vision models", "example": {"image_url": {"detail": "auto", "url": "https://example.com/image.jpg"}, "type": "image_url"}, "properties": {"image_url": {"properties": {"detail": {"description": "Image detail level for vision models. `original` is an OpenRouter extension (not in the OpenAI Chat Completions spec) requesting true original-resolution media; it is downgraded to `high` for providers that lack an original-resolution tier.", "enum": ["auto", "low", "high", "original"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "url": {"description": "URL of the image (data: URLs supported)", "type": "string"}}, "required": ["url"], "type": "object"}, "type": {"enum": ["image_url"], "type": "string"}}, "required": ["type", "image_url"], "type": "object"}
```
