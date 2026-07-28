---
title: Reverse a test-mode authorization
page_id: operation-post-v1-test-helpers-issuing-authorizations-authorization-reverse-eae0469d
path: operations/untagged
description: <p>Reverse a test-mode Authorization.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/issuing/authorizations/{authorization}/reverse
operation_ids:
    - PostTestHelpersIssuingAuthorizationsAuthorizationReverse
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Reverse a test-mode authorization

`POST /v1/test_helpers/issuing/authorizations/{authorization}/reverse`

Operation ID: `PostTestHelpersIssuingAuthorizationsAuthorizationReverse`

<p>Reverse a test-mode Authorization.</p>

## Definition

```yaml
{"summary": "Reverse a test-mode authorization", "description": "<p>Reverse a test-mode Authorization.</p>", "operationId": "PostTestHelpersIssuingAuthorizationsAuthorizationReverse", "parameters": [{"name": "authorization", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "reverse_amount": {"type": "integer", "description": "The amount to reverse from the authorization. If not provided, the full amount of the authorization will be reversed. This amount is in the authorization currency and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal)."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.authorization"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
