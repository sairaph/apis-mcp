---
title: List jobs
page_id: operation-get-accounts-account-id-slurper-jobs-3f32a635
path: operations/r2-super-slurper
description: Lists all R2 Super Slurper migration jobs for the account with their status.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/slurper/jobs
operation_ids:
    - slurper-list-jobs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List jobs

`GET /accounts/{account_id}/slurper/jobs`

Operation ID: `slurper-list-jobs`

Lists all R2 Super Slurper migration jobs for the account with their status.

## Definition

```yaml
{"operationId": "slurper-list-jobs", "summary": "List jobs", "description": "Lists all R2 Super Slurper migration jobs for the account with their status.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "limit", "in": "query", "schema": {"type": "integer", "maximum": 50}}, {"name": "offset", "in": "query", "schema": {"type": "integer"}}], "responses": {"200": {"description": "A list of jobs", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/r2-slurper_api-v4-success"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/r2-slurper_JobResponse"}}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["R2 Super Slurper"]}
```
