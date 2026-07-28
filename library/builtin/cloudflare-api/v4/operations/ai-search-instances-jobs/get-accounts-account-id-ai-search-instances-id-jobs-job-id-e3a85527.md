---
title: Get a Job Details
page_id: operation-get-accounts-account-id-ai-search-instances-id-jobs-job-id-d3f05585
path: operations/ai-search-instances-jobs
description: Retrieves details for a specific AI Search indexing job.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-search/instances/{id}/jobs/{job_id}
operation_ids:
    - ai-search-instance-get-job
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a Job Details

`GET /accounts/{account_id}/ai-search/instances/{id}/jobs/{job_id}`

Operation ID: `ai-search-instance-get-job`

Retrieves details for a specific AI Search indexing job.

## Definition

```yaml
{"operationId": "ai-search-instance-get-job", "summary": "Get a Job Details", "description": "Retrieves details for a specific AI Search indexing job.", "parameters": [{"name": "id", "in": "path", "description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "required": true, "schema": {"description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "type": "string", "example": "my-ai-search", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "job_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}], "responses": {"200": {"description": "Returns a AI Search Job Details.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"description": {"type": "string"}, "end_reason": {"type": "string"}, "ended_at": {"type": "string"}, "id": {"type": "string", "x-auditable": true}, "last_seen_at": {"type": "string"}, "source": {"type": "string", "enum": ["user", "schedule"], "x-auditable": true}, "started_at": {"type": "string"}}, "required": ["id", "source"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "404": {"description": "Job not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "503": {"description": "Unable to connect to ai search.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Instances Jobs"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-ignore": true, "x-fern-sdk-group-name": "ai-search.jobs", "x-fern-sdk-method-name": "get"}
```
