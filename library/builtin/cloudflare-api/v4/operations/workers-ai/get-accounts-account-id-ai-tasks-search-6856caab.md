---
title: Task Search
page_id: operation-get-accounts-account-id-ai-tasks-search-7f94876b
path: operations/workers-ai
description: Searches Workers AI models by task type (e.g., text-generation, embeddings).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai/tasks/search
operation_ids:
    - workers-ai-search-task
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Task Search

`GET /accounts/{account_id}/ai/tasks/search`

Operation ID: `workers-ai-search-task`

Searches Workers AI models by task type (e.g., text-generation, embeddings).

## Definition

```yaml
{"operationId": "workers-ai-search-task", "summary": "Task Search", "description": "Searches Workers AI models by task type (e.g., text-generation, embeddings).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}], "responses": {"200": {"description": "Returns a list of tasks", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "array", "items": {"type": "object"}}, "success": {"type": "boolean", "x-auditable": true}}, "required": ["success", "result", "errors", "messages"]}}}}, "404": {"description": "Object not found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers AI"], "x-api-token-group": ["Workers AI Write", "Workers AI Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
