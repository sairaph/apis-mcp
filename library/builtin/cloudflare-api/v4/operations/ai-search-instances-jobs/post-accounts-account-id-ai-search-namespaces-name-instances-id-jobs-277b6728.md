---
title: Create new job
page_id: operation-post-accounts-account-id-ai-search-namespaces-name-instances-id-jobs-2693d2af
path: operations/ai-search-instances-jobs
description: Creates a new indexing job for an AI Search instance.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/jobs
operation_ids:
    - ai-search-namespace-instance-create-job
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create new job

`POST /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/jobs`

Operation ID: `ai-search-namespace-instance-create-job`

Creates a new indexing job for an AI Search instance.

## Definition

```yaml
{"operationId": "ai-search-namespace-instance-create-job", "summary": "Create new job", "description": "Creates a new indexing job for an AI Search instance.", "parameters": [{"name": "id", "in": "path", "description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "required": true, "schema": {"description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "type": "string", "example": "my-ai-search", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "name", "in": "path", "description": "Namespace name", "required": true, "schema": {"type": "string"}, "example": "my-namespace"}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"type": "string", "maxLength": 255}}}}}}, "responses": {"200": {"description": "Returns the AI Search job id.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"description": {"type": "string"}, "end_reason": {"type": "string"}, "ended_at": {"type": "string"}, "id": {"type": "string", "x-auditable": true}, "last_seen_at": {"type": "string"}, "source": {"type": "string", "enum": ["user", "schedule"], "x-auditable": true}, "started_at": {"type": "string"}}, "required": ["id", "source"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Input Validation Error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}, "path": {"type": "array", "items": {"type": "string"}}}, "required": ["code", "message", "path"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "404": {"description": "Ai search not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "429": {"description": "Sync in cooldown.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "503": {"description": "Unable to connect to ai search.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Instances Jobs"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search.jobs", "x-fern-sdk-method-name": "create", "x-forge-params": {"description": {"description": "Optional description for the indexing job."}, "id": {"description": "AI Search instance ID.", "flagName": "instance-id"}, "name": {"default": "default", "description": "Namespace to use for this operation.", "flagName": "namespace"}}}
```
