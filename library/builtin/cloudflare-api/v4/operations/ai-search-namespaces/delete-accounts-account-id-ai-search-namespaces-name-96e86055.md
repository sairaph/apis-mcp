---
title: Delete a namespace
page_id: operation-delete-accounts-account-id-ai-search-namespaces-name-2234a6cf
path: operations/ai-search-namespaces
description: Permanently delete a namespace. The namespace must be empty (no instances), and the default namespace cannot be deleted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/ai-search/namespaces/{name}
operation_ids:
    - ai-search-delete-namespace
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a namespace

`DELETE /accounts/{account_id}/ai-search/namespaces/{name}`

Operation ID: `ai-search-delete-namespace`

Permanently delete a namespace. The namespace must be empty (no instances), and the default namespace cannot be deleted.

## Definition

```yaml
{"operationId": "ai-search-delete-namespace", "summary": "Delete a namespace", "description": "Permanently delete a namespace. The namespace must be empty (no instances), and the default namespace cannot be deleted.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "name", "in": "path", "required": true, "schema": {"type": "string", "example": "production"}}], "responses": {"200": {"description": "Namespace deleted.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "additionalProperties": false}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result"]}}}}, "400": {"description": "Namespace not empty.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "404": {"description": "Namespace not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Namespaces"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search.namespace", "x-fern-sdk-method-name": "delete", "x-forge-params": {"name": {"description": "Namespace to delete.", "flagName": "namespace"}}, "x-forge-require-confirmation": "This operation permanently deletes the namespace."}
```
