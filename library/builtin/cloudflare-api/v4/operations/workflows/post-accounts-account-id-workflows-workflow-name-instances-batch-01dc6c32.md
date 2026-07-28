---
title: Batch create new Workflow instances
page_id: operation-post-accounts-account-id-workflows-workflow-name-instances-batch-d2d2eead
path: operations/workflows
description: Creates multiple workflow instances in a single batch operation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/workflows/{workflow_name}/instances/batch
operation_ids:
    - wor-batch-create-workflow-instance
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Batch create new Workflow instances

`POST /accounts/{account_id}/workflows/{workflow_name}/instances/batch`

Operation ID: `wor-batch-create-workflow-instance`

Creates multiple workflow instances in a single batch operation.

## Definition

```yaml
{"operationId": "wor-batch-create-workflow-instance", "summary": "Batch create new Workflow instances", "description": "Creates multiple workflow instances in a single batch operation.", "parameters": [{"name": "workflow_name", "in": "path", "required": true, "schema": {"type": "string", "maxLength": 64, "minLength": 1, "pattern": "^[a-zA-Z0-9_][a-zA-Z0-9-_]*$"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"instance_id": {"type": "string", "maximum": 100, "minimum": 1, "pattern": "^[a-zA-Z0-9_][a-zA-Z0-9-_]*$"}, "instance_retention": {"type": "object", "properties": {"error_retention": {"description": "Specifies the duration in milliseconds or as a string like '5 minutes'.", "anyOf": [{"description": "Specifies the duration in milliseconds.", "type": "number"}, {"type": "string"}], "oneOf": [{"description": "Specifies the duration in milliseconds.", "type": "integer"}, {"example": "5 minutes", "pattern": "^(\\d+)\\s+(second|minute|hour|day|week|month|year)s?$", "type": "string"}]}, "success_retention": {"description": "Specifies the duration in milliseconds or as a string like '5 minutes'.", "anyOf": [{"description": "Specifies the duration in milliseconds.", "type": "number"}, {"type": "string"}], "oneOf": [{"description": "Specifies the duration in milliseconds.", "type": "integer"}, {"example": "5 minutes", "pattern": "^(\\d+)\\s+(second|minute|hour|day|week|month|year)s?$", "type": "string"}]}}}, "params": {"description": "JSON-encoded event payload passed into the new instance.", "type": "string"}}, "type": "object"}, "maxItems": 100, "minItems": 1}}}}, "responses": {"200": {"description": "Batch create workflow instances. Body is a JSON list that contain all payloads and ids that are passed into the new instance as the event payload.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "array", "items": {"properties": {"id": {"type": "string", "maxLength": 100, "minLength": 1, "pattern": "^[a-zA-Z0-9_][a-zA-Z0-9-_]*$"}, "status": {"type": "string", "enum": ["queued", "running", "paused", "errored", "terminated", "complete", "waitingForPause", "waiting", "rollingBack"]}, "trigger_source": {"type": "string", "enum": ["unknown", "api", "binding", "event", "cron"]}, "version_id": {"type": "string", "format": "uuid"}, "workflow_id": {"type": "string", "format": "uuid"}}, "required": ["id", "workflow_id", "version_id", "status"], "type": "object"}}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "cursor": {"type": "string"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}, "total_pages": {"type": "number"}}, "required": ["per_page", "count", "total_count"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "result", "messages"]}}}}, "400": {"description": "Provided Workflow ID is not valid.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}, "404": {"description": "Workflow Name not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Workflows"], "x-api-token-group": ["Workers Scripts Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.workers.write", "com.cloudflare.api.workers.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workflows.instances.batch", "x-fern-sdk-method-name": "create"}
```
