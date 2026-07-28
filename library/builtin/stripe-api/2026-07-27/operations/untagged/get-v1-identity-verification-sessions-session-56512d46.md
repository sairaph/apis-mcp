---
title: Retrieve a VerificationSession
page_id: operation-get-v1-identity-verification-sessions-session-eed18a32
path: operations/untagged
description: |-
    <p>Retrieves the details of a VerificationSession that was previously created.</p>

    <p>When the session status is <code>requires_input</code>, you can use this method to retrieve a valid
    <code>client_secret</code> or <code>url</code> to allow re-submission.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/identity/verification_sessions/{session}
operation_ids:
    - GetIdentityVerificationSessionsSession
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a VerificationSession

`GET /v1/identity/verification_sessions/{session}`

Operation ID: `GetIdentityVerificationSessionsSession`

<p>Retrieves the details of a VerificationSession that was previously created.</p>

<p>When the session status is <code>requires_input</code>, you can use this method to retrieve a valid
<code>client_secret</code> or <code>url</code> to allow re-submission.</p>

## Definition

```yaml
{"summary": "Retrieve a VerificationSession", "description": "<p>Retrieves the details of a VerificationSession that was previously created.</p>\n\n<p>When the session status is <code>requires_input</code>, you can use this method to retrieve a valid\n<code>client_secret</code> or <code>url</code> to allow re-submission.</p>", "operationId": "GetIdentityVerificationSessionsSession", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "session", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/identity.verification_session"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
