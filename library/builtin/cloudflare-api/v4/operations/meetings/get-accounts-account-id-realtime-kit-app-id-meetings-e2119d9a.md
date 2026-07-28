---
title: Fetch all meetings for an App
page_id: operation-get-accounts-account-id-realtime-kit-app-id-meetings-0b777ce2
path: operations/meetings
description: Returns all meetings for the given App ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/meetings
operation_ids:
    - get_all_meetings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch all meetings for an App

`GET /accounts/{account_id}/realtime/kit/{app_id}/meetings`

Operation ID: `get_all_meetings`

Returns all meetings for the given App ID.

## Definition

```yaml
{"operationId": "get_all_meetings", "summary": "Fetch all meetings for an App", "description": "Returns all meetings for the given App ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"$ref": "#/components/parameters/realtimekit_pageNo"}, {"$ref": "#/components/parameters/realtimekit_perPage"}, {"$ref": "#/components/parameters/realtimekit_startTime"}, {"$ref": "#/components/parameters/realtimekit_endTime"}, {"$ref": "#/components/parameters/realtimekit_search"}, {"name": "status", "in": "query", "description": "Filter meetings by status.", "schema": {"type": "string", "enum": ["ACTIVE", "INACTIVE"]}}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetAllMeetings"}}, "security": [{"api_token": []}], "tags": ["Meetings"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
