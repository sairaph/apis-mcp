---
title: Increment a test-mode authorization
page_id: operation-post-v1-test-helpers-issuing-authorizations-authorization-increment-32211790
path: operations/untagged
description: <p>Increment a test-mode Authorization.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/issuing/authorizations/{authorization}/increment
operation_ids:
    - PostTestHelpersIssuingAuthorizationsAuthorizationIncrement
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Increment a test-mode authorization

`POST /v1/test_helpers/issuing/authorizations/{authorization}/increment`

Operation ID: `PostTestHelpersIssuingAuthorizationsAuthorizationIncrement`

<p>Increment a test-mode Authorization.</p>

## Definition

```yaml
{"summary": "Increment a test-mode authorization", "description": "<p>Increment a test-mode Authorization.</p>", "operationId": "PostTestHelpersIssuingAuthorizationsAuthorizationIncrement", "parameters": [{"name": "authorization", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["increment_amount"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "increment_amount": {"type": "integer", "description": "The amount to increment the authorization by. This amount is in the authorization currency and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal)."}, "is_amount_controllable": {"type": "boolean", "description": "If set `true`, you may provide [amount](https://docs.stripe.com/api/issuing/authorizations/approve#approve_issuing_authorization-amount) to control how much to hold for the authorization."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.authorization"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
