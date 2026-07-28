---
title: Download Item Content.
page_id: operation-get-accounts-account-id-ai-search-namespaces-name-instances-id-items-ite-7ae01eed
path: operations/ai-search-instances-items
description: Downloads the raw file content for a specific item from the managed AI Search instance storage.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/items/{item_id}/download
operation_ids:
    - ai-search-namespace-instance-get-item-content
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Download Item Content.

`GET /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/items/{item_id}/download`

Operation ID: `ai-search-namespace-instance-get-item-content`

Downloads the raw file content for a specific item from the managed AI Search instance storage.

## Definition

```yaml
{"operationId": "ai-search-namespace-instance-get-item-content", "summary": "Download Item Content.", "description": "Downloads the raw file content for a specific item from the managed AI Search instance storage.", "parameters": [{"name": "id", "in": "path", "description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "required": true, "schema": {"description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "type": "string", "example": "my-ai-search", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "item_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "name", "in": "path", "description": "Namespace name", "required": true, "schema": {"type": "string"}, "example": "my-namespace"}], "responses": {"200": {"description": "Raw file content.", "content": {"application/octet-stream": {"schema": {"type": "string", "format": "binary"}}}}, "400": {"description": "Content download not available for external source items.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "404": {"description": "Item not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "503": {"description": "Unable to connect to ai search.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Instances Items"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search.items", "x-fern-sdk-method-name": "download", "x-forge-params": {"id": {"description": "AI Search instance ID.", "flagName": "instance-id"}, "item_id": {"description": "Indexed item ID.", "flagName": "item-id"}, "name": {"default": "default", "description": "Namespace to use for this operation.", "flagName": "namespace"}}}
```
