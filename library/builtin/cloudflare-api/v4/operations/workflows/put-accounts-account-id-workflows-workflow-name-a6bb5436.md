---
title: Create/modify Workflow
page_id: operation-put-accounts-account-id-workflows-workflow-name-e8d126d8
path: operations/workflows
description: Creates a new workflow or updates an existing workflow definition.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/workflows/{workflow_name}
operation_ids:
    - wor-create-or-modify-workflow
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create/modify Workflow

`PUT /accounts/{account_id}/workflows/{workflow_name}`

Operation ID: `wor-create-or-modify-workflow`

Creates a new workflow or updates an existing workflow definition.

## Definition

```yaml
{"operationId": "wor-create-or-modify-workflow", "summary": "Create/modify Workflow", "description": "Creates a new workflow or updates an existing workflow definition.", "parameters": [{"name": "workflow_name", "in": "path", "required": true, "schema": {"type": "string", "maxLength": 64, "minLength": 1, "pattern": "^[a-zA-Z0-9_][a-zA-Z0-9-_]*$"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"class_name": {"type": "string", "maxLength": 255, "minLength": 1}, "default_retention": {"description": "Default retention applied to instances of this version when they do not set their own retention.", "type": "object", "properties": {"error_retention": {"description": "Specifies the duration in milliseconds or as a string like '5 minutes'.", "anyOf": [{"description": "Specifies the duration in milliseconds.", "type": "number"}, {"type": "string"}], "oneOf": [{"description": "Specifies the duration in milliseconds.", "type": "integer"}, {"example": "5 minutes", "pattern": "^(\\d+)\\s+(second|minute|hour|day|week|month|year)s?$", "type": "string"}]}, "success_retention": {"description": "Specifies the duration in milliseconds or as a string like '5 minutes'.", "anyOf": [{"description": "Specifies the duration in milliseconds.", "type": "number"}, {"type": "string"}], "oneOf": [{"description": "Specifies the duration in milliseconds.", "type": "integer"}, {"example": "5 minutes", "pattern": "^(\\d+)\\s+(second|minute|hour|day|week|month|year)s?$", "type": "string"}]}}}, "limits": {"type": "object", "properties": {"steps": {"type": "integer", "minimum": 1}}}, "schedules": {"type": "array", "items": {"properties": {"cron": {"type": "string", "maxLength": 256, "minLength": 1}}, "required": ["cron"], "type": "object"}, "maxItems": 200}, "script_name": {"type": "string", "maxLength": 255, "minLength": 1}}, "required": ["script_name", "class_name"]}}}}, "responses": {"200": {"description": "Create/modify a Workflow based on a deployed script with an existing `WorkflowEntrypoint` class. Must be done after deploying the corresponding script.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "properties": {"class_name": {"type": "string"}, "created_on": {"type": "string", "format": "date-time"}, "id": {"type": "string", "format": "uuid"}, "is_deleted": {"type": "number"}, "modified_on": {"type": "string", "format": "date-time"}, "name": {"type": "string", "maxLength": 64, "minLength": 1, "pattern": "^[a-zA-Z0-9_][a-zA-Z0-9-_]*$"}, "script_name": {"type": "string"}, "terminator_running": {"type": "number"}, "triggered_on": {"type": "string", "format": "date-time", "nullable": true}, "version_id": {"type": "string", "format": "uuid"}}, "required": ["version_id", "name", "id", "created_on", "modified_on", "script_name", "class_name", "triggered_on", "is_deleted", "terminator_running"]}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "cursor": {"type": "string"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}, "total_pages": {"type": "number"}}, "required": ["per_page", "count", "total_count"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "result", "messages"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Workflows"], "x-api-token-group": ["Workers Scripts Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.workers.write", "com.cloudflare.api.workers.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-ignore": true}
```
