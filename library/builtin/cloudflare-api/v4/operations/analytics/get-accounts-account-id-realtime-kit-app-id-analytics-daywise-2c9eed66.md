---
title: Fetch day-wise session and recording analytics data for an App
page_id: operation-get-accounts-account-id-realtime-kit-app-id-analytics-daywise-c0994b7a
path: operations/analytics
description: Returns day-wise session and recording analytics data of an App for the specified time range start_date to end_date. If start_date and end_date are not provided, the default time range is set from 30 days ago to the current date.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/analytics/daywise
operation_ids:
    - get-org-analytics
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch day-wise session and recording analytics data for an App

`GET /accounts/{account_id}/realtime/kit/{app_id}/analytics/daywise`

Operation ID: `get-org-analytics`

Returns day-wise session and recording analytics data of an App for the specified time range start_date to end_date. If start_date and end_date are not provided, the default time range is set from 30 days ago to the current date.

## Definition

```yaml
{"operationId": "get-org-analytics", "summary": "Fetch day-wise session and recording analytics data for an App", "description": "Returns day-wise session and recording analytics data of an App for the specified time range start_date to end_date. If start_date and end_date are not provided, the default time range is set from 30 days ago to the current date.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"$ref": "#/components/parameters/realtimekit_startDate"}, {"$ref": "#/components/parameters/realtimekit_endDate"}], "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"type": "object", "properties": {"data": {"type": "object", "properties": {"recording_stats": {"description": "Recording statistics of an App during the range specified", "type": "object", "properties": {"day_stats": {"description": "Day wise recording stats", "type": "array", "items": {"properties": {"day": {"type": "string"}, "total_recording_minutes": {"description": "Total recording minutes for a specific day", "type": "integer"}, "total_recordings": {"description": "Total number of recordings for a specific day", "type": "integer"}}, "type": "object"}}, "recording_count": {"description": "Total number of recordings during the range specified", "type": "integer"}, "recording_minutes_consumed": {"description": "Total recording minutes during the range specified", "type": "number"}}}, "session_stats": {"description": "Session statistics of an App during the range specified", "type": "object", "properties": {"day_stats": {"description": "Day wise session stats", "type": "array", "items": {"properties": {"day": {"type": "string"}, "total_session_minutes": {"description": "Total session minutes for a specific day", "type": "number"}, "total_sessions": {"description": "Total number of sessions for a specific day", "type": "integer"}}, "type": "object"}}, "sessions_count": {"description": "Total number of sessions during the range specified", "type": "integer"}, "sessions_minutes_consumed": {"description": "Total session minutes during the range specified", "type": "number"}}}}}, "success": {"type": "boolean"}}}}}}}, "security": [{"api_token": []}], "tags": ["Analytics", "Organizations"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
