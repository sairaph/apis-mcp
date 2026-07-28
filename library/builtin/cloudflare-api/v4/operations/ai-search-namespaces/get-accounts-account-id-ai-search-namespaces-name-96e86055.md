---
title: Get a namespace
page_id: operation-get-accounts-account-id-ai-search-namespaces-name-30b6f787
path: operations/ai-search-namespaces
description: Retrieve a namespace and its description.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-search/namespaces/{name}
operation_ids:
    - ai-search-fetch-namespace
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a namespace

`GET /accounts/{account_id}/ai-search/namespaces/{name}`

Operation ID: `ai-search-fetch-namespace`

Retrieve a namespace and its description.

## Definition

```yaml
{"operationId": "ai-search-fetch-namespace", "summary": "Get a namespace", "description": "Retrieve a namespace and its description.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "name", "in": "path", "required": true, "schema": {"type": "string", "example": "production"}}], "responses": {"200": {"description": "Namespace details.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time", "readOnly": true}, "description": {"description": "Optional description for the namespace. Max 256 characters.", "type": "string", "example": "Production environment", "maxLength": 256, "nullable": true}, "name": {"type": "string", "example": "production", "pattern": "^[a-z0-9]([a-z0-9-]{0,26}[a-z0-9])?$"}}, "required": ["name", "created_at"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result"]}}}}, "404": {"description": "Namespace not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Namespaces"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search.namespace", "x-fern-sdk-method-name": "get", "x-forge-params": {"name": {"description": "Namespace to retrieve.", "flagName": "namespace"}}}
```
