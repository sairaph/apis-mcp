---
title: List Benchmarks
page_id: operation-get-benchmarks-5dc229fb
path: operations/benchmarks
description: Unified benchmark endpoint that aggregates scores from multiple benchmark sources (Artificial Analysis, Design Arena). Filter by source to reproduce the exact shapes from the legacy per-source endpoints, or use task_type to find models suited for specific workloads. Authenticate with any valid OpenRouter API key. Rate-limited to 30 requests/minute per key and 500 requests/day per account.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /benchmarks
operation_ids:
    - getBenchmarks
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List Benchmarks

`GET /benchmarks`

Operation ID: `getBenchmarks`

Unified benchmark endpoint that aggregates scores from multiple benchmark sources (Artificial Analysis, Design Arena). Filter by source to reproduce the exact shapes from the legacy per-source endpoints, or use task_type to find models suited for specific workloads. Authenticate with any valid OpenRouter API key. Rate-limited to 30 requests/minute per key and 500 requests/day per account.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Unified benchmark endpoint that aggregates scores from multiple benchmark sources (Artificial Analysis, Design Arena). Filter by source to reproduce the exact shapes from the legacy per-source endpoints, or use task_type to find models suited for specific workloads. Authenticate with any valid OpenRouter API key. Rate-limited to 30 requests/minute per key and 500 requests/day per account.", "operationId": "getBenchmarks", "parameters": [{"description": "Benchmark source to query. Determines the shape of the returned items. When omitted, returns results from all sources.", "in": "query", "name": "source", "required": false, "schema": {"description": "Benchmark source to query. Determines the shape of the returned items. When omitted, returns results from all sources.", "enum": ["artificial-analysis", "design-arena"], "example": "artificial-analysis", "type": "string", "x-speakeasy-unknown-values": "allow"}}, {"description": "Filter results by task type. For Artificial Analysis, maps to the corresponding index. For Design Arena, maps to the matching category.", "in": "query", "name": "task_type", "required": false, "schema": {"description": "Filter results by task type. For Artificial Analysis, maps to the corresponding index. For Design Arena, maps to the matching category.", "enum": ["coding", "intelligence", "agentic"], "example": "coding", "type": "string", "x-speakeasy-unknown-values": "allow"}}, {"description": "Design Arena only: arena to query. Defaults to `models` when source is `design-arena`.", "in": "query", "name": "arena", "required": false, "schema": {"description": "Design Arena only: arena to query. Defaults to `models` when source is `design-arena`.", "enum": ["models", "builders", "agents"], "example": "models", "type": "string", "x-speakeasy-unknown-values": "allow"}}, {"description": "Design Arena only: category within the arena (e.g. `codecategories`, `uicomponent`, `gamedev`, `3d`, `dataviz`, `image`, `video`, `svg`). When omitted, returns all categories.", "in": "query", "name": "category", "required": false, "schema": {"description": "Design Arena only: category within the arena (e.g. `codecategories`, `uicomponent`, `gamedev`, `3d`, `dataviz`, `image`, `video`, `svg`). When omitted, returns all categories.", "example": "codecategories", "type": "string"}}, {"description": "Maximum number of items to return. When omitted, all matching results are returned.", "in": "query", "name": "max_results", "required": false, "schema": {"description": "Maximum number of items to return. When omitted, all matching results are returned.", "example": 50, "minimum": 1, "type": "integer"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": [{"agentic_index": 58.3, "coding_index": 65.8, "display_name": "GPT-4o", "intelligence_index": 71.2, "model_permaslug": "openai/gpt-4o", "pricing": {"completion": "0.00001", "prompt": "0.0000025"}, "source": "artificial-analysis"}], "meta": {"as_of": "2026-06-03T12:00:00Z", "citation": null, "model_count": 1, "source": null, "source_url": null, "task_type": null, "version": "v1"}}, "schema": {"$ref": "#/components/schemas/UnifiedBenchmarksResponse"}}}, "description": "Benchmark results filtered by the specified source and optional task type."}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "429": {"content": {"application/json": {"example": {"error": {"code": 429, "message": "Rate limit exceeded"}}, "schema": {"$ref": "#/components/schemas/TooManyRequestsResponse"}}}, "description": "Too Many Requests - Rate limit exceeded"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "List Benchmarks", "tags": ["Benchmarks"]}
```
