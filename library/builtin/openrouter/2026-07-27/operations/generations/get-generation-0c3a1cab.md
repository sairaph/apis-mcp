---
title: Get request & usage metadata for a generation
page_id: operation-get-generation-e25d5906
path: operations/generations
description: Get request & usage metadata for a generation
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /generation
operation_ids:
    - getGeneration
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Get request & usage metadata for a generation

`GET /generation`

Operation ID: `getGeneration`

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"operationId": "getGeneration", "parameters": [{"description": "The generation ID", "in": "query", "name": "id", "required": true, "schema": {"description": "The generation ID", "example": "gen-1234567890", "minLength": 1, "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": {"api_type": "completions", "app_id": 12345, "cache_discount": null, "cancelled": false, "created_at": "2024-07-15T23:33:19.433273+00:00", "data_region": "global", "external_user": "user-123", "finish_reason": "stop", "generation_time": 1200, "http_referer": "https://openrouter.ai/", "id": "gen-3bhGkxlo4XFrqiabUM7NDtwDzWwG", "is_byok": false, "latency": 1250, "model": "sao10k/l3-stheno-8b", "moderation_latency": 50, "native_finish_reason": "stop", "native_tokens_cached": 3, "native_tokens_completion": 25, "native_tokens_completion_images": 0, "native_tokens_prompt": 10, "native_tokens_reasoning": 5, "num_fetches": 0, "num_input_audio_prompt": 0, "num_media_completion": 0, "num_media_prompt": 1, "num_search_results": 5, "origin": "https://openrouter.ai/", "preset_id": "a9e8d400-592a-494f-908c-375efa66cafd", "provider_name": "Infermatic", "provider_responses": null, "request_id": "req-1727282430-aBcDeFgHiJkLmNoPqRsT", "router": "openrouter/auto", "service_tier": "priority", "session_id": null, "streamed": true, "tokens_completion": 25, "tokens_prompt": 10, "total_cost": 0.0015, "upstream_id": "chatcmpl-791bcf62-080e-4568-87d0-94c72e3b4946", "upstream_inference_cost": 0.0012, "usage": 0.0015, "user_agent": "Mozilla/5.0", "web_search_engine": "exa"}}, "schema": {"$ref": "#/components/schemas/GenerationResponse"}}}, "description": "Returns the request metadata for this generation"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "402": {"content": {"application/json": {"example": {"error": {"code": 402, "message": "Insufficient credits. Add more using https://openrouter.ai/credits"}}, "schema": {"$ref": "#/components/schemas/PaymentRequiredResponse"}}}, "description": "Payment Required - Insufficient credits or quota to complete request"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "429": {"content": {"application/json": {"example": {"error": {"code": 429, "message": "Rate limit exceeded"}}, "schema": {"$ref": "#/components/schemas/TooManyRequestsResponse"}}}, "description": "Too Many Requests - Rate limit exceeded"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}, "502": {"content": {"application/json": {"example": {"error": {"code": 502, "message": "Provider returned error"}}, "schema": {"$ref": "#/components/schemas/BadGatewayResponse"}}}, "description": "Bad Gateway - Provider/upstream API failure"}, "524": {"content": {"application/json": {"example": {"error": {"code": 524, "message": "Request timed out. Please try again later."}}, "schema": {"$ref": "#/components/schemas/EdgeNetworkTimeoutResponse"}}}, "description": "Infrastructure Timeout - Provider request timed out at edge network"}, "529": {"content": {"application/json": {"example": {"error": {"code": 529, "message": "Provider returned error"}}, "schema": {"$ref": "#/components/schemas/ProviderOverloadedResponse"}}}, "description": "Provider Overloaded - Provider is temporarily overloaded"}}, "summary": "Get request & usage metadata for a generation", "tags": ["Generations"]}
```
