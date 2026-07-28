---
title: Create a new Finetune
page_id: operation-post-accounts-account-id-ai-finetunes-a74c9202
path: operations/workers-ai-finetune
description: Creates a new fine-tuning job for a Workers AI model using custom training data.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/ai/finetunes
operation_ids:
    - workers-ai-create-finetune
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new Finetune

`POST /accounts/{account_id}/ai/finetunes`

Operation ID: `workers-ai-create-finetune`

Creates a new fine-tuning job for a Workers AI model using custom training data.

## Definition

```yaml
{"operationId": "workers-ai-create-finetune", "summary": "Create a new Finetune", "description": "Creates a new fine-tuning job for a Workers AI model using custom training data.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"type": "string"}, "model": {"type": "string", "x-auditable": true}, "name": {"type": "string", "x-auditable": true}, "public": {"type": "boolean", "default": false, "x-auditable": true}}, "required": ["model", "name"]}}}}, "responses": {"200": {"description": "Returns the created finetune", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "description": {"type": "string"}, "id": {"type": "string", "format": "uuid", "x-auditable": true}, "model": {"type": "string", "x-auditable": true}, "modified_at": {"type": "string", "format": "date-time"}, "name": {"type": "string", "x-auditable": true}, "public": {"type": "boolean", "x-auditable": true}}, "required": ["id", "created_at", "modified_at", "public", "name", "model"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Finetune creation failed", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}}, "success": {"type": "boolean"}}, "required": ["errors", "success"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers AI Finetune"], "x-api-token-group": ["Workers AI Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
