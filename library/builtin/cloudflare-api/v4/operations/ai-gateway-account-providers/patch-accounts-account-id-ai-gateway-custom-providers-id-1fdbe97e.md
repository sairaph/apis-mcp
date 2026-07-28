---
title: Update a Account Provider
page_id: operation-patch-accounts-account-id-ai-gateway-custom-providers-id-74ad8660
path: operations/ai-gateway-account-providers
description: Updates an existing AI Gateway dataset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/ai-gateway/custom-providers/{id}
operation_ids:
    - aig-config-update-account-provider
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a Account Provider

`PATCH /accounts/{account_id}/ai-gateway/custom-providers/{id}`

Operation ID: `aig-config-update-account-provider`

Updates an existing AI Gateway dataset.

## Definition

```yaml
{"operationId": "aig-config-update-account-provider", "summary": "Update a Account Provider", "description": "Updates an existing AI Gateway dataset.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "3ebbcb006d4d46d7bb6a8c7f14676cb0"}}, {"name": "id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"base_url": {"type": "string", "format": "uri"}, "beta": {"type": "boolean"}, "curl_example": {"type": "string"}, "description": {"type": "string"}, "enable": {"type": "boolean"}, "headers": {"type": "string", "maxLength": 8192}, "js_example": {"type": "string"}, "link": {"type": "string"}, "logo": {"type": "string"}, "name": {"type": "string"}, "position": {"type": "integer"}, "slug": {"type": "string"}}}}}}, "responses": {"200": {"description": "Returns the updated Object", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"base_url": {"type": "string", "format": "uri"}, "beta": {"type": "boolean"}, "created_at": {"type": "string", "format": "date-time"}, "curl_example": {"type": "string"}, "description": {"type": "string"}, "enable": {"type": "boolean"}, "headers": {"type": "string", "maxLength": 8192}, "id": {"type": "string", "format": "uuid"}, "js_example": {"type": "string"}, "link": {"type": "string"}, "logo": {"type": "string"}, "modified_at": {"type": "string", "format": "date-time"}, "name": {"type": "string"}, "position": {"type": "integer"}, "slug": {"type": "string"}}, "required": ["id", "created_at", "modified_at", "name", "slug", "base_url"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Input Validation Error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7001}, "message": {"type": "string", "example": "Input Validation Error"}, "path": {"type": "array", "items": {"example": "body", "type": "string"}}}, "required": ["code", "message", "path"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}, "404": {"description": "Not Found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7002}, "message": {"type": "string", "example": "Not Found"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Account Providers"], "x-api-token-group": ["AI Gateway Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
