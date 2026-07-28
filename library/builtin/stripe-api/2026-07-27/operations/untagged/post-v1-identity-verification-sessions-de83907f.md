---
title: Create a VerificationSession
page_id: operation-post-v1-identity-verification-sessions-5cd2068d
path: operations/untagged
description: |-
    <p>Creates a VerificationSession object.</p>

    <p>After the VerificationSession is created, display a verification modal using the session <code>client_secret</code> or send your users to the session’s <code>url</code>.</p>

    <p>If your API key is in test mode, verification checks won’t actually process, though everything else will occur as if in live mode.</p>

    <p>Related guide: <a href="/docs/identity/verify-identity-documents">Verify your users’ identity documents</a></p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/identity/verification_sessions
operation_ids:
    - PostIdentityVerificationSessions
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a VerificationSession

`POST /v1/identity/verification_sessions`

Operation ID: `PostIdentityVerificationSessions`

<p>Creates a VerificationSession object.</p>

<p>After the VerificationSession is created, display a verification modal using the session <code>client_secret</code> or send your users to the session’s <code>url</code>.</p>

<p>If your API key is in test mode, verification checks won’t actually process, though everything else will occur as if in live mode.</p>

<p>Related guide: <a href="/docs/identity/verify-identity-documents">Verify your users’ identity documents</a></p>

## Definition

```yaml
{"summary": "Create a VerificationSession", "description": "<p>Creates a VerificationSession object.</p>\n\n<p>After the VerificationSession is created, display a verification modal using the session <code>client_secret</code> or send your users to the session’s <code>url</code>.</p>\n\n<p>If your API key is in test mode, verification checks won’t actually process, though everything else will occur as if in live mode.</p>\n\n<p>Related guide: <a href=\"/docs/identity/verify-identity-documents\">Verify your users’ identity documents</a></p>", "operationId": "PostIdentityVerificationSessions", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"client_reference_id": {"maxLength": 5000, "type": "string", "description": "A string to reference this user. This can be a customer ID, a session ID, or similar, and can be used to reconcile this verification with your internal systems."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "options": {"title": "session_options_param", "type": "object", "properties": {"document": {"anyOf": [{"title": "document_options", "type": "object", "properties": {"allowed_types": {"type": "array", "items": {"type": "string", "enum": ["driving_license", "id_card", "passport"]}}, "require_id_number": {"type": "boolean"}, "require_live_capture": {"type": "boolean"}, "require_matching_selfie": {"type": "boolean"}}}, {"type": "string", "enum": [""]}]}}, "description": "A set of options for the session’s verification checks."}, "provided_details": {"title": "provided_details_param", "type": "object", "properties": {"email": {"type": "string"}, "phone": {"type": "string"}}, "description": "Details provided about the user being verified. These details might be shown to the user."}, "related_customer": {"maxLength": 5000, "type": "string", "description": "Customer ID"}, "related_customer_account": {"maxLength": 5000, "type": "string", "description": "The ID of the Account representing a customer."}, "related_person": {"title": "related_person_param", "required": ["account", "person"], "type": "object", "properties": {"account": {"maxLength": 5000, "type": "string"}, "person": {"maxLength": 5000, "type": "string"}}, "description": "Tokens referencing a Person resource and its associated account."}, "return_url": {"type": "string", "description": "The URL that the user will be redirected to upon completing the verification flow."}, "type": {"type": "string", "description": "The type of [verification check](https://docs.stripe.com/identity/verification-checks) to be performed. You must provide a `type` if not passing `verification_flow`.", "enum": ["document", "id_number"], "x-stripeBypassValidation": true}, "verification_flow": {"maxLength": 5000, "type": "string", "description": "The ID of a verification flow from the Dashboard. See https://docs.stripe.com/identity/verification-flows."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "options": {"style": "deepObject", "explode": true}, "provided_details": {"style": "deepObject", "explode": true}, "related_person": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/identity.verification_session"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
