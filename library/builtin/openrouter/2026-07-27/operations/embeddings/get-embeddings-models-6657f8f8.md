---
title: List all embeddings models
page_id: operation-get-embeddings-models-e864cc14
path: operations/embeddings
description: Returns a list of all available embeddings models and their properties
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /embeddings/models
operation_ids:
    - listEmbeddingsModels
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List all embeddings models

`GET /embeddings/models`

Operation ID: `listEmbeddingsModels`

Returns a list of all available embeddings models and their properties

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Returns a list of all available embeddings models and their properties", "operationId": "listEmbeddingsModels", "parameters": [{"description": "Number of records to skip for pagination. When both offset and limit are omitted, the full list is returned", "in": "query", "name": "offset", "required": false, "schema": {"default": 0, "description": "Number of records to skip for pagination. When both offset and limit are omitted, the full list is returned", "example": 0, "minimum": 0, "type": ["integer", "null"]}}, {"description": "Maximum number of records to return (max 1000). When both offset and limit are omitted, the full list is returned", "in": "query", "name": "limit", "required": false, "schema": {"default": 500, "description": "Maximum number of records to return (max 1000). When both offset and limit are omitted, the full list is returned", "example": 500, "maximum": 1000, "minimum": 1, "type": "integer"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": [{"architecture": {"input_modalities": ["text"], "instruct_type": null, "modality": "text->text", "output_modalities": ["embeddings"], "tokenizer": "GPT"}, "canonical_slug": "openai/text-embedding-3-small", "context_length": 8192, "created": 1692901234, "default_parameters": null, "description": "OpenAI text embedding model optimized for performance.", "expiration_date": null, "id": "openai/text-embedding-3-small", "knowledge_cutoff": null, "links": {"details": "/api/v1/models/openai/text-embedding-3-small/endpoints"}, "name": "Text Embedding 3 Small", "per_request_limits": null, "pricing": {"completion": "0", "image": "0", "prompt": "0.00000002", "request": "0"}, "supported_parameters": [], "supported_voices": null, "top_provider": {"context_length": 8192, "is_moderated": false, "max_completion_tokens": null}}]}, "schema": {"$ref": "#/components/schemas/ModelsListResponse"}}}, "description": "Returns a list of embeddings models"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "List all embeddings models", "tags": ["Embeddings"], "x-speakeasy-name-override": "listModels", "x-speakeasy-pagination": {"inputs": [{"in": "parameters", "name": "offset", "type": "offset"}, {"in": "parameters", "name": "limit", "type": "limit"}], "outputs": {"results": "$.data"}, "type": "offsetLimit"}}
```
