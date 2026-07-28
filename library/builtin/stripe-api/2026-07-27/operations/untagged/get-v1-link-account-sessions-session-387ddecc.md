---
title: Retrieve a Session
page_id: operation-get-v1-link-account-sessions-session-e82d665b
path: operations/untagged
description: <p>Retrieves the details of a Financial Connections <code>Session</code></p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/link_account_sessions/{session}
operation_ids:
    - GetLinkAccountSessionsSession
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a Session

`GET /v1/link_account_sessions/{session}`

Operation ID: `GetLinkAccountSessionsSession`

<p>Retrieves the details of a Financial Connections <code>Session</code></p>

## Definition

```yaml
{"summary": "Retrieve a Session", "description": "<p>Retrieves the details of a Financial Connections <code>Session</code></p>", "operationId": "GetLinkAccountSessionsSession", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "session", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/financial_connections.session"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
