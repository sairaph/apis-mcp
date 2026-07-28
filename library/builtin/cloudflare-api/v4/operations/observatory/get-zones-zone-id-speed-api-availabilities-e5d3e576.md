---
title: Get quota and availability
page_id: operation-get-zones-zone-id-speed-api-availabilities-cddb7d90
path: operations/observatory
description: Retrieves quota for all plans, as well as the current zone quota.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/speed_api/availabilities
operation_ids:
    - speed-get-availabilities
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get quota and availability

`GET /zones/{zone_id}/speed_api/availabilities`

Operation ID: `speed-get-availabilities`

Retrieves quota for all plans, as well as the current zone quota.

## Definition

```yaml
{"operationId": "speed-get-availabilities", "summary": "Get quota and availability", "description": "Retrieves quota for all plans, as well as the current zone quota.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_identifier"}}], "responses": {"200": {"description": "Page test availability.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_availabilities-response"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Observatory"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "speed.availabilities", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
