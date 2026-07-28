---
title: List tokens
page_id: operation-get-accounts-account-id-ai-search-tokens-3267142a
path: operations/ai-search-tokens
description: List stored AI Search credentials in the account without exposing their secrets.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-search/tokens
operation_ids:
    - ai-search-list-tokens
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List tokens

`GET /accounts/{account_id}/ai-search/tokens`

Operation ID: `ai-search-list-tokens`

List stored AI Search credentials in the account without exposing their secrets.

## Definition

```yaml
{"operationId": "ai-search-list-tokens", "summary": "List tokens", "description": "List stored AI Search credentials in the account without exposing their secrets.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "page", "in": "query", "description": "Page number (1-indexed).", "schema": {"description": "Page number (1-indexed).", "type": "integer", "example": 1, "default": 1, "maximum": 100, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Number of results per page.", "schema": {"description": "Number of results per page.", "type": "integer", "example": 20, "default": 20, "maximum": 100, "minimum": 1}}, {"name": "search", "in": "query", "description": "Filter tokens whose name contains this string (case-insensitive).", "schema": {"description": "Filter tokens whose name contains this string (case-insensitive).", "type": "string", "example": "my-token", "maxLength": 256}}], "responses": {"200": {"description": "List of tokens.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "array", "items": {"properties": {"cf_api_id": {"type": "string", "x-auditable": true}, "created_at": {"type": "string", "format": "date-time", "readOnly": true}, "created_by": {"type": "string", "nullable": true, "readOnly": true}, "enabled": {"type": "boolean", "default": true, "x-auditable": true}, "id": {"type": "string", "format": "uuid", "x-auditable": true}, "legacy": {"type": "boolean", "default": true, "readOnly": true}, "modified_at": {"type": "string", "format": "date-time", "readOnly": true}, "modified_by": {"type": "string", "nullable": true, "readOnly": true}, "name": {"type": "string", "x-auditable": true}}, "required": ["id", "name", "cf_api_id", "created_at", "modified_at"], "type": "object"}}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}}, "required": ["page", "per_page", "count", "total_count"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result", "result_info"]}}}}, "400": {"description": "Input Validation Error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}, "path": {"type": "array", "items": {"type": "string"}}}, "required": ["code", "message", "path"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Tokens"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search.tokens", "x-fern-sdk-method-name": "list"}
```
