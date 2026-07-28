---
title: Get total count of available models
page_id: operation-get-models-count-3ad387cd
path: operations/models
description: Get total count of available models
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /models/count
operation_ids:
    - listModelsCount
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Get total count of available models

`GET /models/count`

Operation ID: `listModelsCount`

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"operationId": "listModelsCount", "parameters": [{"description": "Filter models by output modality. Accepts a comma-separated list of modalities (text, image, audio, embeddings) or \"all\" to include all models. Defaults to \"text\".", "in": "query", "name": "output_modalities", "required": false, "schema": {"description": "Filter models by output modality. Accepts a comma-separated list of modalities (text, image, audio, embeddings) or \"all\" to include all models. Defaults to \"text\".", "example": "text", "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": {"count": 150}}, "schema": {"$ref": "#/components/schemas/ModelsCountResponse"}}}, "description": "Returns the total count of available models"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Get total count of available models", "tags": ["Models"], "x-speakeasy-name-override": "count"}
```
