---
title: Create a token
page_id: operation-post-accounts-account-id-ai-search-tokens-87116c39
path: operations/ai-search-tokens
description: Create a stored Cloudflare credential for an AI Search instance to access its data source.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/ai-search/tokens
operation_ids:
    - ai-search-create-tokens
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a token

`POST /accounts/{account_id}/ai-search/tokens`

Operation ID: `ai-search-create-tokens`

Create a stored Cloudflare credential for an AI Search instance to access its data source.

## Definition

```yaml
{"operationId": "ai-search-create-tokens", "summary": "Create a token", "description": "Create a stored Cloudflare credential for an AI Search instance to access its data source.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"cf_api_id": {"type": "string", "example": "a1b2c3d4e5f6"}, "cf_api_key": {"type": "string", "example": "abc123", "writeOnly": true, "x-sensitive": true}, "legacy": {"type": "boolean", "default": true}, "name": {"type": "string", "example": "my-token"}}, "required": ["name", "cf_api_id", "cf_api_key"]}}}}, "responses": {"201": {"description": "Token created.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"cf_api_id": {"type": "string", "x-auditable": true}, "created_at": {"type": "string", "format": "date-time", "readOnly": true}, "created_by": {"type": "string", "nullable": true, "readOnly": true}, "enabled": {"type": "boolean", "default": true, "x-auditable": true}, "id": {"type": "string", "format": "uuid", "x-auditable": true}, "legacy": {"type": "boolean", "default": true, "readOnly": true}, "modified_at": {"type": "string", "format": "date-time", "readOnly": true}, "modified_by": {"type": "string", "nullable": true, "readOnly": true}, "name": {"type": "string", "x-auditable": true}}, "required": ["id", "name", "cf_api_id", "created_at", "modified_at"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result"]}}}}, "400": {"description": "Ai search instance invalid token.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Tokens"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search.tokens", "x-fern-sdk-method-name": "create", "x-forge-params": {"cf_api_id": {"description": "Cloudflare API token ID.", "flagName": "api-token-id"}, "cf_api_key": {"description": "Cloudflare API token value.", "flagName": "api-token"}, "legacy": {"description": "Whether the credential uses the legacy API key format."}, "name": {"description": "Human-readable name for the credential.", "positional": true}}}
```
