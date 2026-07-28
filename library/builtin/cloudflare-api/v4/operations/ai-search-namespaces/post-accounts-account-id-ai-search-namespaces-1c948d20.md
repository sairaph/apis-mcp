---
title: Create a namespace
page_id: operation-post-accounts-account-id-ai-search-namespaces-f874350f
path: operations/ai-search-namespaces
description: Create a namespace for organizing AI Search instances.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/ai-search/namespaces
operation_ids:
    - ai-search-create-namespace
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a namespace

`POST /accounts/{account_id}/ai-search/namespaces`

Operation ID: `ai-search-create-namespace`

Create a namespace for organizing AI Search instances.

## Definition

```yaml
{"operationId": "ai-search-create-namespace", "summary": "Create a namespace", "description": "Create a namespace for organizing AI Search instances.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"description": "Optional description for the namespace. Max 256 characters.", "type": "string", "example": "Production environment", "maxLength": 256, "nullable": true}, "name": {"type": "string", "pattern": "^[a-z0-9]([a-z0-9-]{0,26}[a-z0-9])?$"}}, "required": ["name"]}}}}, "responses": {"201": {"description": "Namespace created.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time", "readOnly": true}, "description": {"description": "Optional description for the namespace. Max 256 characters.", "type": "string", "example": "Production environment", "maxLength": 256, "nullable": true}, "name": {"type": "string", "example": "production", "pattern": "^[a-z0-9]([a-z0-9-]{0,26}[a-z0-9])?$"}}, "required": ["name", "created_at"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result"]}}}}, "400": {"description": "Input Validation Error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}, "path": {"type": "array", "items": {"type": "string"}}}, "required": ["code", "message", "path"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "403": {"description": "Max namespaces reached.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "409": {"description": "Namespace already exists.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Namespaces"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search.namespace", "x-fern-sdk-method-name": "create", "x-forge-params": {"name": {"description": "Name for the new namespace.", "flagName": "namespace", "positional": true}}}
```
