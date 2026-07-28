---
title: Deletes a Workflow
page_id: operation-delete-accounts-account-id-workflows-workflow-name-b2d85153
path: operations/workflows
description: Deletes a Workflow. This only deletes the Workflow and does not delete or modify any Worker associated to this Workflow or bounded to it.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/workflows/{workflow_name}
operation_ids:
    - wor-delete-workflow
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Deletes a Workflow

`DELETE /accounts/{account_id}/workflows/{workflow_name}`

Operation ID: `wor-delete-workflow`

Deletes a Workflow. This only deletes the Workflow and does not delete or modify any Worker associated to this Workflow or bounded to it.

## Definition

```yaml
{"operationId": "wor-delete-workflow", "summary": "Deletes a Workflow", "description": "Deletes a Workflow. This only deletes the Workflow and does not delete or modify any Worker associated to this Workflow or bounded to it.", "parameters": [{"name": "workflow_name", "in": "path", "required": true, "schema": {"type": "string", "maxLength": 64, "minLength": 1, "pattern": "^[a-zA-Z0-9_][a-zA-Z0-9-_]*$"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Deletes a Workflow.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "properties": {"status": {"type": "string", "enum": ["ok"]}, "success": {"type": "boolean", "nullable": true}}, "required": ["success", "status"]}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "cursor": {"type": "string"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}, "total_pages": {"type": "number"}}, "required": ["per_page", "count", "total_count"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "result", "messages"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}, "404": {"description": "Workflow not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Workflows"], "x-api-token-group": ["Workers Scripts Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.workers.write", "com.cloudflare.api.workers.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workflows", "x-fern-sdk-method-name": "delete", "x-forge-require-confirmation": "This operation permanently deletes the Workflow."}
```
