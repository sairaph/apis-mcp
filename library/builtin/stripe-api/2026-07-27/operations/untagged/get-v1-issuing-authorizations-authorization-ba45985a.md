---
title: Retrieve an authorization
page_id: operation-get-v1-issuing-authorizations-authorization-5dd5afde
path: operations/untagged
description: <p>Retrieves an Issuing <code>Authorization</code> object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/issuing/authorizations/{authorization}
operation_ids:
    - GetIssuingAuthorizationsAuthorization
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve an authorization

`GET /v1/issuing/authorizations/{authorization}`

Operation ID: `GetIssuingAuthorizationsAuthorization`

<p>Retrieves an Issuing <code>Authorization</code> object.</p>

## Definition

```yaml
{"summary": "Retrieve an authorization", "description": "<p>Retrieves an Issuing <code>Authorization</code> object.</p>", "operationId": "GetIssuingAuthorizationsAuthorization", "parameters": [{"name": "authorization", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.authorization"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
