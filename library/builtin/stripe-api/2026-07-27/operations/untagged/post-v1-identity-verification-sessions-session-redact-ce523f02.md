---
title: Redact a VerificationSession
page_id: operation-post-v1-identity-verification-sessions-session-redact-0f2a713d
path: operations/untagged
description: |-
    <p>Redact a VerificationSession to remove all collected information from Stripe. This will redact
    the VerificationSession and all objects related to it, including VerificationReports, Events,
    request logs, etc.</p>

    <p>A VerificationSession object can be redacted when it is in <code>requires_input</code> or <code>verified</code>
    <a href="/docs/identity/how-sessions-work">status</a>. Redacting a VerificationSession in <code>requires_action</code>
    state will automatically cancel it.</p>

    <p>The redaction process may take up to four days. When the redaction process is in progress, the
    VerificationSession’s <code>redaction.status</code> field will be set to <code>processing</code>; when the process is
    finished, it will change to <code>redacted</code> and an <code>identity.verification_session.redacted</code> event
    will be emitted.</p>

    <p>Redaction is irreversible. Redacted objects are still accessible in the Stripe API, but all the
    fields that contain personal data will be replaced by the string <code>[redacted]</code> or a similar
    placeholder. The <code>metadata</code> field will also be erased. Redacted objects cannot be updated or
    used for any purpose.</p>

    <p><a href="/docs/identity/verification-sessions#redact">Learn more</a>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/identity/verification_sessions/{session}/redact
operation_ids:
    - PostIdentityVerificationSessionsSessionRedact
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Redact a VerificationSession

`POST /v1/identity/verification_sessions/{session}/redact`

Operation ID: `PostIdentityVerificationSessionsSessionRedact`

<p>Redact a VerificationSession to remove all collected information from Stripe. This will redact
the VerificationSession and all objects related to it, including VerificationReports, Events,
request logs, etc.</p>

<p>A VerificationSession object can be redacted when it is in <code>requires_input</code> or <code>verified</code>
<a href="/docs/identity/how-sessions-work">status</a>. Redacting a VerificationSession in <code>requires_action</code>
state will automatically cancel it.</p>

<p>The redaction process may take up to four days. When the redaction process is in progress, the
VerificationSession’s <code>redaction.status</code> field will be set to <code>processing</code>; when the process is
finished, it will change to <code>redacted</code> and an <code>identity.verification_session.redacted</code> event
will be emitted.</p>

<p>Redaction is irreversible. Redacted objects are still accessible in the Stripe API, but all the
fields that contain personal data will be replaced by the string <code>[redacted]</code> or a similar
placeholder. The <code>metadata</code> field will also be erased. Redacted objects cannot be updated or
used for any purpose.</p>

<p><a href="/docs/identity/verification-sessions#redact">Learn more</a>.</p>

## Definition

```yaml
{"summary": "Redact a VerificationSession", "description": "<p>Redact a VerificationSession to remove all collected information from Stripe. This will redact\nthe VerificationSession and all objects related to it, including VerificationReports, Events,\nrequest logs, etc.</p>\n\n<p>A VerificationSession object can be redacted when it is in <code>requires_input</code> or <code>verified</code>\n<a href=\"/docs/identity/how-sessions-work\">status</a>. Redacting a VerificationSession in <code>requires_action</code>\nstate will automatically cancel it.</p>\n\n<p>The redaction process may take up to four days. When the redaction process is in progress, the\nVerificationSession’s <code>redaction.status</code> field will be set to <code>processing</code>; when the process is\nfinished, it will change to <code>redacted</code> and an <code>identity.verification_session.redacted</code> event\nwill be emitted.</p>\n\n<p>Redaction is irreversible. Redacted objects are still accessible in the Stripe API, but all the\nfields that contain personal data will be replaced by the string <code>[redacted]</code> or a similar\nplaceholder. The <code>metadata</code> field will also be erased. Redacted objects cannot be updated or\nused for any purpose.</p>\n\n<p><a href=\"/docs/identity/verification-sessions#redact\">Learn more</a>.</p>", "operationId": "PostIdentityVerificationSessionsSessionRedact", "parameters": [{"name": "session", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/identity.verification_session"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
