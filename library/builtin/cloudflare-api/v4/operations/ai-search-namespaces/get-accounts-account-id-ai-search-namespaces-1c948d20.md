---
title: List namespaces
page_id: operation-get-accounts-account-id-ai-search-namespaces-bd05d0db
path: operations/ai-search-namespaces
description: List namespaces in the account, including their descriptions and creation times.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-search/namespaces
operation_ids:
    - ai-search-list-namespaces
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List namespaces

`GET /accounts/{account_id}/ai-search/namespaces`

Operation ID: `ai-search-list-namespaces`

List namespaces in the account, including their descriptions and creation times.

## Definition

```yaml
{"operationId": "ai-search-list-namespaces", "summary": "List namespaces", "description": "List namespaces in the account, including their descriptions and creation times.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "page", "in": "query", "description": "Page number (1-indexed).", "schema": {"description": "Page number (1-indexed).", "type": "integer", "example": 1, "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Number of results per page.", "schema": {"description": "Number of results per page.", "type": "integer", "example": 20, "default": 20, "maximum": 100, "minimum": 1}}, {"name": "search", "in": "query", "description": "Filter namespaces whose name or description contains this string (case-insensitive).", "schema": {"description": "Filter namespaces whose name or description contains this string (case-insensitive).", "type": "string", "example": "prod", "maxLength": 256}}], "responses": {"200": {"description": "List of namespaces.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "array", "items": {"properties": {"created_at": {"type": "string", "format": "date-time", "readOnly": true}, "description": {"description": "Optional description for the namespace. Max 256 characters.", "type": "string", "example": "Production environment", "maxLength": 256, "nullable": true}, "name": {"type": "string", "example": "production", "pattern": "^[a-z0-9]([a-z0-9-]{0,26}[a-z0-9])?$"}}, "required": ["name", "created_at"], "type": "object"}}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}}, "required": ["page", "per_page", "count", "total_count"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result", "result_info"]}}}}, "400": {"description": "Input Validation Error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}, "path": {"type": "array", "items": {"type": "string"}}}, "required": ["code", "message", "path"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Namespaces"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search.namespace", "x-fern-sdk-method-name": "list"}
```
