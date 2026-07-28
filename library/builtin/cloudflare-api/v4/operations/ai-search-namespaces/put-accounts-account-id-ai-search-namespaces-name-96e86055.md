---
title: Update a namespace
page_id: operation-put-accounts-account-id-ai-search-namespaces-name-d78a2dee
path: operations/ai-search-namespaces
description: Update the description of an existing namespace. The default namespace cannot be modified.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/ai-search/namespaces/{name}
operation_ids:
    - ai-search-update-namespace
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a namespace

`PUT /accounts/{account_id}/ai-search/namespaces/{name}`

Operation ID: `ai-search-update-namespace`

Update the description of an existing namespace. The default namespace cannot be modified.

## Definition

```yaml
{"operationId": "ai-search-update-namespace", "summary": "Update a namespace", "description": "Update the description of an existing namespace. The default namespace cannot be modified.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "name", "in": "path", "required": true, "schema": {"type": "string", "example": "production"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"description": "Optional description for the namespace. Max 256 characters.", "type": "string", "example": "Production environment", "maxLength": 256, "nullable": true}}}}}}, "responses": {"200": {"description": "Returns the updated namespace.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time", "readOnly": true}, "description": {"description": "Optional description for the namespace. Max 256 characters.", "type": "string", "example": "Production environment", "maxLength": 256, "nullable": true}, "name": {"type": "string", "example": "production", "pattern": "^[a-z0-9]([a-z0-9-]{0,26}[a-z0-9])?$"}}, "required": ["name", "created_at"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result"]}}}}, "400": {"description": "Cannot modify default namespace.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "404": {"description": "Namespace not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Namespaces"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search.namespace", "x-fern-sdk-method-name": "update", "x-forge-params": {"name": {"description": "Namespace to update.", "flagName": "namespace"}}}
```
