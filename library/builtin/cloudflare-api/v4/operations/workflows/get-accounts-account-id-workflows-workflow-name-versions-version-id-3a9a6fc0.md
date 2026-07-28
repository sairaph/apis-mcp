---
title: Get Workflow version details
page_id: operation-get-accounts-account-id-workflows-workflow-name-versions-version-id-afec45a4
path: operations/workflows
description: Retrieves details for a specific deployed workflow version.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workflows/{workflow_name}/versions/{version_id}
operation_ids:
    - wor-describe-workflow-versions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Workflow version details

`GET /accounts/{account_id}/workflows/{workflow_name}/versions/{version_id}`

Operation ID: `wor-describe-workflow-versions`

Retrieves details for a specific deployed workflow version.

## Definition

```yaml
{"operationId": "wor-describe-workflow-versions", "summary": "Get Workflow version details", "description": "Retrieves details for a specific deployed workflow version.", "parameters": [{"name": "workflow_name", "in": "path", "required": true, "schema": {"type": "string", "maxLength": 64, "minLength": 1, "pattern": "^[a-zA-Z0-9_][a-zA-Z0-9-_]*$"}}, {"name": "version_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Get specific version details.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "properties": {"class_name": {"type": "string"}, "created_on": {"type": "string", "format": "date-time"}, "default_retention": {"type": "object", "properties": {"error_retention": {"description": "Default error retention in milliseconds.", "type": "integer"}, "success_retention": {"description": "Default success retention in milliseconds.", "type": "integer"}}}, "has_dag": {"type": "boolean"}, "id": {"type": "string", "format": "uuid"}, "language": {"description": "The programming language of the workflow implementation", "type": "string", "enum": ["javascript", "python"]}, "limits": {"type": "object", "properties": {"steps": {"type": "integer", "minimum": 1}}}, "modified_on": {"type": "string", "format": "date-time"}, "workflow_id": {"type": "string", "format": "uuid"}}, "required": ["created_on", "modified_on", "id", "workflow_id", "class_name", "has_dag", "language"]}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "cursor": {"type": "string"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}, "total_pages": {"type": "number"}}, "required": ["per_page", "count", "total_count"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "result", "messages"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}, "404": {"description": "Version not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Workflows"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.workers.write", "com.cloudflare.api.workers.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workflows.versions", "x-fern-sdk-method-name": "get"}
```
