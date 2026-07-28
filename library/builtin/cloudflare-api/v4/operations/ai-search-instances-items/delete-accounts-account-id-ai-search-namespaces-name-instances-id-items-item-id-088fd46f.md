---
title: Delete Item.
page_id: operation-delete-accounts-account-id-ai-search-namespaces-name-instances-id-items-df3dfad7
path: operations/ai-search-instances-items
description: Deletes a file from a managed AI Search instance and triggers a reindex.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/items/{item_id}
operation_ids:
    - ai-search-namespace-instance-delete-item
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Item.

`DELETE /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/items/{item_id}`

Operation ID: `ai-search-namespace-instance-delete-item`

Deletes a file from a managed AI Search instance and triggers a reindex.

## Definition

```yaml
{"operationId": "ai-search-namespace-instance-delete-item", "summary": "Delete Item.", "description": "Deletes a file from a managed AI Search instance and triggers a reindex.", "parameters": [{"name": "id", "in": "path", "description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "required": true, "schema": {"description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "type": "string", "example": "my-ai-search", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "item_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "name", "in": "path", "description": "Namespace name", "required": true, "schema": {"type": "string"}, "example": "my-namespace"}], "responses": {"200": {"description": "Item deleted successfully.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"key": {"type": "string"}}, "required": ["key"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "This operation requires a managed instance.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "404": {"description": "Item not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Instances Items"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search.items", "x-fern-sdk-method-name": "delete", "x-forge-params": {"id": {"description": "AI Search instance ID.", "flagName": "instance-id"}, "item_id": {"description": "Indexed item ID.", "flagName": "item-id"}, "name": {"default": "default", "description": "Namespace to use for this operation.", "flagName": "namespace"}}}
```
