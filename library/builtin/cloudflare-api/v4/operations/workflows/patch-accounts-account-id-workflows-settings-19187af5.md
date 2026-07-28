---
title: Update account settings
page_id: operation-patch-accounts-account-id-workflows-settings-0455ffce
path: operations/workflows
description: Partially updates account-level Workflows settings; only the fields present in the body are changed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/workflows/settings
operation_ids:
    - wor-update-workflow-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update account settings

`PATCH /accounts/{account_id}/workflows/settings`

Operation ID: `wor-update-workflow-settings`

Partially updates account-level Workflows settings; only the fields present in the body are changed.

## Definition

```yaml
{"operationId": "wor-update-workflow-settings", "summary": "Update account settings", "description": "Partially updates account-level Workflows settings; only the fields present in the body are changed.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"default_retention": {"description": "Default retention applied to instances of this version when they do not set their own retention.", "type": "object", "properties": {"error_retention": {"description": "Specifies the duration in milliseconds or as a string like '5 minutes'.", "anyOf": [{"description": "Specifies the duration in milliseconds.", "type": "number"}, {"type": "string"}], "oneOf": [{"description": "Specifies the duration in milliseconds.", "type": "integer"}, {"example": "5 minutes", "pattern": "^(\\d+)\\s+(second|minute|hour|day|week|month|year)s?$", "type": "string"}]}, "success_retention": {"description": "Specifies the duration in milliseconds or as a string like '5 minutes'.", "anyOf": [{"description": "Specifies the duration in milliseconds.", "type": "number"}, {"type": "string"}], "oneOf": [{"description": "Specifies the duration in milliseconds.", "type": "integer"}, {"example": "5 minutes", "pattern": "^(\\d+)\\s+(second|minute|hour|day|week|month|year)s?$", "type": "string"}]}}}}}}}}, "responses": {"200": {"description": "Updated account settings.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "properties": {"default_retention": {"type": "object", "properties": {"error_retention": {"description": "Default error retention in milliseconds.", "type": "integer"}, "success_retention": {"description": "Default success retention in milliseconds.", "type": "integer"}}}}}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "cursor": {"type": "string"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}, "total_pages": {"type": "number"}}, "required": ["per_page", "count", "total_count"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "result", "messages"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Workflows"], "x-api-token-group": ["Workers Scripts Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.workers.write", "com.cloudflare.api.workers.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
