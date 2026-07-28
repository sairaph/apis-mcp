---
title: Fetch complete analytics data for your livestreams
page_id: operation-get-accounts-account-id-realtime-kit-app-id-analytics-livestreams-overal-b8e51eb8
path: operations/live-streams
description: Returns livestream analytics for the specified time range.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/analytics/livestreams/overall
operation_ids:
    - get-livestream-analytics-complete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch complete analytics data for your livestreams

`GET /accounts/{account_id}/realtime/kit/{app_id}/analytics/livestreams/overall`

Operation ID: `get-livestream-analytics-complete`

Returns livestream analytics for the specified time range.

## Definition

```yaml
{"operationId": "get-livestream-analytics-complete", "summary": "Fetch complete analytics data for your livestreams", "description": "Returns livestream analytics for the specified time range.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "start_time", "in": "query", "description": "Specify the start time as a Unix timestamp in seconds to access the livestream analytics.", "schema": {"type": "integer", "format": "int64"}}, {"name": "end_time", "in": "query", "description": "Specify the end time as a Unix timestamp in seconds to access the livestream analytics.", "schema": {"type": "integer", "format": "int64"}}, {"name": "filters", "in": "query", "description": "Optional filters for livestream analytics.", "schema": {"type": "string"}}], "responses": {"200": {"description": "OK", "content": {"application/json": {"examples": {"Example 1": {"value": {"data": {"count": 4, "total_ingest_seconds": 531, "total_viewer_seconds": 116}, "success": true}}}, "schema": {"type": "object", "properties": {"data": {"type": "object", "properties": {"count": {"description": "Count of total livestreams.", "type": "integer"}, "total_ingest_seconds": {"description": "Total time duration for which the input was given or the meeting was streamed.", "type": "integer"}, "total_viewer_seconds": {"description": "Total view time for which the viewers watched the stream.", "type": "integer"}}}, "success": {"type": "boolean"}}}}}}}, "security": [{"api_token": []}], "tags": ["Live streams", "LivestreamAnalytics"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
