---
title: List core web vital metrics trend
page_id: operation-get-zones-zone-id-speed-api-pages-url-trend-8b28be63
path: operations/observatory
description: Lists the core web vital metrics trend over time for a specific page.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/speed_api/pages/{url}/trend
operation_ids:
    - speed-list-page-trend
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List core web vital metrics trend

`GET /zones/{zone_id}/speed_api/pages/{url}/trend`

Operation ID: `speed-list-page-trend`

Lists the core web vital metrics trend over time for a specific page.

## Definition

```yaml
{"operationId": "speed-list-page-trend", "summary": "List core web vital metrics trend", "description": "Lists the core web vital metrics trend over time for a specific page.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_identifier"}}, {"name": "url", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_url"}}, {"name": "region", "in": "query", "required": true, "schema": {"$ref": "#/components/schemas/observatory_region"}}, {"name": "deviceType", "in": "query", "required": true, "schema": {"$ref": "#/components/schemas/observatory_device_type"}}, {"name": "start", "in": "query", "required": true, "schema": {"$ref": "#/components/schemas/observatory_timestamp"}}, {"name": "end", "in": "query", "schema": {"$ref": "#/components/schemas/observatory_timestamp"}}, {"name": "tz", "in": "query", "description": "The timezone of the start and end timestamps.", "required": true, "schema": {"type": "string"}, "example": "America/Chicago"}, {"name": "metrics", "in": "query", "description": "A comma-separated list of metrics to include in the results.", "required": true, "schema": {"type": "string", "example": "performanceScore,ttfb,fcp,si,lcp,tti,tbt,cls"}}], "responses": {"200": {"description": "Page trend.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_trend-response"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Observatory"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "speed.pages", "x-fern-sdk-method-name": "trend", "x-forge-hidden": true}
```
