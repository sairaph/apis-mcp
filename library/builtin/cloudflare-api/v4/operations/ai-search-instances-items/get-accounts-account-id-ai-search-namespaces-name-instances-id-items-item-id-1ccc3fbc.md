---
title: List Item Chunks.
page_id: operation-get-accounts-account-id-ai-search-namespaces-name-instances-id-items-ite-e2f087a1
path: operations/ai-search-instances-items
description: Lists chunks for a specific item in an AI Search instance, including their text content.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/items/{item_id}/chunks
operation_ids:
    - ai-search-namespace-instance-list-item-chunks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Item Chunks.

`GET /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/items/{item_id}/chunks`

Operation ID: `ai-search-namespace-instance-list-item-chunks`

Lists chunks for a specific item in an AI Search instance, including their text content.

## Definition

```yaml
{"operationId": "ai-search-namespace-instance-list-item-chunks", "summary": "List Item Chunks.", "description": "Lists chunks for a specific item in an AI Search instance, including their text content.", "parameters": [{"name": "id", "in": "path", "description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "required": true, "schema": {"description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "type": "string", "example": "my-ai-search", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "item_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "limit", "in": "query", "schema": {"type": "integer", "default": 20, "maximum": 100, "minimum": 1}}, {"name": "offset", "in": "query", "schema": {"type": "integer", "default": 0, "minimum": 0}}, {"name": "name", "in": "path", "description": "Namespace name", "required": true, "schema": {"type": "string"}, "example": "my-namespace"}], "responses": {"200": {"description": "Returns the AI Search item chunks with text content.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "array", "items": {"properties": {"end_byte": {"type": "number"}, "id": {"type": "string"}, "item": {"type": "object", "properties": {"key": {"type": "string"}, "metadata": {"type": "object", "additionalProperties": true}, "timestamp": {"type": "number"}}, "required": ["key"]}, "start_byte": {"type": "number"}, "text": {"type": "string"}}, "required": ["id", "text", "item"], "type": "object"}}, "result_info": {"type": "object", "properties": {"count": {"type": "integer"}, "limit": {"type": "integer"}, "offset": {"type": "integer"}, "total": {"type": "integer"}}, "required": ["count", "total", "limit", "offset"]}, "success": {"type": "boolean"}}, "required": ["success", "result", "result_info"]}}}}, "400": {"description": "Input Validation Error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}, "path": {"type": "array", "items": {"type": "string"}}}, "required": ["code", "message", "path"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "404": {"description": "Item not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "503": {"description": "Unable to connect to ai search.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Instances Items"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search.items", "x-fern-sdk-method-name": "chunks", "x-forge-params": {"id": {"description": "AI Search instance ID.", "flagName": "instance-id"}, "item_id": {"description": "Indexed item ID.", "flagName": "item-id"}, "name": {"default": "default", "description": "Namespace to use for this operation.", "flagName": "namespace"}}}
```
