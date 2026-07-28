---
title: Retrieve a cardholder
page_id: operation-get-v1-issuing-cardholders-cardholder-3f66a185
path: operations/untagged
description: <p>Retrieves an Issuing <code>Cardholder</code> object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/issuing/cardholders/{cardholder}
operation_ids:
    - GetIssuingCardholdersCardholder
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a cardholder

`GET /v1/issuing/cardholders/{cardholder}`

Operation ID: `GetIssuingCardholdersCardholder`

<p>Retrieves an Issuing <code>Cardholder</code> object.</p>

## Definition

```yaml
{"summary": "Retrieve a cardholder", "description": "<p>Retrieves an Issuing <code>Cardholder</code> object.</p>", "operationId": "GetIssuingCardholdersCardholder", "parameters": [{"name": "cardholder", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.cardholder"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
