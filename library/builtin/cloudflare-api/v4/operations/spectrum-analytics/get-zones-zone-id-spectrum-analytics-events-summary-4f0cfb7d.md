---
title: Get analytics summary
page_id: operation-get-zones-zone-id-spectrum-analytics-events-summary-c1802762
path: operations/spectrum-analytics
description: Retrieves a list of summarised aggregate metrics over a given time period.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/spectrum/analytics/events/summary
operation_ids:
    - spectrum-analytics-(-summary)-get-analytics-summary
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get analytics summary

`GET /zones/{zone_id}/spectrum/analytics/events/summary`

Operation ID: `spectrum-analytics-(-summary)-get-analytics-summary`

Retrieves a list of summarised aggregate metrics over a given time period.

## Definition

```yaml
{"operationId": "spectrum-analytics-(-summary)-get-analytics-summary", "summary": "Get analytics summary", "description": "Retrieves a list of summarised aggregate metrics over a given time period.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/spectrum-analytics_identifier"}}, {"name": "dimensions", "in": "query", "schema": {"$ref": "#/components/schemas/spectrum-analytics_dimensions"}}, {"name": "sort", "in": "query", "schema": {"$ref": "#/components/schemas/spectrum-analytics_sort"}}, {"name": "until", "in": "query", "schema": {"$ref": "#/components/schemas/spectrum-analytics_until"}}, {"name": "metrics", "in": "query", "schema": {"$ref": "#/components/schemas/spectrum-analytics_metrics"}}, {"name": "filters", "in": "query", "schema": {"$ref": "#/components/schemas/spectrum-analytics_filters"}}, {"name": "since", "in": "query", "schema": {"$ref": "#/components/schemas/spectrum-analytics_since"}}], "responses": {"200": {"description": "Get analytics summary response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-analytics_query-response-single"}}}}, "4xx": {"description": "Get analytics summary response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-analytics_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Spectrum Analytics"], "x-api-token-group": ["Analytics Read"]}
```
