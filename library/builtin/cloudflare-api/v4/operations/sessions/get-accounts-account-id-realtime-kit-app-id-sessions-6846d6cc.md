---
title: Fetch all sessions of an App
page_id: operation-get-accounts-account-id-realtime-kit-app-id-sessions-1b2c93ba
path: operations/sessions
description: Returns details of all sessions of an App.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/sessions
operation_ids:
    - GetSessions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch all sessions of an App

`GET /accounts/{account_id}/realtime/kit/{app_id}/sessions`

Operation ID: `GetSessions`

Returns details of all sessions of an App.

## Definition

```yaml
{"operationId": "GetSessions", "summary": "Fetch all sessions of an App", "description": "Returns details of all sessions of an App.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"$ref": "#/components/parameters/realtimekit_pageNo"}, {"$ref": "#/components/parameters/realtimekit_perPage"}, {"$ref": "#/components/parameters/realtimekit_sortBy"}, {"$ref": "#/components/parameters/realtimekit_sortOrder"}, {"$ref": "#/components/parameters/realtimekit_startTime"}, {"$ref": "#/components/parameters/realtimekit_endTime"}, {"$ref": "#/components/parameters/realtimekit_participants"}, {"$ref": "#/components/parameters/realtimekit_status"}, {"name": "search", "in": "query", "description": "Search string that matches sessions based on meeting title, meeting ID, and session ID", "schema": {"type": "string"}}, {"name": "associated_id", "in": "query", "description": "ID of the meeting that sessions should be associated with", "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetSessions"}}, "security": [{"api_token": []}], "tags": ["Sessions"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
