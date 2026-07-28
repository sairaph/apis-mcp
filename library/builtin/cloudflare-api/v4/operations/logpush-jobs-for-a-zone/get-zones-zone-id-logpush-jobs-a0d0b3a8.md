---
title: List Logpush jobs
page_id: operation-get-zones-zone-id-logpush-jobs-2d9b88a2
path: operations/logpush-jobs-for-a-zone
description: Lists Logpush jobs for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/logpush/jobs
operation_ids:
    - get-zones-zone_id-logpush-jobs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Logpush jobs

`GET /zones/{zone_id}/logpush/jobs`

Operation ID: `get-zones-zone_id-logpush-jobs`

Lists Logpush jobs for a zone.

## Definition

```yaml
{"operationId": "get-zones-zone_id-logpush-jobs", "summary": "List Logpush jobs", "description": "Lists Logpush jobs for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logpush_identifier"}}], "responses": {"200": {"description": "List Logpush jobs response.", "content": {"application/json": {"example": {"errors": [], "messages": [], "result": [{"dataset": "http_requests", "destination_conf": "s3://mybucket/logs?region=us-west-2", "enabled": false, "error_message": null, "filter": "{\"where\":{\"and\":[{\"key\":\"ClientRequestPath\",\"operator\":\"contains\",\"value\":\"/static\"},{\"key\":\"ClientRequestHost\",\"operator\":\"eq\",\"value\":\"example.com\"}]}}", "id": 1, "kind": "", "last_complete": null, "last_error": null, "max_upload_bytes": 5000000, "max_upload_interval_seconds": 30, "max_upload_records": 1000, "name": "example.com", "output_options": {"CVE-2021-44228": false, "batch_prefix": "", "batch_suffix": "", "field_delimiter": ",", "field_names": ["Datetime", "DstIP", "SrcIP"], "output_type": "ndjson", "record_delimiter": "", "record_prefix": "{", "record_suffix": "}\n", "sample_rate": 1, "timestamp_format": "unixnano"}}], "success": true}, "schema": {"$ref": "#/components/schemas/logpush_logpush_job_response_collection"}}}}, "4XX": {"description": "List Logpush jobs response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logpush_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logpush jobs for a zone"], "x-api-token-group": ["Logs Write"], "x-cfPermissionsRequired": {"enum": ["#logs:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logpush.jobs", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
