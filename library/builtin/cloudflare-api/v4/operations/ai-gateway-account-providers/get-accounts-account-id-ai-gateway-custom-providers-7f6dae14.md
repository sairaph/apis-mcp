---
title: List Account Providers
page_id: operation-get-accounts-account-id-ai-gateway-custom-providers-f47d0ba8
path: operations/ai-gateway-account-providers
description: Lists all AI Gateway evaluator types configured for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-gateway/custom-providers
operation_ids:
    - aig-config-list-account-provider
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Account Providers

`GET /accounts/{account_id}/ai-gateway/custom-providers`

Operation ID: `aig-config-list-account-provider`

Lists all AI Gateway evaluator types configured for the account.

## Definition

```yaml
{"operationId": "aig-config-list-account-provider", "summary": "List Account Providers", "description": "Lists all AI Gateway evaluator types configured for the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "3ebbcb006d4d46d7bb6a8c7f14676cb0"}}, {"name": "page", "in": "query", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"type": "integer", "default": 20, "maximum": 100, "minimum": 1}}, {"name": "beta", "in": "query", "schema": {"type": "boolean"}}, {"name": "enable", "in": "query", "schema": {"type": "boolean"}}, {"name": "search", "in": "query", "schema": {"description": "Search by id, name, slug", "type": "string"}}], "responses": {"200": {"description": "List objects", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "array", "items": {"properties": {"base_url": {"type": "string", "format": "uri"}, "beta": {"type": "boolean"}, "created_at": {"type": "string", "format": "date-time"}, "curl_example": {"type": "string"}, "description": {"type": "string"}, "enable": {"type": "boolean"}, "headers": {"type": "string", "maxLength": 8192}, "id": {"type": "string", "format": "uuid"}, "js_example": {"type": "string"}, "link": {"type": "string"}, "logo": {"type": "string"}, "modified_at": {"type": "string", "format": "date-time"}, "name": {"type": "string"}, "position": {"type": "integer"}, "slug": {"type": "string"}}, "required": ["id", "created_at", "modified_at", "name", "slug", "base_url"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Account Providers"], "x-api-token-group": ["AI Gateway Write", "AI Gateway Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
