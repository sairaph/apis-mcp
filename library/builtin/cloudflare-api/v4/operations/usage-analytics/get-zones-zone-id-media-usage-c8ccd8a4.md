---
title: Get zone Media usage
page_id: operation-get-zones-zone-id-media-usage-c63c75ff
path: operations/usage-analytics
description: Retrieve Media usage analytics for a zone. This endpoint shares the same backend handler as the Stream usage endpoint and returns identical Stream metrics (streamMinutesViewed). The gateway resolves the zone to its owning account and rewrites this path to the shared usage handler.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/media/usage
operation_ids:
    - usage-analytics-get-zone-media-usage
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get zone Media usage

`GET /zones/{zone_id}/media/usage`

Operation ID: `usage-analytics-get-zone-media-usage`

Retrieve Media usage analytics for a zone. This endpoint shares the same backend handler as the Stream usage endpoint and returns identical Stream metrics (streamMinutesViewed). The gateway resolves the zone to its owning account and rewrites this path to the shared usage handler.

## Definition

```yaml
{"operationId": "usage-analytics-get-zone-media-usage", "summary": "Get zone Media usage", "description": "Retrieve Media usage analytics for a zone. This endpoint shares the same backend handler as the Stream usage endpoint and returns identical Stream metrics (streamMinutesViewed). The gateway resolves the zone to its owning account and rewrites this path to the shared usage handler.\n", "parameters": [{"$ref": "#/components/parameters/usage-analytics_zone_id"}, {"$ref": "#/components/parameters/usage-analytics_metrics"}, {"$ref": "#/components/parameters/usage-analytics_since"}, {"$ref": "#/components/parameters/usage-analytics_until"}, {"$ref": "#/components/parameters/usage-analytics_time_delta"}, {"$ref": "#/components/parameters/usage-analytics_limit"}, {"$ref": "#/components/parameters/usage-analytics_filters"}], "responses": {"200": {"description": "Usage analytics response.", "content": {"application/json": {"examples": {"success": {"summary": "Successful usage retrieval", "value": {"errors": [], "messages": [], "result": [{"streamMinutesViewed": 12500, "ts": 1693526400}, {"streamMinutesViewed": 9800, "ts": 1693530000}], "success": true}}}, "schema": {"$ref": "#/components/schemas/usage-analytics_stream_usage_response"}}}}, "400": {"$ref": "#/components/responses/usage-analytics_bad_request"}, "401": {"$ref": "#/components/responses/usage-analytics_unauthorized"}, "403": {"$ref": "#/components/responses/usage-analytics_forbidden"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Usage Analytics"], "x-api-token-group": ["Images Read"]}
```
