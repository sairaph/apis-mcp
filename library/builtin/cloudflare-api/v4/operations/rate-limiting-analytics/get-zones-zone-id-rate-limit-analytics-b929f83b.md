---
title: Get Rate Limiting Analytics
page_id: operation-get-zones-zone-id-rate-limit-analytics-f64fd240
path: operations/rate-limiting-analytics
description: |-
    Returns rate limiting analytics for a zone over the specified time period.
    The time period divides into time segments of a given length. Each segment
    contains total action counts and action counts broken down by colo.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/rate_limit_analytics
operation_ids:
    - rate-limit-analytics-get-zone-analytics
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Rate Limiting Analytics

`GET /zones/{zone_id}/rate_limit_analytics`

Operation ID: `rate-limit-analytics-get-zone-analytics`

Returns rate limiting analytics for a zone over the specified time period.
The time period divides into time segments of a given length. Each segment
contains total action counts and action counts broken down by colo.

## Definition

```yaml
{"operationId": "rate-limit-analytics-get-zone-analytics", "summary": "Get Rate Limiting Analytics", "description": "Returns rate limiting analytics for a zone over the specified time period.\nThe time period divides into time segments of a given length. Each segment\ncontains total action counts and action counts broken down by colo.", "parameters": [{"name": "zone_id", "in": "path", "description": "Identifier of the zone.", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "pattern": "^[0-9a-f]{32}$"}}, {"name": "since", "in": "query", "description": "The start of the queried time period. Time must be rounded to the time\nsegment boundary and formatted as RFC 3339.", "required": true, "schema": {"type": "string", "format": "date-time", "example": "2024-01-01T00:00:00Z"}}, {"name": "until", "in": "query", "description": "The exclusive end of the queried time period. Time must be rounded to the\ntime segment boundary and formatted as RFC 3339.", "required": true, "schema": {"type": "string", "format": "date-time", "example": "2024-01-02T00:00:00Z"}}, {"name": "time_delta", "in": "query", "description": "Length (in seconds) of each time segment dividing the entire time period.\nAccepted values are 60 (minute), 3600 (hour), 86400 (day), and 2592000 (month).", "required": true, "schema": {"type": "integer", "example": 3600, "enum": [60, 3600, 86400, 2592000]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rlanalytics_rate_limit_analytics"}}}}, "400": {"description": "Bad request. Invalid or missing parameters.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rlanalytics_api-response-common-failure"}}}}, "401": {"description": "Unauthorized. Missing or invalid authentication credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rlanalytics_api-response-common-failure"}}}}, "403": {"description": "Forbidden. Insufficient permissions for the requested resource.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rlanalytics_api-response-common-failure"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rlanalytics_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Rate Limiting Analytics"], "x-api-token-group": ["Analytics Read"]}
```
