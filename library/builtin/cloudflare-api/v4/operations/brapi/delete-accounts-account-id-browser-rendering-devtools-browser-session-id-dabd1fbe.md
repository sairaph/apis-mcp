---
title: Close browser session.
page_id: operation-delete-accounts-account-id-browser-rendering-devtools-browser-session-id-36cb3389
path: operations/brapi
description: Closes an existing browser session.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/browser-rendering/devtools/browser/{session_id}
operation_ids:
    - brapi-delete_DevtoolsBrowserDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Close browser session.

`DELETE /accounts/{account_id}/browser-rendering/devtools/browser/{session_id}`

Operation ID: `brapi-delete_DevtoolsBrowserDelete`

Closes an existing browser session.

## Definition

```yaml
{"operationId": "brapi-delete_DevtoolsBrowserDelete", "summary": "Close browser session.", "description": "Closes an existing browser session.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "session_id", "in": "path", "description": "Browser session ID to close.", "required": true, "schema": {"description": "Browser session ID to close.", "type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Browser session closing or already closed.", "content": {"application/json": {"schema": {"type": "object", "properties": {"status": {"type": "string", "enum": ["closing", "closed"]}}, "required": ["status"]}}}}, "404": {"description": "Browser session not found."}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["brapi"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.brapi.read"], "type": "string"}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "browser-run.devtools.browser", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
