---
title: ImageConfig
page_id: schema-imageconfig-96799a97
path: schemas
description: Provider-specific image configuration options. Keys and values vary by model/provider. See https://openrouter.ai/docs/guides/overview/multimodal/image-generation for more details.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImageConfig

Provider-specific image configuration options. Keys and values vary by model/provider. See https://openrouter.ai/docs/guides/overview/multimodal/image-generation for more details.

```yaml
{"additionalProperties": {"anyOf": [{"type": "string"}, {"format": "double", "type": "number"}, {"items": {}, "type": "array"}]}, "description": "Provider-specific image configuration options. Keys and values vary by model/provider. See https://openrouter.ai/docs/guides/overview/multimodal/image-generation for more details.", "example": {"aspect_ratio": "16:9", "quality": "high"}, "type": "object"}
```
