---
title: List all endpoints for a model
page_id: operation-get-models-author-slug-endpoints-fbccd557
path: operations/endpoints
description: List all endpoints for a model
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /models/{author}/{slug}/endpoints
operation_ids:
    - listEndpoints
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List all endpoints for a model

`GET /models/{author}/{slug}/endpoints`

Operation ID: `listEndpoints`

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"operationId": "listEndpoints", "parameters": [{"description": "The author/organization of the model", "in": "path", "name": "author", "required": true, "schema": {"description": "The author/organization of the model", "example": "openai", "type": "string"}}, {"description": "The model slug", "in": "path", "name": "slug", "required": true, "schema": {"description": "The model slug", "example": "gpt-4", "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": {"architecture": {"input_modalities": ["text"], "instruct_type": "chatml", "modality": "text->text", "output_modalities": ["text"], "tokenizer": "GPT"}, "created": 1692901234, "description": "GPT-4 is a large multimodal model.", "endpoints": [], "id": "openai/gpt-4", "name": "GPT-4"}}, "schema": {"example": {"data": {"architecture": {"input_modalities": ["text"], "instruct_type": "chatml", "modality": "text->text", "output_modalities": ["text"], "tokenizer": "GPT"}, "created": 1692901234, "description": "GPT-4 is a large multimodal model that can solve difficult problems with greater accuracy.", "endpoints": [{"context_length": 8192, "latency_last_30m": {"p50": 0.25, "p75": 0.35, "p90": 0.48, "p99": 0.85}, "max_completion_tokens": 4096, "max_prompt_tokens": 8192, "model_id": "openai/gpt-4", "model_name": "GPT-4", "name": "OpenAI: GPT-4", "pricing": {"completion": "0.00006", "image": "0", "prompt": "0.00003", "request": "0"}, "provider_name": "OpenAI", "quantization": "fp16", "status": 0, "supported_parameters": ["temperature", "top_p", "max_tokens"], "supports_implicit_caching": true, "tag": "openai", "throughput_last_30m": {"p50": 45.2, "p75": 38.5, "p90": 28.3, "p99": 15.1}, "uptime_last_1d": 99.8, "uptime_last_30m": 99.5, "uptime_last_5m": 100}], "id": "openai/gpt-4", "name": "GPT-4"}}, "properties": {"data": {"$ref": "#/components/schemas/ListEndpointsResponse"}}, "required": ["data"], "type": "object"}}}, "description": "Returns a list of endpoints"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "List all endpoints for a model", "tags": ["Endpoints"], "x-speakeasy-name-override": "list"}
```
