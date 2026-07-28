---
title: Get Logpush job details
page_id: operation-get-accounts-account-id-logpush-jobs-job-id-573bee8c
path: operations/logpush-jobs-for-an-account
description: Gets the details of a Logpush job.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/logpush/jobs/{job_id}
operation_ids:
    - get-accounts-account_id-logpush-jobs-job_id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Logpush job details

`GET /accounts/{account_id}/logpush/jobs/{job_id}`

Operation ID: `get-accounts-account_id-logpush-jobs-job_id`

Gets the details of a Logpush job.

## Definition

```yaml
{"operationId": "get-accounts-account_id-logpush-jobs-job_id", "summary": "Get Logpush job details", "description": "Gets the details of a Logpush job.", "parameters": [{"name": "job_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logpush_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logpush_identifier"}}], "responses": {"200": {"description": "Get Logpush job details response.", "content": {"application/json": {"example": {"errors": [], "messages": [], "result": {"dataset": "gateway_dns", "destination_conf": "s3://mybucket/logs?region=us-west-2", "enabled": false, "error_message": null, "filter": "{\"where\":{\"and\":[{\"key\":\"ClientRequestPath\",\"operator\":\"contains\",\"value\":\"/static\"},{\"key\":\"ClientRequestHost\",\"operator\":\"eq\",\"value\":\"example.com\"}]}}", "id": 1, "kind": "", "last_complete": null, "last_error": null, "max_upload_bytes": 5000000, "max_upload_interval_seconds": 30, "max_upload_records": 1000, "name": "example.com", "output_options": {"CVE-2021-44228": false, "batch_prefix": "", "batch_suffix": "", "field_delimiter": ",", "field_names": ["Datetime", "DstIP", "SrcIP"], "output_type": "ndjson", "record_delimiter": "", "record_prefix": "{", "record_suffix": "}\n", "sample_rate": 1, "timestamp_format": "unixnano"}}, "success": true}, "schema": {"$ref": "#/components/schemas/logpush_logpush_job_response_single"}}}}, "4XX": {"description": "Get Logpush job details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logpush_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logpush jobs for an account"], "x-api-token-group": ["Logs Write"], "x-cfPermissionsRequired": {"enum": ["#logs:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logpush.account-jobs", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
