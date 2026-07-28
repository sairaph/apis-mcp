---
title: Get a page test schedule
page_id: operation-get-zones-zone-id-speed-api-schedule-url-632c2262
path: operations/observatory
description: Retrieves the test schedule for a page in a specific region.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/speed_api/schedule/{url}
operation_ids:
    - speed-get-scheduled-test
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a page test schedule

`GET /zones/{zone_id}/speed_api/schedule/{url}`

Operation ID: `speed-get-scheduled-test`

Retrieves the test schedule for a page in a specific region.

## Definition

```yaml
{"operationId": "speed-get-scheduled-test", "summary": "Get a page test schedule", "description": "Retrieves the test schedule for a page in a specific region.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_identifier"}}, {"name": "url", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_url"}}, {"name": "region", "in": "query", "schema": {"allOf": [{"$ref": "#/components/schemas/observatory_region"}, {"default": "us-central1", "type": "string"}]}}], "responses": {"200": {"description": "Page test schedule.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_schedule-response-single"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Observatory"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "speed.schedule", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
