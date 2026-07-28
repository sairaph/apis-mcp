---
title: Refresh participant's authentication token
page_id: operation-post-accounts-account-id-realtime-kit-app-id-meetings-meeting-id-partici-78d416a3
path: operations/meetings
description: Regenerates participant's authentication token for the given meeting and participant ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/participants/{participant_id}/token
operation_ids:
    - regenerate_token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Refresh participant's authentication token

`POST /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/participants/{participant_id}/token`

Operation ID: `regenerate_token`

Regenerates participant's authentication token for the given meeting and participant ID.

## Definition

```yaml
{"operationId": "regenerate_token", "summary": "Refresh participant's authentication token", "description": "Regenerates participant's authentication token for the given meeting and participant ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "meeting_id", "in": "path", "description": "ID of the meeting. You can fetch the meeting ID using the create a meeting API.", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "participant_id", "in": "path", "description": "ID of the participant. You can fetch the participant ID using the add a  participant API.", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_RegenToken"}, "500": {"$ref": "#/components/responses/realtimekit_GenericError"}}, "security": [{"api_token": []}], "tags": ["Meetings"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
