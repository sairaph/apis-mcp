---
title: Abort all jobs
page_id: operation-put-accounts-account-id-slurper-jobs-abortall-fb785834
path: operations/r2-super-slurper
description: Cancels all running R2 Super Slurper migration jobs for the account. Any objects in the middle of a transfer will finish, but no new objects will start transferring.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/slurper/jobs/abortAll
operation_ids:
    - slurper-abort-all-jobs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Abort all jobs

`PUT /accounts/{account_id}/slurper/jobs/abortAll`

Operation ID: `slurper-abort-all-jobs`

Cancels all running R2 Super Slurper migration jobs for the account. Any objects in the middle of a transfer will finish, but no new objects will start transferring.

## Definition

```yaml
{"operationId": "slurper-abort-all-jobs", "summary": "Abort all jobs", "description": "Cancels all running R2 Super Slurper migration jobs for the account. Any objects in the middle of a transfer will finish, but no new objects will start transferring.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "All jobs aborted", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/r2-slurper_api-v4-success"}, {"properties": {"result": {"type": "string"}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["R2 Super Slurper"]}
```
