---
title: ImageModelListItem
page_id: schema-imagemodellistitem-e5076c8a
path: schemas
description: A single image model in the discovery listing.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImageModelListItem

A single image model in the discovery listing.

```yaml
{"description": "A single image model in the discovery listing.", "example": {"architecture": {"input_modalities": ["text", "image"], "output_modalities": ["image"]}, "created": 1692901234, "description": "A text-to-image model.", "endpoints": "/api/v1/images/models/bytedance-seed/seedream-4.5/endpoints", "id": "bytedance-seed/seedream-4.5", "name": "Seedream 4.5", "supported_parameters": {"resolution": {"type": "enum", "values": ["1K", "2K", "4K"]}, "seed": {"type": "boolean"}}, "supports_streaming": false}, "properties": {"architecture": {"$ref": "#/components/schemas/ImageModelArchitecture"}, "created": {"description": "Unix timestamp (seconds) of when the model was created", "example": 1692901234, "type": "integer"}, "description": {"example": "A text-to-image model.", "type": "string"}, "endpoints": {"description": "Relative URL to the full per-endpoint records for this model", "example": "/api/v1/images/models/bytedance-seed/seedream-4.5/endpoints", "type": "string"}, "id": {"description": "Model slug", "example": "bytedance-seed/seedream-4.5", "type": "string"}, "name": {"description": "Display name", "example": "Seedream 4.5", "type": "string"}, "supported_parameters": {"$ref": "#/components/schemas/SupportedParameters"}, "supports_streaming": {"description": "Whether any endpoint of this model supports native SSE streaming on the dedicated Image API (i.e. `stream: true` in the request). OR across endpoints.", "example": false, "type": "boolean"}}, "required": ["id", "name", "description", "created", "architecture", "supported_parameters", "supports_streaming", "endpoints"], "type": "object"}
```
