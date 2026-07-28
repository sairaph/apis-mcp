---
title: Get job progress
page_id: operation-get-accounts-account-id-slurper-jobs-job-id-progress-59bd918c
path: operations/r2-super-slurper
description: Retrieves current progress metrics for an R2 Super Slurper migration job
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/slurper/jobs/{job_id}/progress
operation_ids:
    - slurper-get-job-progress
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get job progress

`GET /accounts/{account_id}/slurper/jobs/{job_id}/progress`

Operation ID: `slurper-get-job-progress`

Retrieves current progress metrics for an R2 Super Slurper migration job

## Definition

```yaml
{"operationId": "slurper-get-job-progress", "summary": "Get job progress", "description": "Retrieves current progress metrics for an R2 Super Slurper migration job", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "job_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Job progress", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/r2-slurper_api-v4-success"}, {"properties": {"result": {"$ref": "#/components/schemas/r2-slurper_JobProgressResponse"}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["R2 Super Slurper"]}
```
