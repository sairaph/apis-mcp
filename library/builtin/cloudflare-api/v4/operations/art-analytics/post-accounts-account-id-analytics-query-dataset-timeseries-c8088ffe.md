---
title: Query analytics timeseries
page_id: operation-post-accounts-account-id-analytics-query-dataset-timeseries-cac080f0
path: operations/art-analytics
description: Returns time-bucketed analytics data for a dataset. Includes time slots, each containing the requested stats, group-by dimensions, and resolution-controlled bucket size (e.g. `hour`, `day`).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/analytics/query/{dataset}/timeseries
operation_ids:
    - art-analytics-query-timeseries
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Query analytics timeseries

`POST /accounts/{account_id}/analytics/query/{dataset}/timeseries`

Operation ID: `art-analytics-query-timeseries`

Returns time-bucketed analytics data for a dataset. Includes time slots, each containing the requested stats, group-by dimensions, and resolution-controlled bucket size (e.g. `hour`, `day`).

## Definition

```yaml
{"operationId": "art-analytics-query-timeseries", "summary": "Query analytics timeseries", "description": "Returns time-bucketed analytics data for a dataset. Includes time slots, each containing the requested stats, group-by dimensions, and resolution-controlled bucket size (e.g. `hour`, `day`).\n", "parameters": [{"$ref": "#/components/parameters/art_account_id"}, {"name": "dataset", "in": "path", "description": "Dataset name to query. Examples: `access-logins`, `gateway-http`, `gateway-dns`, `gateway-http`, `shadow-it`.\n", "required": true, "schema": {"type": "string", "example": "shadow_it"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"access_logins_daily": {"summary": "Access logins daily timeseries with filter", "value": {"filters": [{"name": "allowed", "op": "eq", "values": [true]}], "from": "2024-11-01T00:00:00Z", "groupBy": ["country", "allowed"], "resolution": "day", "stats": ["attemptsTotal"], "to": "2024-11-08T00:00:00Z"}}, "shadow_it_hourly": {"summary": "Shadow IT hourly timeseries", "value": {"filters": [], "from": "2024-11-05T00:00:00Z", "groupBy": ["appName"], "resolution": "hour", "stats": ["bytesTotal"], "to": "2024-11-06T00:00:00Z"}}}, "schema": {"$ref": "#/components/schemas/art_QueryTimeseries"}}}}, "responses": {"200": {"description": "Timeseries query result.", "content": {"application/json": {"examples": {"success": {"summary": "Successful timeseries result", "value": {"errors": [], "messages": [{"code": 1000, "message": "API in beta: expect breaking changes."}], "result": {"resolution": "hour", "slots": [{"appName": "Slack", "bytesTotal": 1048576, "time_bucket": "2024-11-05T00:00:00Z"}, {"appName": "Slack", "bytesTotal": 2097152, "time_bucket": "2024-11-05T01:00:00Z"}]}, "success": true}}}, "schema": {"$ref": "#/components/schemas/art_TimeseriesResponse"}}}}, "400": {"$ref": "#/components/responses/art_BadRequest"}, "401": {"$ref": "#/components/responses/art_Unauthorized"}, "403": {"$ref": "#/components/responses/art_Forbidden"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["ART Analytics"], "x-api-token-group": ["Zero Trust Read"]}
```
