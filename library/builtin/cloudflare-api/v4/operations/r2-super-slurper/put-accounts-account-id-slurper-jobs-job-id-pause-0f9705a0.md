---
title: Pause a job
page_id: operation-put-accounts-account-id-slurper-jobs-job-id-pause-f1ea97ea
path: operations/r2-super-slurper
description: Pauses a running R2 Super Slurper migration job. The job can be resumed later to continue transferring.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/slurper/jobs/{job_id}/pause
operation_ids:
    - slurper-pause-job
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Pause a job

`PUT /accounts/{account_id}/slurper/jobs/{job_id}/pause`

Operation ID: `slurper-pause-job`

Pauses a running R2 Super Slurper migration job. The job can be resumed later to continue transferring.

## Definition

```yaml
{"operationId": "slurper-pause-job", "summary": "Pause a job", "description": "Pauses a running R2 Super Slurper migration job. The job can be resumed later to continue transferring.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "job_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Job paused", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/r2-slurper_api-v4-success"}, {"properties": {"result": {"type": "string"}}, "type": "object"}]}}}}, "409": {"description": "Job is not paused", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_api-v4-failure"}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["R2 Super Slurper"]}
```
