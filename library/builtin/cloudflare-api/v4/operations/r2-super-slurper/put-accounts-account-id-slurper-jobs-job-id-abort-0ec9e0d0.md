---
title: Abort a job
page_id: operation-put-accounts-account-id-slurper-jobs-job-id-abort-02699c04
path: operations/r2-super-slurper
description: Cancels a specific R2 Super Slurper migration job. Any objects in the middle of a transfer will finish, but no new objects will start transferring.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/slurper/jobs/{job_id}/abort
operation_ids:
    - slurper-abort-job
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Abort a job

`PUT /accounts/{account_id}/slurper/jobs/{job_id}/abort`

Operation ID: `slurper-abort-job`

Cancels a specific R2 Super Slurper migration job. Any objects in the middle of a transfer will finish, but no new objects will start transferring.

## Definition

```yaml
{"operationId": "slurper-abort-job", "summary": "Abort a job", "description": "Cancels a specific R2 Super Slurper migration job. Any objects in the middle of a transfer will finish, but no new objects will start transferring.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "job_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Job aborted", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/r2-slurper_api-v4-success"}, {"properties": {"result": {"type": "string"}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["R2 Super Slurper"]}
```
