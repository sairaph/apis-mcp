---
title: Create Instant Logs job
page_id: operation-post-zones-zone-id-logpush-edge-jobs-85944425
path: operations/instant-logs-jobs-for-a-zone
description: Creates a new Instant Logs job for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/logpush/edge/jobs
operation_ids:
    - post-zones-zone_id-logpush-edge-jobs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Instant Logs job

`POST /zones/{zone_id}/logpush/edge/jobs`

Operation ID: `post-zones-zone_id-logpush-edge-jobs`

Creates a new Instant Logs job for a zone.

## Definition

```yaml
{"operationId": "post-zones-zone_id-logpush-edge-jobs", "summary": "Create Instant Logs job", "description": "Creates a new Instant Logs job for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logpush_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"fields": {"$ref": "#/components/schemas/logpush_fields"}, "filter": {"$ref": "#/components/schemas/logpush_filter-2"}, "sample": {"$ref": "#/components/schemas/logpush_sample"}}}}}}, "responses": {"200": {"description": "Create Instant Logs job response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/logpush_instant_logs_job_response_single"}}}}, "4XX": {"description": "Create Instant Logs job response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logpush_instant_logs_job_response_single"}, {"$ref": "#/components/schemas/logpush_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Instant Logs jobs for a zone"], "x-api-token-group": ["Logs Read"], "x-cfPermissionsRequired": {"enum": ["#logs:read"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logpush.edge", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
