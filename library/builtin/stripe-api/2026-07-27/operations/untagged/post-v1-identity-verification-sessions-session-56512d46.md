---
title: Update a VerificationSession
page_id: operation-post-v1-identity-verification-sessions-session-fa64509d
path: operations/untagged
description: |-
    <p>Updates a VerificationSession object.</p>

    <p>When the session status is <code>requires_input</code>, you can use this method to update the
    verification check and options.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/identity/verification_sessions/{session}
operation_ids:
    - PostIdentityVerificationSessionsSession
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a VerificationSession

`POST /v1/identity/verification_sessions/{session}`

Operation ID: `PostIdentityVerificationSessionsSession`

<p>Updates a VerificationSession object.</p>

<p>When the session status is <code>requires_input</code>, you can use this method to update the
verification check and options.</p>

## Definition

```yaml
{"summary": "Update a VerificationSession", "description": "<p>Updates a VerificationSession object.</p>\n\n<p>When the session status is <code>requires_input</code>, you can use this method to update the\nverification check and options.</p>", "operationId": "PostIdentityVerificationSessionsSession", "parameters": [{"name": "session", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "options": {"title": "session_options_param", "type": "object", "properties": {"document": {"anyOf": [{"title": "document_options", "type": "object", "properties": {"allowed_types": {"type": "array", "items": {"type": "string", "enum": ["driving_license", "id_card", "passport"]}}, "require_id_number": {"type": "boolean"}, "require_live_capture": {"type": "boolean"}, "require_matching_selfie": {"type": "boolean"}}}, {"type": "string", "enum": [""]}]}}, "description": "A set of options for the session’s verification checks."}, "provided_details": {"title": "provided_details_param", "type": "object", "properties": {"email": {"type": "string"}, "phone": {"type": "string"}}, "description": "Details provided about the user being verified. These details may be shown to the user."}, "type": {"type": "string", "description": "The type of [verification check](https://docs.stripe.com/identity/verification-checks) to be performed.", "enum": ["document", "id_number"], "x-stripeBypassValidation": true}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "options": {"style": "deepObject", "explode": true}, "provided_details": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/identity.verification_session"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
