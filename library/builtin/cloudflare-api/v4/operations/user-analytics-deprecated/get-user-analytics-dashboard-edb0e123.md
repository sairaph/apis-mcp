---
title: Get user analytics dashboard
page_id: operation-get-user-analytics-dashboard-d0cd4c7f
path: operations/user-analytics-deprecated
description: |-
    The user analytics dashboard provides totals and timeseries data aggregated
    across all zones owned by the authenticated user for the given time period.
    Only zones for which the user has the `#analytics:read` permission are included.

    This endpoint is deprecated. Please use the GraphQL Analytics API instead:
    https://developers.cloudflare.com/analytics/graphql-api/
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/analytics/dashboard
operation_ids:
    - user-analytics-get-dashboard
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get user analytics dashboard

`GET /user/analytics/dashboard`

Operation ID: `user-analytics-get-dashboard`

The user analytics dashboard provides totals and timeseries data aggregated
across all zones owned by the authenticated user for the given time period.
Only zones for which the user has the `#analytics:read` permission are included.

This endpoint is deprecated. Please use the GraphQL Analytics API instead:
https://developers.cloudflare.com/analytics/graphql-api/

## Definition

```yaml
{"operationId": "user-analytics-get-dashboard", "summary": "Get user analytics dashboard", "description": "The user analytics dashboard provides totals and timeseries data aggregated\nacross all zones owned by the authenticated user for the given time period.\nOnly zones for which the user has the `#analytics:read` permission are included.\n\nThis endpoint is deprecated. Please use the GraphQL Analytics API instead:\nhttps://developers.cloudflare.com/analytics/graphql-api/", "parameters": [{"name": "since", "in": "query", "schema": {"$ref": "#/components/schemas/zone-analytics-api_since"}}, {"name": "until", "in": "query", "schema": {"$ref": "#/components/schemas/zone-analytics-api_until"}}, {"name": "continuous", "in": "query", "schema": {"description": "When set to true, the API will move the requested time window backward,\nuntil it finds a region with completely aggregated data.\n\nThe API response _may not represent the requested time window_.", "type": "boolean", "default": true}}], "responses": {"200": {"description": "Get user analytics dashboard response", "content": {"application/json": {"examples": {"success": {"summary": "Successful user analytics dashboard response", "value": {"errors": [], "messages": [], "query": {"since": "2015-01-01T12:23:00Z", "until": "2015-01-02T12:23:00Z"}, "result": [{"timeseries": [{"bandwidth": {"all": 190290, "cached": 97717, "uncached": 92573}, "requests": {"all": 51440}, "since": "2015-01-01T12:23:00Z", "until": "2015-01-01T13:23:00Z"}], "totals": {"bandwidth": {"all": 4567890, "cached": 2345678, "uncached": 2222212}, "requests": {"all": 1234567}, "since": "2015-01-01T12:23:00Z", "until": "2015-01-02T12:23:00Z"}, "zone_id": "023e105f4ecef8ad9ca31a8372d0c353"}], "success": true}}}, "schema": {"$ref": "#/components/schemas/zone-analytics-api_user_dashboard_response"}}}}, "4XX": {"description": "Get user analytics dashboard response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zone-analytics-api_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_token": []}], "tags": ["User Analytics (Deprecated)"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPermissionsRequired": {"enum": ["#analytics:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-stainless-deprecation-message": "This endpoint is deprecated and will be removed in a future version. Please use the GraphQL Analytics API instead: https://developers.cloudflare.com/analytics/graphql-api/ It provides equivalent data and more features, including the ability to select only the metrics you need. Migration guide: https://developers.cloudflare.com/analytics/migration-guides/zone-analytics/\n"}
```
