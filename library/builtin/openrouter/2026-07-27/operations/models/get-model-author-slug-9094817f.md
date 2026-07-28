---
title: Get a model by its slug
page_id: operation-get-model-author-slug-54352fff
path: operations/models
description: Returns full details for a single model identified by its author and slug (e.g. openai/gpt-4). Supports variant suffixes (e.g. openai/gpt-4:free) and resolves known slug aliases.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /model/{author}/{slug}
operation_ids:
    - getModel
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Get a model by its slug

`GET /model/{author}/{slug}`

Operation ID: `getModel`

Returns full details for a single model identified by its author and slug (e.g. openai/gpt-4). Supports variant suffixes (e.g. openai/gpt-4:free) and resolves known slug aliases.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Returns full details for a single model identified by its author and slug (e.g. openai/gpt-4). Supports variant suffixes (e.g. openai/gpt-4:free) and resolves known slug aliases.", "operationId": "getModel", "parameters": [{"description": "The author/organization of the model", "in": "path", "name": "author", "required": true, "schema": {"description": "The author/organization of the model", "example": "openai", "type": "string"}}, {"description": "The model slug, optionally including a variant suffix (e.g. gpt-4 or gpt-4:free)", "in": "path", "name": "slug", "required": true, "schema": {"description": "The model slug, optionally including a variant suffix (e.g. gpt-4 or gpt-4:free)", "example": "gpt-4", "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": {"architecture": {"input_modalities": ["text"], "instruct_type": "chatml", "modality": "text->text", "output_modalities": ["text"], "tokenizer": "GPT"}, "canonical_slug": "openai/gpt-4", "context_length": 8192, "created": 1692901234, "default_parameters": null, "description": "GPT-4 is a large multimodal model that can solve difficult problems with greater accuracy.", "expiration_date": null, "id": "openai/gpt-4", "knowledge_cutoff": null, "links": {"details": "/api/v1/models/openai/gpt-4/endpoints"}, "name": "GPT-4", "per_request_limits": null, "pricing": {"completion": "0.00006", "image": "0", "prompt": "0.00003", "request": "0"}, "supported_parameters": ["temperature", "top_p", "max_tokens"], "supported_voices": null, "top_provider": {"context_length": 8192, "is_moderated": true, "max_completion_tokens": 4096}}}, "schema": {"$ref": "#/components/schemas/ModelResponse"}}}, "description": "Returns the model details"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Get a model by its slug", "tags": ["Models"], "x-speakeasy-name-override": "get"}
```
