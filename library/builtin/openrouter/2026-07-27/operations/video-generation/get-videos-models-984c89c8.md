---
title: List all video generation models
page_id: operation-get-videos-models-00270286
path: operations/video-generation
description: Returns a list of all available video generation models and their properties
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /videos/models
operation_ids:
    - listVideosModels
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List all video generation models

`GET /videos/models`

Operation ID: `listVideosModels`

Returns a list of all available video generation models and their properties

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Returns a list of all available video generation models and their properties", "operationId": "listVideosModels", "responses": {"200": {"content": {"application/json": {"example": {"data": [{"allowed_passthrough_parameters": [], "canonical_slug": "google/veo-3.1", "created": 1700000000, "description": "Google video generation model", "generate_audio": true, "id": "google/veo-3.1", "name": "Veo 3.1", "pricing_skus": {"generate": "0.50"}, "seed": null, "supported_aspect_ratios": ["16:9"], "supported_durations": [5, 8], "supported_frame_images": ["first_frame", "last_frame"], "supported_resolutions": ["720p"], "supported_sizes": null}]}, "schema": {"$ref": "#/components/schemas/VideoModelsListResponse"}}}, "description": "Returns a list of video generation models"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "List all video generation models", "tags": ["Video Generation"]}
```
