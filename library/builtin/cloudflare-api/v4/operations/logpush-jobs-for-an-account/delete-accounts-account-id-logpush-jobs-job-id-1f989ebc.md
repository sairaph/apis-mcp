---
title: Delete Logpush job
page_id: operation-delete-accounts-account-id-logpush-jobs-job-id-fcfbba12
path: operations/logpush-jobs-for-an-account
description: Deletes a Logpush job.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/logpush/jobs/{job_id}
operation_ids:
    - delete-accounts-account_id-logpush-jobs-job_id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Logpush job

`DELETE /accounts/{account_id}/logpush/jobs/{job_id}`

Operation ID: `delete-accounts-account_id-logpush-jobs-job_id`

Deletes a Logpush job.

## Definition

```yaml
{"operationId": "delete-accounts-account_id-logpush-jobs-job_id", "summary": "Delete Logpush job", "description": "Deletes a Logpush job.", "parameters": [{"name": "job_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logpush_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logpush_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Logpush job response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logpush_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/logpush_id"}}}}, "type": "object"}]}}}}, "4XX": {"description": "Delete Logpush job response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logpush_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logpush jobs for an account"], "x-api-token-group": ["Logs Write"], "x-cfPermissionsRequired": {"enum": ["#logs:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logpush.account-jobs", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
