---
title: Batch terminate instances of a workflow
page_id: operation-post-accounts-account-id-workflows-workflow-name-instances-batch-termina-1bc6f56e
path: operations/workflows
description: Terminates multiple workflow instances in a single batch operation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/workflows/{workflow_name}/instances/batch/terminate
operation_ids:
    - wor-batch-terminate-workflow-instances
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Batch terminate instances of a workflow

`POST /accounts/{account_id}/workflows/{workflow_name}/instances/batch/terminate`

Operation ID: `wor-batch-terminate-workflow-instances`

Terminates multiple workflow instances in a single batch operation.

## Definition

```yaml
{"operationId": "wor-batch-terminate-workflow-instances", "summary": "Batch terminate instances of a workflow", "description": "Terminates multiple workflow instances in a single batch operation.", "parameters": [{"name": "workflow_name", "in": "path", "required": true, "schema": {"type": "string", "maxLength": 64, "minLength": 1, "pattern": "^[a-zA-Z0-9_][a-zA-Z0-9-_]*$"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "array", "items": {"maximum": 100, "minimum": 1, "pattern": "^[a-zA-Z0-9_][a-zA-Z0-9-_]*$", "type": "string"}, "maxItems": 100, "minItems": 1}}}}, "responses": {"200": {"description": "Batch terminate instances of a workflow, via a async job. Body is a JSON list that contain the ids of the instances to terminate.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "properties": {"instancesTerminated": {"type": "number"}, "status": {"type": "string", "enum": ["ok", "already_running"]}}, "required": ["status", "instancesTerminated"]}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "cursor": {"type": "string"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}, "total_pages": {"type": "number"}}, "required": ["per_page", "count", "total_count"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "result", "messages"]}}}}, "400": {"description": "Provided Workflow ID is not valid.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}, "404": {"description": "Workflow Name not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Workflows"], "x-api-token-group": ["Workers Scripts Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.workers.write", "com.cloudflare.api.workers.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-ignore": true}
```
