---
title: Approve an authorization
page_id: operation-post-v1-issuing-authorizations-authorization-approve-4bdcab11
path: operations/untagged
description: "<p>[Deprecated] Approves a pending Issuing <code>Authorization</code> object. This request should be made within the timeout window of the <a href=\"/docs/issuing/controls/real-time-authorizations\">real-time authorization</a> flow. \nThis method is deprecated. Instead, <a href=\"/docs/issuing/controls/real-time-authorizations#authorization-handling\">respond directly to the webhook request to approve an authorization</a>.</p>"
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/issuing/authorizations/{authorization}/approve
operation_ids:
    - PostIssuingAuthorizationsAuthorizationApprove
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Approve an authorization

`POST /v1/issuing/authorizations/{authorization}/approve`

Operation ID: `PostIssuingAuthorizationsAuthorizationApprove`

<p>[Deprecated] Approves a pending Issuing <code>Authorization</code> object. This request should be made within the timeout window of the <a href="/docs/issuing/controls/real-time-authorizations">real-time authorization</a> flow.
This method is deprecated. Instead, <a href="/docs/issuing/controls/real-time-authorizations#authorization-handling">respond directly to the webhook request to approve an authorization</a>.</p>

## Definition

```yaml
{"summary": "Approve an authorization", "description": "<p>[Deprecated] Approves a pending Issuing <code>Authorization</code> object. This request should be made within the timeout window of the <a href=\"/docs/issuing/controls/real-time-authorizations\">real-time authorization</a> flow. \nThis method is deprecated. Instead, <a href=\"/docs/issuing/controls/real-time-authorizations#authorization-handling\">respond directly to the webhook request to approve an authorization</a>.</p>", "operationId": "PostIssuingAuthorizationsAuthorizationApprove", "parameters": [{"name": "authorization", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"amount": {"type": "integer", "description": "If the authorization's `pending_request.is_amount_controllable` property is `true`, you may provide this value to control how much to hold for the authorization. Must be positive (use [`decline`](https://docs.stripe.com/api/issuing/authorizations/decline) to decline an authorization request)."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.authorization"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}, "deprecated": true}
```
