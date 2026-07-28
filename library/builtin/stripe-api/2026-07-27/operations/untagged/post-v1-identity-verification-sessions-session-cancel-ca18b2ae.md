---
title: Cancel a VerificationSession
page_id: operation-post-v1-identity-verification-sessions-session-cancel-20915369
path: operations/untagged
description: |-
    <p>A VerificationSession object can be canceled when it is in <code>requires_input</code> <a href="/docs/identity/how-sessions-work">status</a>.</p>

    <p>Once canceled, future submission attempts are disabled. This cannot be undone. <a href="/docs/identity/verification-sessions#cancel">Learn more</a>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/identity/verification_sessions/{session}/cancel
operation_ids:
    - PostIdentityVerificationSessionsSessionCancel
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Cancel a VerificationSession

`POST /v1/identity/verification_sessions/{session}/cancel`

Operation ID: `PostIdentityVerificationSessionsSessionCancel`

<p>A VerificationSession object can be canceled when it is in <code>requires_input</code> <a href="/docs/identity/how-sessions-work">status</a>.</p>

<p>Once canceled, future submission attempts are disabled. This cannot be undone. <a href="/docs/identity/verification-sessions#cancel">Learn more</a>.</p>

## Definition

```yaml
{"summary": "Cancel a VerificationSession", "description": "<p>A VerificationSession object can be canceled when it is in <code>requires_input</code> <a href=\"/docs/identity/how-sessions-work\">status</a>.</p>\n\n<p>Once canceled, future submission attempts are disabled. This cannot be undone. <a href=\"/docs/identity/verification-sessions#cancel\">Learn more</a>.</p>", "operationId": "PostIdentityVerificationSessionsSessionCancel", "parameters": [{"name": "session", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/identity.verification_session"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
