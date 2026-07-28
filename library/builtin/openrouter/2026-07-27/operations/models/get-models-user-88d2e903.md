---
title: List models filtered by user provider preferences, privacy settings, and guardrails
page_id: operation-get-models-user-33adc33e
path: operations/models
description: List models filtered by user provider preferences, [privacy settings](https://openrouter.ai/docs/guides/privacy/provider-logging), and [guardrails](https://openrouter.ai/docs/guides/features/guardrails). If requesting through `eu.openrouter.ai/api/v1/...` the results will be filtered to models that satisfy [EU in-region routing](https://openrouter.ai/docs/guides/privacy/provider-logging#enterprise-eu-in-region-routing).
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /models/user
operation_ids:
    - listModelsUser
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List models filtered by user provider preferences, privacy settings, and guardrails

`GET /models/user`

Operation ID: `listModelsUser`

List models filtered by user provider preferences, [privacy settings](https://openrouter.ai/docs/guides/privacy/provider-logging), and [guardrails](https://openrouter.ai/docs/guides/features/guardrails). If requesting through `eu.openrouter.ai/api/v1/...` the results will be filtered to models that satisfy [EU in-region routing](https://openrouter.ai/docs/guides/privacy/provider-logging#enterprise-eu-in-region-routing).

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "List models filtered by user provider preferences, [privacy settings](https://openrouter.ai/docs/guides/privacy/provider-logging), and [guardrails](https://openrouter.ai/docs/guides/features/guardrails). If requesting through `eu.openrouter.ai/api/v1/...` the results will be filtered to models that satisfy [EU in-region routing](https://openrouter.ai/docs/guides/privacy/provider-logging#enterprise-eu-in-region-routing).", "operationId": "listModelsUser", "parameters": [{"description": "Number of records to skip for pagination. When both offset and limit are omitted, the full list is returned", "in": "query", "name": "offset", "required": false, "schema": {"default": 0, "description": "Number of records to skip for pagination. When both offset and limit are omitted, the full list is returned", "example": 0, "minimum": 0, "type": ["integer", "null"]}}, {"description": "Maximum number of records to return (max 1000). When both offset and limit are omitted, the full list is returned", "in": "query", "name": "limit", "required": false, "schema": {"default": 500, "description": "Maximum number of records to return (max 1000). When both offset and limit are omitted, the full list is returned", "example": 500, "maximum": 1000, "minimum": 1, "type": "integer"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": [{"architecture": {"input_modalities": ["text"], "instruct_type": "chatml", "modality": "text->text", "output_modalities": ["text"], "tokenizer": "GPT"}, "canonical_slug": "openai/gpt-4", "context_length": 8192, "created": 1692901234, "default_parameters": null, "description": "GPT-4 is a large multimodal model that can solve difficult problems with greater accuracy.", "expiration_date": null, "id": "openai/gpt-4", "knowledge_cutoff": null, "links": {"details": "/api/v1/models/openai/gpt-4/endpoints"}, "name": "GPT-4", "per_request_limits": null, "pricing": {"completion": "0.00006", "image": "0", "prompt": "0.00003", "request": "0"}, "supported_parameters": ["temperature", "top_p", "max_tokens"], "supported_voices": null, "top_provider": {"context_length": 8192, "is_moderated": true, "max_completion_tokens": 4096}}]}, "schema": {"$ref": "#/components/schemas/ModelsListResponse"}}}, "description": "Returns a list of models filtered by user provider preferences"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "security": [{"bearer": []}], "summary": "List models filtered by user provider preferences, privacy settings, and guardrails", "tags": ["Models"], "x-speakeasy-name-override": "listForUser", "x-speakeasy-pagination": {"inputs": [{"in": "parameters", "name": "offset", "type": "offset"}, {"in": "parameters", "name": "limit", "type": "limit"}], "outputs": {"results": "$.data"}, "type": "offsetLimit"}}
```
