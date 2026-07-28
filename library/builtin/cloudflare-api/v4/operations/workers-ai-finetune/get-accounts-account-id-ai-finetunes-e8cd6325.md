---
title: List Finetunes
page_id: operation-get-accounts-account-id-ai-finetunes-efa9319b
path: operations/workers-ai-finetune
description: Lists all fine-tuning jobs created by the account, including status and metrics.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai/finetunes
operation_ids:
    - workers-ai-list-finetunes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Finetunes

`GET /accounts/{account_id}/ai/finetunes`

Operation ID: `workers-ai-list-finetunes`

Lists all fine-tuning jobs created by the account, including status and metrics.

## Definition

```yaml
{"operationId": "workers-ai-list-finetunes", "summary": "List Finetunes", "description": "Lists all fine-tuning jobs created by the account, including status and metrics.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}], "responses": {"200": {"description": "Returns all finetunes", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "description": {"type": "string"}, "id": {"type": "string", "format": "uuid", "x-auditable": true}, "model": {"type": "string", "x-auditable": true}, "modified_at": {"type": "string", "format": "date-time"}, "name": {"type": "string", "x-auditable": true}}, "required": ["id", "model", "name", "created_at", "modified_at"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers AI Finetune"], "x-api-token-group": ["Workers AI Write", "Workers AI Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
