---
title: Get Item.
page_id: operation-get-accounts-account-id-ai-search-namespaces-name-instances-id-items-ite-7131f3e1
path: operations/ai-search-instances-items
description: Retrieves a specific indexed item from an AI Search instance.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/items/{item_id}
operation_ids:
    - ai-search-namespace-instance-get-item
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Item.

`GET /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/items/{item_id}`

Operation ID: `ai-search-namespace-instance-get-item`

Retrieves a specific indexed item from an AI Search instance.

## Definition

```yaml
{"operationId": "ai-search-namespace-instance-get-item", "summary": "Get Item.", "description": "Retrieves a specific indexed item from an AI Search instance.", "parameters": [{"name": "id", "in": "path", "description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "required": true, "schema": {"description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "type": "string", "example": "my-ai-search", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "item_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "name", "in": "path", "description": "Namespace name", "required": true, "schema": {"type": "string"}, "example": "my-namespace"}], "responses": {"200": {"description": "Returns a AI Search Item detail.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"checksum": {"type": "string", "readOnly": true}, "chunks_count": {"type": "integer", "nullable": true, "readOnly": true, "x-auditable": true}, "created_at": {"type": "string", "format": "date-time", "readOnly": true}, "error": {"type": "string", "readOnly": true}, "file_size": {"type": "number", "nullable": true, "readOnly": true, "x-auditable": true}, "id": {"type": "string", "readOnly": true}, "key": {"type": "string", "readOnly": true}, "last_seen_at": {"type": "string", "format": "date-time", "readOnly": true}, "metadata": {"description": "Built-in, configured filterable, and retained source metadata for the item.", "type": "object", "additionalProperties": {"anyOf": [{"type": "string"}, {"type": "number"}, {"type": "boolean"}]}, "nullable": true, "readOnly": true}, "namespace": {"type": "string", "readOnly": true}, "next_action": {"type": "string", "enum": ["INDEX", "DELETE", null], "nullable": true, "readOnly": true}, "source_id": {"description": "Identifies which data source this item belongs to. \"builtin\" for uploaded files, \"{type}:{source}\" for external sources, null for legacy items.", "type": "string", "nullable": true, "readOnly": true}, "status": {"type": "string", "enum": ["queued", "running", "completed", "error", "skipped", "outdated"], "readOnly": true}}, "required": ["id", "key", "status", "next_action", "checksum", "namespace", "chunks_count", "file_size", "metadata", "source_id", "last_seen_at", "created_at"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "404": {"description": "Job not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "503": {"description": "Unable to connect to ai search.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Instances Items"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search.items", "x-fern-sdk-method-name": "get", "x-forge-params": {"id": {"description": "AI Search instance ID.", "flagName": "instance-id"}, "item_id": {"description": "Indexed item ID.", "flagName": "item-id"}, "name": {"default": "default", "description": "Namespace to use for this operation.", "flagName": "namespace"}}}
```
