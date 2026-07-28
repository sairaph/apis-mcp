---
title: Get account settings
page_id: operation-get-accounts-account-id-workflows-settings-8831dd72
path: operations/workflows
description: Retrieves account-level Workflows settings, such as the default instance retention.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workflows/settings
operation_ids:
    - wor-get-workflow-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get account settings

`GET /accounts/{account_id}/workflows/settings`

Operation ID: `wor-get-workflow-settings`

Retrieves account-level Workflows settings, such as the default instance retention.

## Definition

```yaml
{"operationId": "wor-get-workflow-settings", "summary": "Get account settings", "description": "Retrieves account-level Workflows settings, such as the default instance retention.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Account settings.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "properties": {"default_retention": {"type": "object", "properties": {"error_retention": {"description": "Default error retention in milliseconds.", "type": "integer"}, "success_retention": {"description": "Default success retention in milliseconds.", "type": "integer"}}}}}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "cursor": {"type": "string"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}, "total_pages": {"type": "number"}}, "required": ["per_page", "count", "total_count"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "result", "messages"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Workflows"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.workers.write", "com.cloudflare.api.workers.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
