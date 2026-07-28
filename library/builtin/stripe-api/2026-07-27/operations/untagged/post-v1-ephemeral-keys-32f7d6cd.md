---
title: Create an ephemeral key
page_id: operation-post-v1-ephemeral-keys-03229f3c
path: operations/untagged
description: <p>Creates a short-lived API key for a given resource.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/ephemeral_keys
operation_ids:
    - PostEphemeralKeys
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create an ephemeral key

`POST /v1/ephemeral_keys`

Operation ID: `PostEphemeralKeys`

<p>Creates a short-lived API key for a given resource.</p>

## Definition

```yaml
{"summary": "Create an ephemeral key", "description": "<p>Creates a short-lived API key for a given resource.</p>", "operationId": "PostEphemeralKeys", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"customer": {"maxLength": 5000, "type": "string", "description": "The ID of the Customer you'd like to modify using the resulting ephemeral key."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "issuing_card": {"maxLength": 5000, "type": "string", "description": "The ID of the Issuing Card you'd like to access using the resulting ephemeral key."}, "nonce": {"maxLength": 5000, "type": "string", "description": "A single-use token, created by Stripe.js, used for creating ephemeral keys for Issuing Cards without exchanging sensitive information."}, "verification_session": {"maxLength": 5000, "type": "string", "description": "The ID of the Identity VerificationSession you'd like to access using the resulting ephemeral key"}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ephemeral_key"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
