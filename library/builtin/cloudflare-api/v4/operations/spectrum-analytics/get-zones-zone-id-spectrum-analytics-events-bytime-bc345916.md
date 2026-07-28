---
title: Get analytics by time
page_id: operation-get-zones-zone-id-spectrum-analytics-events-bytime-7b3393cd
path: operations/spectrum-analytics
description: Retrieves a list of aggregate metrics grouped by time interval.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/spectrum/analytics/events/bytime
operation_ids:
    - spectrum-analytics-(-by-time)-get-analytics-by-time
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get analytics by time

`GET /zones/{zone_id}/spectrum/analytics/events/bytime`

Operation ID: `spectrum-analytics-(-by-time)-get-analytics-by-time`

Retrieves a list of aggregate metrics grouped by time interval.

## Definition

```yaml
{"operationId": "spectrum-analytics-(-by-time)-get-analytics-by-time", "summary": "Get analytics by time", "description": "Retrieves a list of aggregate metrics grouped by time interval.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/spectrum-analytics_identifier"}}, {"name": "dimensions", "in": "query", "schema": {"$ref": "#/components/schemas/spectrum-analytics_dimensions"}}, {"name": "sort", "in": "query", "schema": {"$ref": "#/components/schemas/spectrum-analytics_sort"}}, {"name": "until", "in": "query", "schema": {"$ref": "#/components/schemas/spectrum-analytics_until"}}, {"name": "metrics", "in": "query", "schema": {"$ref": "#/components/schemas/spectrum-analytics_metrics"}}, {"name": "filters", "in": "query", "schema": {"$ref": "#/components/schemas/spectrum-analytics_filters"}}, {"name": "since", "in": "query", "schema": {"$ref": "#/components/schemas/spectrum-analytics_since"}}, {"name": "time_delta", "in": "query", "required": true, "schema": {"description": "Used to select time series resolution.", "type": "string", "example": "minute", "enum": ["year", "quarter", "month", "week", "day", "hour", "dekaminute", "minute"]}}], "responses": {"200": {"description": "Get analytics by time response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-analytics_query-response-single"}}}}, "4xx": {"description": "Get analytics by time response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-analytics_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Spectrum Analytics"], "x-api-token-group": ["Analytics Read"]}
```
