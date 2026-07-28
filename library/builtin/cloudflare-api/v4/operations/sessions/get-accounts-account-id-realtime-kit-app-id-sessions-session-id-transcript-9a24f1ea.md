---
title: Fetch the complete transcript for a session
page_id: operation-get-accounts-account-id-realtime-kit-app-id-sessions-session-id-transcri-8d652ebb
path: operations/sessions
description: Returns a URL to download the transcript for the session ID in CSV format.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/sessions/{session_id}/transcript
operation_ids:
    - GetSessionTranscript
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch the complete transcript for a session

`GET /accounts/{account_id}/realtime/kit/{app_id}/sessions/{session_id}/transcript`

Operation ID: `GetSessionTranscript`

Returns a URL to download the transcript for the session ID in CSV format.

## Definition

```yaml
{"operationId": "GetSessionTranscript", "summary": "Fetch the complete transcript for a session", "description": "Returns a URL to download the transcript for the session ID in CSV format.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "session_id", "in": "path", "description": "ID of the session", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "format", "in": "query", "description": "Transcript file format to fetch.", "schema": {"type": "string", "default": "CSV", "enum": ["SRT", "VTT", "JSON", "CSV"]}}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetSessionTranscript"}}, "security": [{"api_token": []}], "tags": ["Sessions"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
