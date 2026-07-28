---
title: Get dashboard
page_id: operation-get-zones-zone-identifier-analytics-dashboard-de7ed08e
path: operations/zone-analytics-deprecated
description: The dashboard view provides both totals and timeseries data for the given zone and time period across the entire Cloudflare network.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_identifier}/analytics/dashboard
operation_ids:
    - zone-analytics-(-deprecated)-get-dashboard
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get dashboard

`GET /zones/{zone_identifier}/analytics/dashboard`

Operation ID: `zone-analytics-(-deprecated)-get-dashboard`

The dashboard view provides both totals and timeseries data for the given zone and time period across the entire Cloudflare network.

## Definition

```yaml
{"operationId": "zone-analytics-(-deprecated)-get-dashboard", "summary": "Get dashboard", "description": "The dashboard view provides both totals and timeseries data for the given zone and time period across the entire Cloudflare network.", "parameters": [{"name": "zone_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zone-analytics-api_identifier"}}, {"name": "until", "in": "query", "schema": {"$ref": "#/components/schemas/zone-analytics-api_until"}}, {"name": "since", "in": "query", "schema": {"description": "The (inclusive) beginning of the requested time frame. This value can be a negative integer representing the number of minutes in the past relative to time the request is made, or can be an absolute timestamp that conforms to RFC 3339. At this point in time, it cannot exceed a time in the past greater than one year.\n\nRanges that the Cloudflare web application provides will provide the following period length for each point:\n- Last 60 minutes (from -59 to -1): 1 minute resolution\n- Last 7 hours (from -419 to -60): 15 minutes resolution\n- Last 15 hours (from -899 to -420): 30 minutes resolution\n- Last 72 hours (from -4320 to -900): 1 hour resolution\n- Older than 3 days (-525600 to -4320): 1 day resolution.", "example": "2015-01-01T12:23:00Z", "default": -10080, "anyOf": [{"type": "string"}, {"type": "integer"}]}}, {"name": "continuous", "in": "query", "schema": {"description": "When set to true, the API will move the requested time window backward, until it finds a region with completely aggregated data.\n\nThe API response _may not represent the requested time window_.", "type": "boolean", "default": true}}], "responses": {"200": {"description": "Get dashboard response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zone-analytics-api_dashboard_response"}}}}, "4XX": {"description": "Get dashboard response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zone-analytics-api_dashboard_response"}, {"$ref": "#/components/schemas/zone-analytics-api_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zone Analytics (Deprecated)"], "x-api-token-group": ["Analytics Read"], "x-cfDeprecation": {"description": "Please use the new GraphQL Analytics API instead: https://developers.cloudflare.com/analytics/graphql-api/. It provides equivalent data and more features, including the ability to select only the metrics you need. Migration guide: https://developers.cloudflare.com/analytics/migration-guides/zone-analytics/.", "display": true, "eol": "2021-03-01", "id": "zone_analytics_deprecation"}, "x-cfPermissionsRequired": {"enum": ["#analytics:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
