---
title: Resume a job
page_id: operation-put-accounts-account-id-slurper-jobs-job-id-resume-577ad688
path: operations/r2-super-slurper
description: Resumes a paused R2 Super Slurper migration job, continuing the transfer from where it stopped.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/slurper/jobs/{job_id}/resume
operation_ids:
    - slurper-resume-job
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Resume a job

`PUT /accounts/{account_id}/slurper/jobs/{job_id}/resume`

Operation ID: `slurper-resume-job`

Resumes a paused R2 Super Slurper migration job, continuing the transfer from where it stopped.

## Definition

```yaml
{"operationId": "slurper-resume-job", "summary": "Resume a job", "description": "Resumes a paused R2 Super Slurper migration job, continuing the transfer from where it stopped.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "job_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Job resumed", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/r2-slurper_api-v4-success"}, {"properties": {"result": {"type": "string"}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["R2 Super Slurper"]}
```
