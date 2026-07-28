---
title: Get job logs
page_id: operation-get-accounts-account-id-slurper-jobs-job-id-logs-bf143715
path: operations/r2-super-slurper
description: Gets log entries for an R2 Super Slurper migration job, showing migration status changes, errors, etc.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/slurper/jobs/{job_id}/logs
operation_ids:
    - slurper-get-job-logs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get job logs

`GET /accounts/{account_id}/slurper/jobs/{job_id}/logs`

Operation ID: `slurper-get-job-logs`

Gets log entries for an R2 Super Slurper migration job, showing migration status changes, errors, etc.

## Definition

```yaml
{"operationId": "slurper-get-job-logs", "summary": "Get job logs", "description": "Gets log entries for an R2 Super Slurper migration job, showing migration status changes, errors, etc.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "job_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "limit", "in": "query", "schema": {"type": "integer", "maximum": 50}}, {"name": "offset", "in": "query", "schema": {"type": "integer"}}], "responses": {"200": {"description": "Job logs", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/r2-slurper_api-v4-success"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/r2-slurper_JobLogResponse"}}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["R2 Super Slurper"]}
```
