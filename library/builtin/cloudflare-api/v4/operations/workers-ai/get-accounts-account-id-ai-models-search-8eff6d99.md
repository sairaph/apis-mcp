---
title: Model Search
page_id: operation-get-accounts-account-id-ai-models-search-06b4cbdf
path: operations/workers-ai
description: Searches Workers AI models by name or description.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai/models/search
operation_ids:
    - workers-ai-search-model
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Model Search

`GET /accounts/{account_id}/ai/models/search`

Operation ID: `workers-ai-search-model`

Searches Workers AI models by name or description.

## Definition

```yaml
{"operationId": "workers-ai-search-model", "summary": "Model Search", "description": "Searches Workers AI models by name or description.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}, {"name": "per_page", "in": "query", "schema": {"type": "integer", "default": 100}}, {"name": "page", "in": "query", "schema": {"type": "integer", "default": 1}}, {"name": "task", "in": "query", "description": "Filter by Task Name", "schema": {"description": "Filter by Task Name", "type": "string", "example": "Text Generation", "default": ""}}, {"name": "author", "in": "query", "description": "Filter by Author", "schema": {"description": "Filter by Author", "type": "string", "default": ""}}, {"name": "source", "in": "query", "description": "Filter by Source Id", "schema": {"description": "Filter by Source Id", "type": "number"}}, {"name": "hide_experimental", "in": "query", "description": "Filter to hide experimental models", "schema": {"description": "Filter to hide experimental models", "type": "boolean", "default": false}}, {"name": "search", "in": "query", "description": "Search", "schema": {"description": "Search", "type": "string", "default": ""}}, {"name": "include_deprecated", "in": "query", "description": "If true, include models whose planned_deprecation_date is in the past — but only within a three-month grace window after that date. Models whose planned_deprecation_date is more than three months in the past remain hidden regardless of this flag. Future planned-deprecation dates are always included regardless of this flag. Defaults to false, preserving the existing behavior of hiding all past-dated deprecations.", "schema": {"description": "If true, include models whose planned_deprecation_date is in the past — but only within a three-month grace window after that date. Models whose planned_deprecation_date is more than three months in the past remain hidden regardless of this flag. Future planned-deprecation dates are always included regardless of this flag. Defaults to false, preserving the existing behavior of hiding all past-dated deprecations.", "type": "boolean", "default": false}}, {"name": "format", "in": "query", "description": "If set, return models in the requested marketplace format instead of the default response.", "schema": {"description": "If set, return models in the requested marketplace format instead of the default response.", "type": "string", "enum": ["openrouter"]}}], "responses": {"200": {"description": "Returns a list of models. Default shape is the standard envelope; when `format` is supplied the marketplace-specific shape is returned instead.", "content": {"application/json": {"schema": {"anyOf": [{"properties": {"errors": {"type": "array", "items": {"type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "array", "items": {"type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "result", "errors", "messages"], "type": "object"}, {"description": "Marketplace-format response. See https://openrouter.ai/docs/guides/get-started/for-providers", "properties": {"data": {"type": "array", "items": {"type": "object"}}}, "required": ["data"], "type": "object"}]}}}}, "404": {"description": "Object not found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers AI"], "x-api-token-group": ["Workers AI Write", "Workers AI Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
