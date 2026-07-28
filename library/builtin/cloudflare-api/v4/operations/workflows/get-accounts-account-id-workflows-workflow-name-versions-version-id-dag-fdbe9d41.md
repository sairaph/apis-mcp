---
title: Get Workflow version dag
page_id: operation-get-accounts-account-id-workflows-workflow-name-versions-version-id-dag-184c5b36
path: operations/workflows
description: Retrieves the directed acyclic graph (DAG) representation of a workflow version.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workflows/{workflow_name}/versions/{version_id}/dag
operation_ids:
    - wor-describe-workflow-versions-dag
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Workflow version dag

`GET /accounts/{account_id}/workflows/{workflow_name}/versions/{version_id}/dag`

Operation ID: `wor-describe-workflow-versions-dag`

Retrieves the directed acyclic graph (DAG) representation of a workflow version.

## Definition

```yaml
{"operationId": "wor-describe-workflow-versions-dag", "summary": "Get Workflow version dag", "description": "Retrieves the directed acyclic graph (DAG) representation of a workflow version.", "parameters": [{"name": "workflow_name", "in": "path", "required": true, "schema": {"type": "string", "maxLength": 64, "minLength": 1, "pattern": "^[a-zA-Z0-9_][a-zA-Z0-9-_]*$"}}, {"name": "version_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Get the parsed DAG for a specific workflow version.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "properties": {"class_name": {"type": "string"}, "created_on": {"type": "string", "format": "date-time"}, "dag": {"type": "object", "nullable": true}, "id": {"type": "string", "format": "uuid"}, "modified_on": {"type": "string", "format": "date-time"}, "workflow_id": {"type": "string", "format": "uuid"}}, "required": ["created_on", "modified_on", "id", "workflow_id", "class_name", "dag"]}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "cursor": {"type": "string"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}, "total_pages": {"type": "number"}}, "required": ["per_page", "count", "total_count"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "result", "messages"]}}}}, "404": {"description": "Version not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Workflows"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.workers.write", "com.cloudflare.api.workers.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-ignore": true}
```
