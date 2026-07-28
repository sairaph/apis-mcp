---
title: Fetch summary of transcripts for a session
page_id: operation-get-accounts-account-id-realtime-kit-app-id-sessions-session-id-summary-32d1fc5f
path: operations/sessions
description: Returns a Summary URL to download the Summary of Transcripts for the session ID as plain text.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/sessions/{session_id}/summary
operation_ids:
    - GetSessionSummary
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch summary of transcripts for a session

`GET /accounts/{account_id}/realtime/kit/{app_id}/sessions/{session_id}/summary`

Operation ID: `GetSessionSummary`

Returns a Summary URL to download the Summary of Transcripts for the session ID as plain text.

## Definition

```yaml
{"operationId": "GetSessionSummary", "summary": "Fetch summary of transcripts for a session", "description": "Returns a Summary URL to download the Summary of Transcripts for the session ID as plain text.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "session_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetSessionTranscriptSummary"}}, "security": [{"api_token": []}], "tags": ["Sessions"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
