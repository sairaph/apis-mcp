---
title: Pause/Resume/Stop recording
page_id: operation-put-accounts-account-id-realtime-kit-app-id-recordings-recording-id-8ac0f5ec
path: operations/recordings
description: Pause/Resume/Stop a given recording ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/recordings/{recording_id}
operation_ids:
    - pause_resume_stop_recording
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Pause/Resume/Stop recording

`PUT /accounts/{account_id}/realtime/kit/{app_id}/recordings/{recording_id}`

Operation ID: `pause_resume_stop_recording`

Pause/Resume/Stop a given recording ID.

## Definition

```yaml
{"operationId": "pause_resume_stop_recording", "summary": "Pause/Resume/Stop recording", "description": "Pause/Resume/Stop a given recording ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "description": "A Cloudflare-generated unique identifier for an item.", "required": true, "schema": {"type": "string", "maxLength": 32, "minLength": 32, "pattern": "^[a-f0-9]{32}$"}, "example": "2a95132c15732412d22c1476fa83f27a"}, {"name": "recording_id", "in": "path", "description": "ID of the recording", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"action": {"type": "string", "enum": ["stop", "pause", "resume"]}}, "required": ["action"]}}}}, "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetRecording"}}, "security": [{"api_token": []}], "tags": ["Recordings"], "x-api-token-group": ["Realtime Admin", "Realtime"], "x-stability": "beta"}
```
