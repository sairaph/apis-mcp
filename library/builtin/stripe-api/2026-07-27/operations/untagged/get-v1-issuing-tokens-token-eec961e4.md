---
title: Retrieve an issuing token
page_id: operation-get-v1-issuing-tokens-token-d609a46f
path: operations/untagged
description: <p>Retrieves an Issuing <code>Token</code> object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/issuing/tokens/{token}
operation_ids:
    - GetIssuingTokensToken
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve an issuing token

`GET /v1/issuing/tokens/{token}`

Operation ID: `GetIssuingTokensToken`

<p>Retrieves an Issuing <code>Token</code> object.</p>

## Definition

```yaml
{"summary": "Retrieve an issuing token", "description": "<p>Retrieves an Issuing <code>Token</code> object.</p>", "operationId": "GetIssuingTokensToken", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "token", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.token"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
