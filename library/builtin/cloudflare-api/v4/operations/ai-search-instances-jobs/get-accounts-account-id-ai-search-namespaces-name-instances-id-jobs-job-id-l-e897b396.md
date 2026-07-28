---
title: List Job Logs
page_id: operation-get-accounts-account-id-ai-search-namespaces-name-instances-id-jobs-job-698f14a2
path: operations/ai-search-instances-jobs
description: Lists log entries for an AI Search indexing job.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/jobs/{job_id}/logs
operation_ids:
    - ai-search-namespace-instance-list-job-logs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Job Logs

`GET /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/jobs/{job_id}/logs`

Operation ID: `ai-search-namespace-instance-list-job-logs`

Lists log entries for an AI Search indexing job.

## Definition

```yaml
{"operationId": "ai-search-namespace-instance-list-job-logs", "summary": "List Job Logs", "description": "Lists log entries for an AI Search indexing job.", "parameters": [{"name": "id", "in": "path", "description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "required": true, "schema": {"description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "type": "string", "example": "my-ai-search", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "job_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "page", "in": "query", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"type": "integer", "default": 20, "maximum": 500, "minimum": 0}}, {"name": "name", "in": "path", "description": "Namespace name", "required": true, "schema": {"type": "string"}, "example": "my-namespace"}], "responses": {"200": {"description": "Returns a list of AI Search Job Logs.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "array", "items": {"properties": {"created_at": {"type": "number"}, "id": {"type": "integer"}, "message": {"type": "string"}, "message_type": {"type": "integer"}}, "required": ["id", "message", "message_type", "created_at"], "type": "object"}}, "result_info": {"type": "object", "properties": {"count": {"type": "integer"}, "page": {"type": "integer"}, "per_page": {"type": "integer"}, "total_count": {"type": "integer"}}, "required": ["count", "page", "per_page", "total_count"]}, "success": {"type": "boolean"}}, "required": ["success", "result", "result_info"]}}}}, "400": {"description": "Input Validation Error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}, "path": {"type": "array", "items": {"type": "string"}}}, "required": ["code", "message", "path"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "404": {"description": "Ai search not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "503": {"description": "Unable to connect to ai search.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Instances Jobs"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search.jobs", "x-fern-sdk-method-name": "logs", "x-forge-params": {"id": {"description": "AI Search instance ID.", "flagName": "instance-id"}, "job_id": {"description": "Indexing job ID.", "flagName": "job-id"}, "name": {"default": "default", "description": "Namespace to use for this operation.", "flagName": "namespace"}}}
```
