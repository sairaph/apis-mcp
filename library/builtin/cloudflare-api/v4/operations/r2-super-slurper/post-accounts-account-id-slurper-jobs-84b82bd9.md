---
title: Create a job
page_id: operation-post-accounts-account-id-slurper-jobs-92b2567a
path: operations/r2-super-slurper
description: Creates a new R2 Super Slurper migration job to transfer objects from a source bucket (e.g. S3, GCS, R2) to R2.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/slurper/jobs
operation_ids:
    - slurper-create-job
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a job

`POST /accounts/{account_id}/slurper/jobs`

Operation ID: `slurper-create-job`

Creates a new R2 Super Slurper migration job to transfer objects from a source bucket (e.g. S3, GCS, R2) to R2.

## Definition

```yaml
{"operationId": "slurper-create-job", "summary": "Create a job", "description": "Creates a new R2 Super Slurper migration job to transfer objects from a source bucket (e.g. S3, GCS, R2) to R2.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_CreateJobRequest"}}}}, "responses": {"201": {"description": "Job created", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/r2-slurper_api-v4-success"}, {"properties": {"result": {"type": "object", "properties": {"id": {"type": "string"}}}}, "type": "object"}]}}}}, "409": {"description": "Maximum number of concurrent jobs has been reached", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_api-v4-failure"}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["R2 Super Slurper"]}
```
