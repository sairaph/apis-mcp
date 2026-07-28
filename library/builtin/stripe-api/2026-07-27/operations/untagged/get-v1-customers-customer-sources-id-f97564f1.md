---
title: GET /v1/customers/{customer}/sources/{id}
page_id: operation-get-v1-customers-customer-sources-id-379bd5f1
path: operations/untagged
description: <p>Retrieve a specified source for a given customer.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/customers/{customer}/sources/{id}
operation_ids:
    - GetCustomersCustomerSourcesId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# GET /v1/customers/{customer}/sources/{id}

`GET /v1/customers/{customer}/sources/{id}`

Operation ID: `GetCustomersCustomerSourcesId`

<p>Retrieve a specified source for a given customer.</p>

## Definition

```yaml
{"description": "<p>Retrieve a specified source for a given customer.</p>", "operationId": "GetCustomersCustomerSourcesId", "parameters": [{"name": "customer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 500, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_source"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
