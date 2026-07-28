---
title: Fetch details of a recording
page_id: operation-get-accounts-account-id-realtime-kit-app-id-recordings-recording-id-54aa9c51
path: operations/recordings
description: Returns details of a recording for the given recording ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/recordings/{recording_id}
operation_ids:
    - get_one_recording
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch details of a recording

`GET /accounts/{account_id}/realtime/kit/{app_id}/recordings/{recording_id}`

Operation ID: `get_one_recording`

Returns details of a recording for the given recording ID.

## Definition

```yaml
{"operationId": "get_one_recording", "summary": "Fetch details of a recording", "description": "Returns details of a recording for the given recording ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "recording_id", "in": "path", "description": "ID of the recording", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetRecording"}}, "security": [{"api_token": []}], "tags": ["Recordings"], "x-api-token-group": ["Realtime Admin", "Realtime"], "x-stability": "beta"}
```
