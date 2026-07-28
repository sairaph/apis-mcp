---
title: Add script triggers
page_id: operation-patch-accounts-account-id-triggers-script-name-45383fdb
path: operations/workflows
description: Adds event trigger declarations without removing existing declarations owned by the script.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/triggers/{script_name}
operation_ids:
    - wor-add-script-triggers
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add script triggers

`PATCH /accounts/{account_id}/triggers/{script_name}`

Operation ID: `wor-add-script-triggers`

Adds event trigger declarations without removing existing declarations owned by the script.

## Definition

```yaml
{"operationId": "wor-add-script-triggers", "summary": "Add script triggers", "description": "Adds event trigger declarations without removing existing declarations owned by the script.", "parameters": [{"name": "script_name", "in": "path", "required": true, "schema": {"type": "string", "maxLength": 255, "minLength": 1}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"filter": {"type": "object", "default": {}, "additionalProperties": {"anyOf": [{"maxLength": 1024, "type": "string"}, {"type": "number"}, {"type": "boolean"}]}}, "targets": {"type": "array", "items": {"properties": {"script_name": {"type": "string", "maxLength": 255, "minLength": 1}, "type": {"type": "string", "enum": ["workflow"]}, "workflow_name": {"type": "string", "maxLength": 255, "minLength": 1}}, "required": ["type", "script_name", "workflow_name"], "type": "object"}, "maxItems": 100, "minItems": 1}, "type": {"type": "string", "minLength": 1}}, "required": ["type", "targets"], "type": "object"}, "maxItems": 100}}}}, "responses": {"200": {"description": "Effective script trigger declarations.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "properties": {"script_name": {"type": "string"}, "triggers": {"type": "array", "items": {"properties": {"filter": {"type": "object", "default": {}, "additionalProperties": {"anyOf": [{"maxLength": 1024, "type": "string"}, {"type": "number"}, {"type": "boolean"}]}}, "targets": {"type": "array", "items": {"properties": {"script_name": {"type": "string", "maxLength": 255, "minLength": 1}, "type": {"type": "string", "enum": ["workflow"]}, "workflow_name": {"type": "string", "maxLength": 255, "minLength": 1}}, "required": ["type", "script_name", "workflow_name"], "type": "object"}, "maxItems": 100, "minItems": 1}, "type": {"type": "string", "minLength": 1}}, "required": ["type", "targets"], "type": "object"}, "maxItems": 100}}, "required": ["script_name", "triggers"]}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "cursor": {"type": "string"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}, "total_pages": {"type": "number"}}, "required": ["per_page", "count", "total_count"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "result", "messages"]}}}}, "400": {"description": "Invalid trigger configuration.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}, "404": {"description": "Target workflow not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Workflows"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.workers.write", "com.cloudflare.api.workers.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-ignore": true}
```
