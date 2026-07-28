---
title: Fetch all recordings for an App
page_id: operation-get-accounts-account-id-realtime-kit-app-id-recordings-9fdd4892
path: operations/recordings
description: Returns all recordings for an App. If the `meeting_id` parameter is passed, returns all recordings for the given meeting ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/recordings
operation_ids:
    - get_all_recordings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch all recordings for an App

`GET /accounts/{account_id}/realtime/kit/{app_id}/recordings`

Operation ID: `get_all_recordings`

Returns all recordings for an App. If the `meeting_id` parameter is passed, returns all recordings for the given meeting ID.

## Definition

```yaml
{"operationId": "get_all_recordings", "summary": "Fetch all recordings for an App", "description": "Returns all recordings for an App. If the `meeting_id` parameter is passed, returns all recordings for the given meeting ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "meeting_id", "in": "query", "description": "ID of a meeting. Optional. Will limit results to only this meeting if passed.", "schema": {"type": "string", "format": "uuid"}}, {"$ref": "#/components/parameters/realtimekit_pageNo"}, {"$ref": "#/components/parameters/realtimekit_perPage"}, {"name": "expired", "in": "query", "description": "If passed, only shows expired/non-expired recordings on RealtimeKit's bucket", "schema": {"type": "boolean"}}, {"$ref": "#/components/parameters/realtimekit_search"}, {"$ref": "#/components/parameters/realtimekit_recordingSortBy"}, {"$ref": "#/components/parameters/realtimekit_sortOrder"}, {"$ref": "#/components/parameters/realtimekit_startTime"}, {"$ref": "#/components/parameters/realtimekit_endTime"}, {"$ref": "#/components/parameters/realtimekit_recordingStatus"}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetRecordings"}}, "security": [{"api_token": []}], "tags": ["Recordings"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
