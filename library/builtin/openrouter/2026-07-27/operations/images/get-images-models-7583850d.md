---
title: List image generation models
page_id: operation-get-images-models-f6c15378
path: operations/images
description: Lists every image generation model with its top-level supported-parameter superset and a URL to its full per-endpoint records.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /images/models
operation_ids:
    - listImageModels
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List image generation models

`GET /images/models`

Operation ID: `listImageModels`

Lists every image generation model with its top-level supported-parameter superset and a URL to its full per-endpoint records.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Lists every image generation model with its top-level supported-parameter superset and a URL to its full per-endpoint records.", "operationId": "listImageModels", "responses": {"200": {"content": {"application/json": {"example": {"data": [{"architecture": {"input_modalities": ["text"], "output_modalities": ["image"]}, "created": 1692901234, "description": "A text-to-image model.", "endpoints": "/api/v1/images/models/bytedance-seed/seedream-4.5/endpoints", "id": "bytedance-seed/seedream-4.5", "name": "Seedream 4.5", "supported_parameters": {"resolution": {"type": "enum", "values": ["1K", "2K", "4K"]}}, "supports_streaming": false}]}, "schema": {"$ref": "#/components/schemas/ImageModelsListResponse"}}}, "description": "List of image generation models"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "List image generation models", "tags": ["Images"], "x-speakeasy-name-override": "listModels"}
```
