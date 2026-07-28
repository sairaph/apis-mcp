---
title: Preview the impact of ZDR on the available endpoints
page_id: operation-get-endpoints-zdr-6e619222
path: operations/endpoints
description: Preview the impact of ZDR on the available endpoints
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /endpoints/zdr
operation_ids:
    - listEndpointsZdr
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Preview the impact of ZDR on the available endpoints

`GET /endpoints/zdr`

Operation ID: `listEndpointsZdr`

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"operationId": "listEndpointsZdr", "responses": {"200": {"content": {"application/json": {"example": {"data": [{"context_length": 8192, "latency_last_30m": {"p50": 0.25, "p75": 0.35, "p90": 0.48, "p99": 0.85}, "max_completion_tokens": 4096, "max_prompt_tokens": 8192, "model_id": "openai/gpt-4", "model_name": "GPT-4", "name": "OpenAI: GPT-4", "pricing": {"completion": "0.00006", "image": "0", "prompt": "0.00003", "request": "0"}, "provider_name": "OpenAI", "quantization": "fp16", "status": 0, "supported_parameters": ["temperature", "top_p", "max_tokens"], "supports_implicit_caching": true, "tag": "openai", "throughput_last_30m": {"p50": 45.2, "p75": 38.5, "p90": 28.3, "p99": 15.1}, "uptime_last_1d": 99.8, "uptime_last_30m": 99.5, "uptime_last_5m": 100}]}, "schema": {"example": {"data": [{"context_length": 8192, "latency_last_30m": {"p50": 0.25, "p75": 0.35, "p90": 0.48, "p99": 0.85}, "max_completion_tokens": 4096, "max_prompt_tokens": 8192, "model_id": "openai/gpt-4", "model_name": "GPT-4", "name": "OpenAI: GPT-4", "pricing": {"completion": "0.00006", "image": "0", "prompt": "0.00003", "request": "0"}, "provider_name": "OpenAI", "quantization": "fp16", "status": 0, "supported_parameters": ["temperature", "top_p", "max_tokens"], "supports_implicit_caching": true, "tag": "openai", "throughput_last_30m": {"p50": 45.2, "p75": 38.5, "p90": 28.3, "p99": 15.1}, "uptime_last_1d": 99.8, "uptime_last_30m": 99.5, "uptime_last_5m": 100}]}, "properties": {"data": {"items": {"$ref": "#/components/schemas/PublicEndpoint"}, "type": "array"}}, "required": ["data"], "type": "object"}}}, "description": "Returns a list of endpoints"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Preview the impact of ZDR on the available endpoints", "tags": ["Endpoints"], "x-speakeasy-name-override": "listZdrEndpoints"}
```
