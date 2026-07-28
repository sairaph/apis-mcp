---
title: List Public Finetunes
page_id: operation-get-accounts-account-id-ai-finetunes-public-4b979683
path: operations/workers-ai-finetune
description: Lists publicly available fine-tuned models that can be used with Workers AI.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai/finetunes/public
operation_ids:
    - workers-ai-list-public-finetunes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Public Finetunes

`GET /accounts/{account_id}/ai/finetunes/public`

Operation ID: `workers-ai-list-public-finetunes`

Lists publicly available fine-tuned models that can be used with Workers AI.

## Definition

```yaml
{"operationId": "workers-ai-list-public-finetunes", "summary": "List Public Finetunes", "description": "Lists publicly available fine-tuned models that can be used with Workers AI.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}, {"name": "limit", "in": "query", "description": "Pagination Limit", "schema": {"description": "Pagination Limit", "type": "number", "default": 20}}, {"name": "offset", "in": "query", "description": "Pagination Offset", "schema": {"description": "Pagination Offset", "type": "number", "default": 0}}, {"name": "orderBy", "in": "query", "description": "Order By Column Name", "schema": {"description": "Order By Column Name", "type": "string"}}], "responses": {"200": {"description": "Returns all public finetunes", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "array", "items": {"properties": {"created_at": {"type": "string", "format": "date-time"}, "description": {"type": "string"}, "id": {"type": "string", "format": "uuid", "x-auditable": true}, "model": {"type": "string", "x-auditable": true}, "modified_at": {"type": "string", "format": "date-time"}, "name": {"type": "string", "x-auditable": true}, "public": {"type": "boolean", "x-auditable": true}}, "required": ["id", "model", "name", "public", "created_at", "modified_at"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers AI Finetune"], "x-api-token-group": ["Workers AI Write", "Workers AI Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
