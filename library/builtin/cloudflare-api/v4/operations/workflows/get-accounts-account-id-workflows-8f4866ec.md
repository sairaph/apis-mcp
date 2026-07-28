---
title: List all Workflows
page_id: operation-get-accounts-account-id-workflows-69a21119
path: operations/workflows
description: Lists all workflows configured for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workflows
operation_ids:
    - wor-list-workflows
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all Workflows

`GET /accounts/{account_id}/workflows`

Operation ID: `wor-list-workflows`

Lists all workflows configured for the account.

## Definition

```yaml
{"operationId": "wor-list-workflows", "summary": "List all Workflows", "description": "Lists all workflows configured for the account.", "parameters": [{"name": "per_page", "in": "query", "schema": {"type": "number", "default": 10, "maximum": 100, "minimum": 1}}, {"name": "page", "in": "query", "schema": {"type": "number", "default": 1, "minimum": 1}}, {"name": "search", "in": "query", "description": "Allows filtering workflows` name.", "schema": {"description": "Allows filtering workflows` name.", "type": "string", "maxLength": 64, "minLength": 1, "pattern": "^[a-zA-Z0-9_][a-zA-Z0-9-_]*$"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "List of all Workflows belonging to a account.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "array", "items": {"properties": {"class_name": {"type": "string"}, "created_on": {"type": "string", "format": "date-time"}, "id": {"type": "string", "format": "uuid"}, "instances": {"type": "object", "properties": {"complete": {"type": "number"}, "errored": {"type": "number"}, "paused": {"type": "number"}, "queued": {"type": "number"}, "rollingBack": {"type": "number"}, "running": {"type": "number"}, "terminated": {"type": "number"}, "waiting": {"type": "number"}, "waitingForPause": {"type": "number"}}}, "modified_on": {"type": "string", "format": "date-time"}, "name": {"type": "string", "maxLength": 64, "minLength": 1, "pattern": "^[a-zA-Z0-9_][a-zA-Z0-9-_]*$"}, "schedules": {"type": "array", "items": {"properties": {"cron": {"type": "string"}, "next_instance": {"type": "string"}}, "required": ["cron", "next_instance"], "type": "object"}}, "script_name": {"type": "string"}, "triggered_on": {"type": "string", "format": "date-time", "nullable": true}}, "required": ["name", "id", "created_on", "modified_on", "script_name", "class_name", "triggered_on", "instances"], "type": "object"}}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "cursor": {"type": "string"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}, "total_pages": {"type": "number"}}, "required": ["per_page", "count", "total_count"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "result", "messages"]}}}}, "400": {"description": "Input Validation Error.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Workflows"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.workers.write", "com.cloudflare.api.workers.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workflows", "x-fern-sdk-method-name": "list"}
```
