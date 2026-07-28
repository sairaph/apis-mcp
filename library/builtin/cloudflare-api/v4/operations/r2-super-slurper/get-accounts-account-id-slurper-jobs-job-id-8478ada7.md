---
title: Get job details
page_id: operation-get-accounts-account-id-slurper-jobs-job-id-c17a70c8
path: operations/r2-super-slurper
description: Retrieves detailed status and configuration for a specific R2 Super Slurper migration job.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/slurper/jobs/{job_id}
operation_ids:
    - slurper-get-job
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get job details

`GET /accounts/{account_id}/slurper/jobs/{job_id}`

Operation ID: `slurper-get-job`

Retrieves detailed status and configuration for a specific R2 Super Slurper migration job.

## Definition

```yaml
{"operationId": "slurper-get-job", "summary": "Get job details", "description": "Retrieves detailed status and configuration for a specific R2 Super Slurper migration job.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "job_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Job details", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/r2-slurper_api-v4-success"}, {"properties": {"result": {"$ref": "#/components/schemas/r2-slurper_JobResponse"}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["R2 Super Slurper"]}
```
