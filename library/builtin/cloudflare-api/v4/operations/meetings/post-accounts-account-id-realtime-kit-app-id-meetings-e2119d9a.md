---
title: Create a meeting
page_id: operation-post-accounts-account-id-realtime-kit-app-id-meetings-a558862f
path: operations/meetings
description: Create a meeting for the given App ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/meetings
operation_ids:
    - create_meeting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a meeting

`POST /accounts/{account_id}/realtime/kit/{app_id}/meetings`

Operation ID: `create_meeting`

Create a meeting for the given App ID.

## Definition

```yaml
{"operationId": "create_meeting", "summary": "Create a meeting", "description": "Create a meeting for the given App ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}], "requestBody": {"$ref": "#/components/requestBodies/realtimekit_CreateMeetingBody"}, "responses": {"201": {"$ref": "#/components/responses/realtimekit_GetMeeting"}}, "security": [{"api_token": []}], "tags": ["Meetings"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
