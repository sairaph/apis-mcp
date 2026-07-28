---
title: Update Logpush job
page_id: operation-put-accounts-account-id-logpush-jobs-job-id-ab5296ff
path: operations/logpush-jobs-for-an-account
description: Updates a Logpush job.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/logpush/jobs/{job_id}
operation_ids:
    - put-accounts-account_id-logpush-jobs-job_id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Logpush job

`PUT /accounts/{account_id}/logpush/jobs/{job_id}`

Operation ID: `put-accounts-account_id-logpush-jobs-job_id`

Updates a Logpush job.

## Definition

```yaml
{"operationId": "put-accounts-account_id-logpush-jobs-job_id", "summary": "Update Logpush job", "description": "Updates a Logpush job.", "parameters": [{"name": "job_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logpush_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logpush_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"example": {"destination_conf": "s3://mybucket/logs?region=us-west-2", "enabled": false, "filter": "{\"where\":{\"and\":[{\"key\":\"ClientRequestPath\",\"operator\":\"contains\",\"value\":\"/static\"},{\"key\":\"ClientRequestHost\",\"operator\":\"eq\",\"value\":\"example.com\"}]}}", "kind": "", "max_upload_bytes": 5000000, "max_upload_interval_seconds": 30, "max_upload_records": 1000, "output_options": {"CVE-2021-44228": false, "batch_prefix": "", "batch_suffix": "", "field_delimiter": ",", "field_names": ["Datetime", "DstIP", "SrcIP"], "output_type": "ndjson", "record_delimiter": "", "record_prefix": "{", "record_suffix": "}\n", "sample_rate": 1, "timestamp_format": "unixnano"}, "ownership_challenge": "00000000000000000000"}, "schema": {"type": "object", "properties": {"destination_conf": {"$ref": "#/components/schemas/logpush_destination_conf"}, "enabled": {"$ref": "#/components/schemas/logpush_enabled"}, "filter": {"$ref": "#/components/schemas/logpush_filter"}, "frequency": {"$ref": "#/components/schemas/logpush_frequency"}, "kind": {"$ref": "#/components/schemas/logpush_kind"}, "logpull_options": {"$ref": "#/components/schemas/logpush_logpull_options"}, "max_upload_bytes": {"$ref": "#/components/schemas/logpush_max_upload_bytes"}, "max_upload_interval_seconds": {"$ref": "#/components/schemas/logpush_max_upload_interval_seconds"}, "max_upload_records": {"$ref": "#/components/schemas/logpush_max_upload_records"}, "name": {"$ref": "#/components/schemas/logpush_name"}, "output_options": {"$ref": "#/components/schemas/logpush_output_options"}, "ownership_challenge": {"$ref": "#/components/schemas/logpush_ownership_challenge"}}}}}}, "responses": {"200": {"description": "Update Logpush job response.", "content": {"application/json": {"example": {"errors": [], "messages": [], "result": {"dataset": "gateway_dns", "destination_conf": "s3://mybucket/logs?region=us-west-2", "enabled": false, "error_message": null, "filter": "{\"where\":{\"and\":[{\"key\":\"ClientRequestPath\",\"operator\":\"contains\",\"value\":\"/static\"},{\"key\":\"ClientRequestHost\",\"operator\":\"eq\",\"value\":\"example.com\"}]}}", "id": 1, "kind": "", "last_complete": null, "last_error": null, "max_upload_bytes": 5000000, "max_upload_interval_seconds": 30, "max_upload_records": 1000, "name": "example.com", "output_options": {"CVE-2021-44228": false, "batch_prefix": "", "batch_suffix": "", "field_delimiter": ",", "field_names": ["Datetime", "DstIP", "SrcIP"], "output_type": "ndjson", "record_delimiter": "", "record_prefix": "{", "record_suffix": "}\n", "sample_rate": 1, "timestamp_format": "unixnano"}}, "success": true}, "schema": {"$ref": "#/components/schemas/logpush_logpush_job_response_single"}}}}, "4XX": {"description": "Update Logpush job response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logpush_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logpush jobs for an account"], "x-api-token-group": ["Logs Write"], "x-cfPermissionsRequired": {"enum": ["#logs:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logpush.account-jobs", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
