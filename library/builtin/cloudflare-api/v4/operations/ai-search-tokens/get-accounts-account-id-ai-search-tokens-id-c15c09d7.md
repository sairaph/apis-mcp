---
title: Get a token
page_id: operation-get-accounts-account-id-ai-search-tokens-id-d5633576
path: operations/ai-search-tokens
description: Retrieve a stored AI Search credential without exposing its secret.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-search/tokens/{id}
operation_ids:
    - ai-search-fetch-tokens
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a token

`GET /accounts/{account_id}/ai-search/tokens/{id}`

Operation ID: `ai-search-fetch-tokens`

Retrieve a stored AI Search credential without exposing its secret.

## Definition

```yaml
{"operationId": "ai-search-fetch-tokens", "summary": "Get a token", "description": "Retrieve a stored AI Search credential without exposing its secret.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid", "example": "62af0db3-c410-40b2-9ee3-0e93f6dd1de0"}}], "responses": {"200": {"description": "Token details.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"cf_api_id": {"type": "string", "x-auditable": true}, "created_at": {"type": "string", "format": "date-time", "readOnly": true}, "created_by": {"type": "string", "nullable": true, "readOnly": true}, "enabled": {"type": "boolean", "default": true, "x-auditable": true}, "id": {"type": "string", "format": "uuid", "x-auditable": true}, "legacy": {"type": "boolean", "default": true, "readOnly": true}, "modified_at": {"type": "string", "format": "date-time", "readOnly": true}, "modified_by": {"type": "string", "nullable": true, "readOnly": true}, "name": {"type": "string", "x-auditable": true}}, "required": ["id", "name", "cf_api_id", "created_at", "modified_at"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result"]}}}}, "400": {"description": "Input Validation Error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}, "path": {"type": "array", "items": {"type": "string"}}}, "required": ["code", "message", "path"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "404": {"description": "Token not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Tokens"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search.tokens", "x-fern-sdk-method-name": "get", "x-forge-params": {"id": {"description": "Stored credential record ID.", "flagName": "token-id"}}}
```
