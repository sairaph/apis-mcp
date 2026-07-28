---
title: Fetch a Account Provider
page_id: operation-get-accounts-account-id-ai-gateway-custom-providers-id-59773f51
path: operations/ai-gateway-account-providers
description: Retrieves details for a specific AI Gateway dataset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-gateway/custom-providers/{id}
operation_ids:
    - aig-config-fetch-account-provider
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch a Account Provider

`GET /accounts/{account_id}/ai-gateway/custom-providers/{id}`

Operation ID: `aig-config-fetch-account-provider`

Retrieves details for a specific AI Gateway dataset.

## Definition

```yaml
{"operationId": "aig-config-fetch-account-provider", "summary": "Fetch a Account Provider", "description": "Retrieves details for a specific AI Gateway dataset.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "3ebbcb006d4d46d7bb6a8c7f14676cb0"}}, {"name": "id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Returns a single object if found", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"base_url": {"type": "string", "format": "uri"}, "beta": {"type": "boolean"}, "created_at": {"type": "string", "format": "date-time"}, "curl_example": {"type": "string"}, "description": {"type": "string"}, "enable": {"type": "boolean"}, "headers": {"type": "string", "maxLength": 8192}, "id": {"type": "string", "format": "uuid"}, "js_example": {"type": "string"}, "link": {"type": "string"}, "logo": {"type": "string"}, "modified_at": {"type": "string", "format": "date-time"}, "name": {"type": "string"}, "position": {"type": "integer"}, "slug": {"type": "string"}}, "required": ["id", "created_at", "modified_at", "name", "slug", "base_url"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "404": {"description": "Not Found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7002}, "message": {"type": "string", "example": "Not Found"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Account Providers"], "x-api-token-group": ["AI Gateway Write", "AI Gateway Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
