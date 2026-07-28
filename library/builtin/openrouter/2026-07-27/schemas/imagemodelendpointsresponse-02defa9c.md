---
title: ImageModelEndpointsResponse
page_id: schema-imagemodelendpointsresponse-02defa9c
path: schemas
description: The full per-endpoint records for an image model.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImageModelEndpointsResponse

The full per-endpoint records for an image model.

```yaml
{"description": "The full per-endpoint records for an image model.", "example": {"endpoints": [{"allowed_passthrough_parameters": [], "pricing": [{"billable": "output_image", "cost_usd": 0.05, "unit": "image"}], "provider_name": "Bytedance", "provider_slug": "bytedance", "provider_tag": "bytedance", "supported_parameters": {"resolution": {"type": "enum", "values": ["1K", "2K", "4K"]}}, "supports_streaming": false}], "id": "bytedance-seed/seedream-4.5"}, "properties": {"endpoints": {"items": {"$ref": "#/components/schemas/ImageEndpoint"}, "type": "array"}, "id": {"description": "Model slug", "example": "bytedance-seed/seedream-4.5", "type": "string"}}, "required": ["id", "endpoints"], "type": "object"}
```
