---
title: Expire a test-mode authorization
page_id: operation-post-v1-test-helpers-issuing-authorizations-authorization-expire-0acf46ac
path: operations/untagged
description: <p>Expire a test-mode Authorization.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/issuing/authorizations/{authorization}/expire
operation_ids:
    - PostTestHelpersIssuingAuthorizationsAuthorizationExpire
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Expire a test-mode authorization

`POST /v1/test_helpers/issuing/authorizations/{authorization}/expire`

Operation ID: `PostTestHelpersIssuingAuthorizationsAuthorizationExpire`

<p>Expire a test-mode Authorization.</p>

## Definition

```yaml
{"summary": "Expire a test-mode authorization", "description": "<p>Expire a test-mode Authorization.</p>", "operationId": "PostTestHelpersIssuingAuthorizationsAuthorizationExpire", "parameters": [{"name": "authorization", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.authorization"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
