---
title: Send event to instance
page_id: operation-post-accounts-account-id-workflows-workflow-name-instances-instance-id-e-763350e3
path: operations/workflows
description: Sends an event to a running workflow instance to trigger state transitions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/workflows/{workflow_name}/instances/{instance_id}/events/{event_type}
operation_ids:
    - wor-send-event-workflow-instance
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Send event to instance

`POST /accounts/{account_id}/workflows/{workflow_name}/instances/{instance_id}/events/{event_type}`

Operation ID: `wor-send-event-workflow-instance`

Sends an event to a running workflow instance to trigger state transitions.

## Definition

```yaml
{"operationId": "wor-send-event-workflow-instance", "summary": "Send event to instance", "description": "Sends an event to a running workflow instance to trigger state transitions.", "parameters": [{"name": "workflow_name", "in": "path", "required": true, "schema": {"type": "string", "maxLength": 64, "minLength": 1, "pattern": "^[a-zA-Z0-9_][a-zA-Z0-9-_]*$"}}, {"name": "instance_id", "in": "path", "required": true, "schema": {"description": "Instance identifier. User-created instances match `^[a-zA-Z0-9_][a-zA-Z0-9-_]*$` (max 100 characters); cron-triggered instances can use a longer, system-generated id derived from the cron expression.", "type": "string", "maxLength": 271, "minLength": 1}}, {"name": "event_type", "in": "path", "required": true, "schema": {"type": "string", "maxLength": 100, "minLength": 1, "pattern": "^[a-zA-Z0-9_][a-zA-Z0-9-_]*$"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object"}}}}, "responses": {"200": {"description": "Send an event to an instance.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object"}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "cursor": {"type": "string"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}, "total_pages": {"type": "number"}}, "required": ["per_page", "count", "total_count"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}, "404": {"description": "Workflow not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Workflows"], "x-api-token-group": ["Workers Scripts Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.workers.write", "com.cloudflare.api.workers.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workflows.instances.events", "x-fern-sdk-method-name": "send"}
```
