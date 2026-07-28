---
title: ImageEndpoint
page_id: schema-imageendpoint-e7b0cec2
path: schemas
description: An endpoint that serves a given image model.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImageEndpoint

An endpoint that serves a given image model.

```yaml
{"description": "An endpoint that serves a given image model.", "example": {"allowed_passthrough_parameters": [], "pricing": [{"billable": "output_image", "cost_usd": 0.05, "unit": "image"}], "provider_name": "Bytedance", "provider_slug": "bytedance", "provider_tag": "bytedance", "supported_parameters": {"resolution": {"type": "enum", "values": ["1K", "2K", "4K"]}, "seed": {"type": "boolean"}}, "supports_streaming": false}, "properties": {"allowed_passthrough_parameters": {"description": "Provider-specific options accepted under provider.options[provider_slug].", "example": [], "items": {"type": "string"}, "type": "array"}, "pricing": {"description": "Billable pricing lines for this endpoint.", "example": [{"billable": "output_image", "cost_usd": 0.05, "unit": "image"}], "items": {"$ref": "#/components/schemas/ImagePricingEntry"}, "type": "array"}, "provider_name": {"description": "Provider display name", "example": "Bytedance", "type": "string"}, "provider_slug": {"description": "Provider slug", "example": "bytedance", "type": "string"}, "provider_tag": {"description": "Provider tag for request-side selection", "example": "bytedance", "type": ["string", "null"]}, "supported_parameters": {"allOf": [{"$ref": "#/components/schemas/SupportedParameters"}, {"description": "The definitive set of parameters this endpoint accepts for this model."}]}, "supports_streaming": {"description": "Whether this endpoint supports native SSE streaming (`stream: true` in the request).", "example": false, "type": "boolean"}}, "required": ["provider_name", "provider_slug", "provider_tag", "supported_parameters", "allowed_passthrough_parameters", "supports_streaming", "pricing"], "type": "object"}
```
