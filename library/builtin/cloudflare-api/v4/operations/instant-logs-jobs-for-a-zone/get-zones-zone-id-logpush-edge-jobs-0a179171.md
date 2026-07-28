---
title: List Instant Logs jobs
page_id: operation-get-zones-zone-id-logpush-edge-jobs-4b4c78dc
path: operations/instant-logs-jobs-for-a-zone
description: Lists Instant Logs jobs for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/logpush/edge/jobs
operation_ids:
    - get-zones-zone_id-logpush-edge-jobs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Instant Logs jobs

`GET /zones/{zone_id}/logpush/edge/jobs`

Operation ID: `get-zones-zone_id-logpush-edge-jobs`

Lists Instant Logs jobs for a zone.

## Definition

```yaml
{"operationId": "get-zones-zone_id-logpush-edge-jobs", "summary": "List Instant Logs jobs", "description": "Lists Instant Logs jobs for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logpush_identifier"}}], "responses": {"200": {"description": "List Instant Logs jobs response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/logpush_instant_logs_job_response_collection"}}}}, "4XX": {"description": "List Instant Logs jobs response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logpush_instant_logs_job_response_collection"}, {"$ref": "#/components/schemas/logpush_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Instant Logs jobs for a zone"], "x-api-token-group": ["Logs Read"], "x-cfPermissionsRequired": {"enum": ["#logs:read"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logpush.edge", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
