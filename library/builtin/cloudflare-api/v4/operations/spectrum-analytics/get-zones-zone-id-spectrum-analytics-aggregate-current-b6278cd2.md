---
title: Get current aggregated analytics
page_id: operation-get-zones-zone-id-spectrum-analytics-aggregate-current-5b099607
path: operations/spectrum-analytics
description: Retrieves analytics aggregated from the last minute of usage on Spectrum applications underneath a given zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/spectrum/analytics/aggregate/current
operation_ids:
    - spectrum-aggregate-analytics-get-current-aggregated-analytics
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get current aggregated analytics

`GET /zones/{zone_id}/spectrum/analytics/aggregate/current`

Operation ID: `spectrum-aggregate-analytics-get-current-aggregated-analytics`

Retrieves analytics aggregated from the last minute of usage on Spectrum applications underneath a given zone.

## Definition

```yaml
{"operationId": "spectrum-aggregate-analytics-get-current-aggregated-analytics", "summary": "Get current aggregated analytics", "description": "Retrieves analytics aggregated from the last minute of usage on Spectrum applications underneath a given zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/spectrum-analytics_identifier"}}, {"name": "appID", "in": "query", "schema": {"$ref": "#/components/schemas/spectrum-analytics_app_id_param"}}, {"name": "colo_name", "in": "query", "schema": {"description": "Co-location identifier.", "type": "string", "example": "PDX", "maxLength": 3}}], "responses": {"200": {"description": "Get current aggregated analytics response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-analytics_query-response-aggregate"}}}}, "4xx": {"description": "Get current aggregated analytics response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-analytics_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Spectrum Analytics"], "x-api-token-group": ["Analytics Read"]}
```
