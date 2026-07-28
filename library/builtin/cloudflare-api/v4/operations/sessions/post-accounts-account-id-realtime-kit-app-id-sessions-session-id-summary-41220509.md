---
title: Generate summary of Transcripts for the session
page_id: operation-post-accounts-account-id-realtime-kit-app-id-sessions-session-id-summary-31616329
path: operations/sessions
description: Trigger Summary generation of Transcripts for the session ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/sessions/{session_id}/summary
operation_ids:
    - post-sessions-session_id-summary
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Generate summary of Transcripts for the session

`POST /accounts/{account_id}/realtime/kit/{app_id}/sessions/{session_id}/summary`

Operation ID: `post-sessions-session_id-summary`

Trigger Summary generation of Transcripts for the session ID.

## Definition

```yaml
{"operationId": "post-sessions-session_id-summary", "summary": "Generate summary of Transcripts for the session", "description": "Trigger Summary generation of Transcripts for the session ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "session_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Success", "content": {"application/json": {"schema": {"type": "object", "properties": {"data": {"type": "object", "properties": {"session_id": {"type": "string", "format": "uuid"}, "status": {"type": "string"}}}, "success": {"type": "boolean"}}}}}}}, "security": [{"api_token": []}], "tags": ["Sessions"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
