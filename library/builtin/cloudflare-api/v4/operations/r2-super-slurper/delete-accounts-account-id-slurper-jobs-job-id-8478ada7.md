---
title: Delete a job
page_id: operation-delete-accounts-account-id-slurper-jobs-job-id-78764f74
path: operations/r2-super-slurper
description: Deletes a completed, aborted, or errored R2 Super Slurper migration job. Active jobs cannot be deleted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/slurper/jobs/{job_id}
operation_ids:
    - slurper-delete-job
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a job

`DELETE /accounts/{account_id}/slurper/jobs/{job_id}`

Operation ID: `slurper-delete-job`

Deletes a completed, aborted, or errored R2 Super Slurper migration job. Active jobs cannot be deleted.

## Definition

```yaml
{"operationId": "slurper-delete-job", "summary": "Delete a job", "description": "Deletes a completed, aborted, or errored R2 Super Slurper migration job. Active jobs cannot be deleted.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "job_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Job deleted", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/r2-slurper_api-v4-success"}, {"properties": {"result": {"type": "string"}}, "type": "object"}]}}}}, "409": {"description": "Job is still active", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_api-v4-failure"}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["R2 Super Slurper"]}
```
