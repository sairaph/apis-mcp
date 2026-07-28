---
title: Get a Job Details
page_id: operation-get-accounts-account-id-autorag-rags-id-jobs-job-id-4380d0a7
path: operations/autorag-jobs
description: Returns details for an AutoRAG job.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/autorag/rags/{id}/jobs/{job_id}
operation_ids:
    - autorag-config-get-job
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a Job Details

`GET /accounts/{account_id}/autorag/rags/{id}/jobs/{job_id}`

Operation ID: `autorag-config-get-job`

Returns details for an AutoRAG job.

## Definition

```yaml
{"operationId": "autorag-config-get-job", "summary": "Get a Job Details", "description": "Returns details for an AutoRAG job.", "parameters": [{"name": "id", "in": "path", "description": "rag id", "required": true, "schema": {"description": "rag id", "type": "string", "example": "my-rag", "maxLength": 32, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$"}}, {"name": "job_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}], "responses": {"200": {"description": "Returns a AutoRAG Job Details", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"end_reason": {"type": "string"}, "ended_at": {"type": "string"}, "id": {"type": "string"}, "last_seen_at": {"type": "string"}, "source": {"type": "string", "enum": ["user", "schedule"]}, "started_at": {"type": "string"}}, "required": ["id", "source"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "404": {"description": "job_not_found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7021}, "message": {"type": "string", "example": "job_not_found"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "503": {"description": "unable_to_connect_to_autorag", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7017}, "message": {"type": "string", "example": "unable_to_connect_to_autorag"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AutoRAG Jobs"], "x-api-token-group": ["Auto Rag Write", "Auto Rag Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.rag"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-ignore": true, "x-fern-sdk-group-name": "autorag.jobs", "x-fern-sdk-method-name": "get"}
```
