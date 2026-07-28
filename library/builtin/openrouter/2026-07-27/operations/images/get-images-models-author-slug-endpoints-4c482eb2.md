---
title: List endpoints for an image model
page_id: operation-get-images-models-author-slug-endpoints-3877ad11
path: operations/images
description: 'Returns the full per-endpoint records for an image model: each endpoint''s definitive supported parameters, pricing, and passthrough allowlist.'
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /images/models/{author}/{slug}/endpoints
operation_ids:
    - listImageModelEndpoints
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List endpoints for an image model

`GET /images/models/{author}/{slug}/endpoints`

Operation ID: `listImageModelEndpoints`

Returns the full per-endpoint records for an image model: each endpoint's definitive supported parameters, pricing, and passthrough allowlist.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Returns the full per-endpoint records for an image model: each endpoint's definitive supported parameters, pricing, and passthrough allowlist.", "operationId": "listImageModelEndpoints", "parameters": [{"description": "Model author/organization", "in": "path", "name": "author", "required": true, "schema": {"description": "Model author/organization", "example": "bytedance-seed", "type": "string"}}, {"description": "Model slug", "in": "path", "name": "slug", "required": true, "schema": {"description": "Model slug", "example": "seedream-4.5", "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"endpoints": [{"allowed_passthrough_parameters": [], "pricing": [{"billable": "output_image", "cost_usd": 0.05, "unit": "image"}], "provider_name": "Bytedance", "provider_slug": "bytedance", "provider_tag": "bytedance", "supported_parameters": {"resolution": {"type": "enum", "values": ["1K", "2K", "4K"]}}, "supports_streaming": false}], "id": "bytedance-seed/seedream-4.5"}, "schema": {"$ref": "#/components/schemas/ImageModelEndpointsResponse"}}}, "description": "The full per-endpoint records for an image model"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "List endpoints for an image model", "tags": ["Images"], "x-speakeasy-name-override": "listModelEndpoints"}
```
