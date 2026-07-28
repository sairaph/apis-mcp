---
title: ImageModelsListResponse
page_id: schema-imagemodelslistresponse-024b7511
path: schemas
description: List of image generation models.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImageModelsListResponse

List of image generation models.

```yaml
{"description": "List of image generation models.", "example": {"data": [{"architecture": {"input_modalities": ["text"], "output_modalities": ["image"]}, "created": 1692901234, "description": "A text-to-image model.", "endpoints": "/api/v1/images/models/bytedance-seed/seedream-4.5/endpoints", "id": "bytedance-seed/seedream-4.5", "name": "Seedream 4.5", "supported_parameters": {"resolution": {"type": "enum", "values": ["1K", "2K", "4K"]}}, "supports_streaming": false}]}, "properties": {"data": {"items": {"$ref": "#/components/schemas/ImageModelListItem"}, "type": "array"}}, "required": ["data"], "type": "object"}
```
