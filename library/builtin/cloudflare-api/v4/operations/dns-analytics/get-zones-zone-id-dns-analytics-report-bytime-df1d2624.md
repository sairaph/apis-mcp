---
title: By Time
page_id: operation-get-zones-zone-id-dns-analytics-report-bytime-42f3cbd1
path: operations/dns-analytics
description: |-
    Retrieves a list of aggregate metrics grouped by time interval.

    See [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/) for detailed information about the available query parameters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/dns_analytics/report/bytime
operation_ids:
    - dns-analytics-by-time
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# By Time

`GET /zones/{zone_id}/dns_analytics/report/bytime`

Operation ID: `dns-analytics-by-time`

Retrieves a list of aggregate metrics grouped by time interval.

See [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/) for detailed information about the available query parameters.

## Definition

```yaml
{"operationId": "dns-analytics-by-time", "summary": "By Time", "description": "Retrieves a list of aggregate metrics grouped by time interval.\n\nSee [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/) for detailed information about the available query parameters.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-analytics_identifier"}}, {"name": "metrics", "in": "query", "schema": {"$ref": "#/components/schemas/dns-analytics_metrics"}}, {"name": "dimensions", "in": "query", "schema": {"$ref": "#/components/schemas/dns-analytics_dimensions"}}, {"name": "since", "in": "query", "schema": {"$ref": "#/components/schemas/dns-analytics_since"}}, {"name": "until", "in": "query", "schema": {"$ref": "#/components/schemas/dns-analytics_until"}}, {"name": "limit", "in": "query", "schema": {"$ref": "#/components/schemas/dns-analytics_limit"}}, {"name": "sort", "in": "query", "schema": {"$ref": "#/components/schemas/dns-analytics_sort"}}, {"name": "filters", "in": "query", "schema": {"$ref": "#/components/schemas/dns-analytics_filters"}}, {"name": "time_delta", "in": "query", "schema": {"$ref": "#/components/schemas/dns-analytics_time_delta"}}], "responses": {"200": {"description": "By Time response", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-analytics_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dns-analytics_report_bytime"}}, "type": "object"}]}}}}, "4XX": {"description": "By Time response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-analytics_api-response-common-failure"}, {"properties": {"result": {"$ref": "#/components/schemas/dns-analytics_report_bytime"}}, "type": "object"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Analytics"], "x-api-token-group": ["Analytics Read"], "x-cfPermissionsRequired": {"enum": ["#analytics:read"]}}
```
