---
title: Item Logs.
page_id: operation-get-accounts-account-id-ai-search-namespaces-name-instances-id-items-ite-46c9d6c8
path: operations/ai-search-instances-items
description: Lists processing logs for a specific item in an AI Search instance.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/items/{item_id}/logs
operation_ids:
    - ai-search-namespace-instance-logs-item
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Item Logs.

`GET /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/items/{item_id}/logs`

Operation ID: `ai-search-namespace-instance-logs-item`

Lists processing logs for a specific item in an AI Search instance.

## Definition

```yaml
{"operationId": "ai-search-namespace-instance-logs-item", "summary": "Item Logs.", "description": "Lists processing logs for a specific item in an AI Search instance.", "parameters": [{"name": "id", "in": "path", "description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "required": true, "schema": {"description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "type": "string", "example": "my-ai-search", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "item_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "limit", "in": "query", "schema": {"type": "integer", "default": 50, "maximum": 100, "minimum": 1}}, {"name": "cursor", "in": "query", "schema": {"type": "string", "maxLength": 512}}, {"name": "name", "in": "path", "description": "Namespace name", "required": true, "schema": {"type": "string"}, "example": "my-namespace"}], "responses": {"200": {"description": "Returns the AI Search item logs.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "array", "items": {"properties": {"action": {"type": "string"}, "chunkCount": {"type": "integer", "nullable": true}, "errorType": {"type": "string", "nullable": true}, "fileKey": {"type": "string"}, "message": {"type": "string", "nullable": true}, "processingTimeMs": {"type": "integer", "nullable": true}, "timestamp": {"type": "string", "format": "date-time"}}, "required": ["timestamp", "action", "message", "fileKey", "chunkCount", "processingTimeMs", "errorType"], "type": "object"}}, "result_info": {"type": "object", "properties": {"count": {"type": "integer"}, "cursor": {"type": "string", "nullable": true}, "per_page": {"type": "integer"}, "truncated": {"type": "boolean"}}, "required": ["count", "per_page", "cursor", "truncated"]}, "success": {"type": "boolean"}}, "required": ["success", "result", "result_info"]}}}}, "400": {"description": "Input Validation Error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}, "path": {"type": "array", "items": {"type": "string"}}}, "required": ["code", "message", "path"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "404": {"description": "Item not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "503": {"description": "Unable to connect to ai search.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Instances Items"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search.items", "x-fern-sdk-method-name": "logs", "x-forge-params": {"id": {"description": "AI Search instance ID.", "flagName": "instance-id"}, "item_id": {"description": "Indexed item ID.", "flagName": "item-id"}, "name": {"default": "default", "description": "Namespace to use for this operation.", "flagName": "namespace"}}}
```
