---
title: Retrieve a ConfirmationToken
page_id: operation-get-v1-confirmation-tokens-confirmation-token-6ab734ab
path: operations/untagged
description: <p>Retrieves an existing ConfirmationToken object</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/confirmation_tokens/{confirmation_token}
operation_ids:
    - GetConfirmationTokensConfirmationToken
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a ConfirmationToken

`GET /v1/confirmation_tokens/{confirmation_token}`

Operation ID: `GetConfirmationTokensConfirmationToken`

<p>Retrieves an existing ConfirmationToken object</p>

## Definition

```yaml
{"summary": "Retrieve a ConfirmationToken", "description": "<p>Retrieves an existing ConfirmationToken object</p>", "operationId": "GetConfirmationTokensConfirmationToken", "parameters": [{"name": "confirmation_token", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/confirmation_token"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
