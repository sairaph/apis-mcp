---
title: Fetch details of an active session
page_id: operation-get-accounts-account-id-realtime-kit-app-id-meetings-meeting-id-active-s-e7e69ff8
path: operations/active-session
description: Returns details of an ongoing active session for the given meeting ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/active-session
operation_ids:
    - GetActiveSession
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch details of an active session

`GET /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/active-session`

Operation ID: `GetActiveSession`

Returns details of an ongoing active session for the given meeting ID.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"$ref": "#/components/parameters/realtimekit_meeting_id"}]
```

## Definition

```yaml
{"operationId": "GetActiveSession", "summary": "Fetch details of an active session", "description": "Returns details of an ongoing active session for the given meeting ID.", "parameters": [{"$ref": "#/components/parameters/realtimekit_meetingId"}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetActiveSession"}, "404": {"$ref": "#/components/responses/realtimekit_GetActiveSessionNotFound"}}, "security": [{"api_token": []}], "tags": ["Active session"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
