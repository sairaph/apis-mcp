---
title: Update a token
page_id: operation-put-accounts-account-id-ai-search-tokens-id-112d97fa
path: operations/ai-search-tokens
description: Replace a stored AI Search credential and invalidate cached credentials for instances that use it.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/ai-search/tokens/{id}
operation_ids:
    - ai-search-update-tokens
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a token

`PUT /accounts/{account_id}/ai-search/tokens/{id}`

Operation ID: `ai-search-update-tokens`

Replace a stored AI Search credential and invalidate cached credentials for instances that use it.

## Definition

```yaml
{"operationId": "ai-search-update-tokens", "summary": "Update a token", "description": "Replace a stored AI Search credential and invalidate cached credentials for instances that use it.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid", "example": "62af0db3-c410-40b2-9ee3-0e93f6dd1de0"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"cf_api_id": {"type": "string", "example": "a1b2c3d4e5f6"}, "cf_api_key": {"type": "string", "example": "abc123", "writeOnly": true, "x-sensitive": true}, "legacy": {"type": "boolean", "default": true}, "name": {"type": "string", "example": "my-token"}}, "required": ["name", "cf_api_id", "cf_api_key"]}}}}, "responses": {"200": {"description": "Returns the updated token.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"cf_api_id": {"type": "string", "x-auditable": true}, "created_at": {"type": "string", "format": "date-time", "readOnly": true}, "created_by": {"type": "string", "nullable": true, "readOnly": true}, "enabled": {"type": "boolean", "default": true, "x-auditable": true}, "id": {"type": "string", "format": "uuid", "x-auditable": true}, "legacy": {"type": "boolean", "default": true, "readOnly": true}, "modified_at": {"type": "string", "format": "date-time", "readOnly": true}, "modified_by": {"type": "string", "nullable": true, "readOnly": true}, "name": {"type": "string", "x-auditable": true}}, "required": ["id", "name", "cf_api_id", "created_at", "modified_at"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result"]}}}}, "400": {"description": "Ai search instance invalid token.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "404": {"description": "Token not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Tokens"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search.tokens", "x-fern-sdk-method-name": "update", "x-forge-params": {"cf_api_id": {"description": "Cloudflare API token ID.", "flagName": "api-token-id"}, "cf_api_key": {"description": "Cloudflare API token value.", "flagName": "api-token"}, "id": {"description": "Stored credential record ID.", "flagName": "token-id"}, "legacy": {"description": "Whether the credential uses the legacy API key format."}, "name": {"description": "Human-readable name for the credential."}}}
```
