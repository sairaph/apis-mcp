---
title: Get Workflow details
page_id: operation-get-accounts-account-id-workflows-workflow-name-58bd6b96
path: operations/workflows
description: Retrieves configuration and metadata for a specific workflow.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workflows/{workflow_name}
operation_ids:
    - wor-get-workflow-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Workflow details

`GET /accounts/{account_id}/workflows/{workflow_name}`

Operation ID: `wor-get-workflow-details`

Retrieves configuration and metadata for a specific workflow.

## Definition

```yaml
{"operationId": "wor-get-workflow-details", "summary": "Get Workflow details", "description": "Retrieves configuration and metadata for a specific workflow.", "parameters": [{"name": "workflow_name", "in": "path", "required": true, "schema": {"type": "string", "maxLength": 64, "minLength": 1, "pattern": "^[a-zA-Z0-9_][a-zA-Z0-9-_]*$"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Get Workflow details.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "properties": {"class_name": {"type": "string"}, "created_on": {"type": "string", "format": "date-time"}, "id": {"type": "string", "format": "uuid"}, "instances": {"type": "object", "properties": {"complete": {"type": "number"}, "errored": {"type": "number"}, "paused": {"type": "number"}, "queued": {"type": "number"}, "rollingBack": {"type": "number"}, "running": {"type": "number"}, "terminated": {"type": "number"}, "waiting": {"type": "number"}, "waitingForPause": {"type": "number"}}}, "modified_on": {"type": "string", "format": "date-time"}, "name": {"type": "string", "maxLength": 64, "minLength": 1, "pattern": "^[a-zA-Z0-9_][a-zA-Z0-9-_]*$"}, "schedules": {"type": "array", "items": {"properties": {"cron": {"type": "string"}, "next_instance": {"type": "string"}}, "required": ["cron", "next_instance"], "type": "object"}}, "script_name": {"type": "string"}, "triggered_on": {"type": "string", "format": "date-time", "nullable": true}}, "required": ["name", "id", "created_on", "modified_on", "script_name", "class_name", "triggered_on", "instances"]}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "cursor": {"type": "string"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}, "total_pages": {"type": "number"}}, "required": ["per_page", "count", "total_count"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "result", "messages"]}}}}, "400": {"description": "Workflow has no deployed versions.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}, "404": {"description": "Workflow not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Workflows"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.workers.write", "com.cloudflare.api.workers.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workflows", "x-fern-sdk-method-name": "get"}
```
