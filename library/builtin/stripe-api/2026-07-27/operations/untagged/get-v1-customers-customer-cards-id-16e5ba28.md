---
title: Retrieve a card
page_id: operation-get-v1-customers-customer-cards-id-f00bbf43
path: operations/untagged
description: <p>You can always see the 10 most recent cards directly on a customer; this method lets you retrieve details about a specific card stored on the customer.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/customers/{customer}/cards/{id}
operation_ids:
    - GetCustomersCustomerCardsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a card

`GET /v1/customers/{customer}/cards/{id}`

Operation ID: `GetCustomersCustomerCardsId`

<p>You can always see the 10 most recent cards directly on a customer; this method lets you retrieve details about a specific card stored on the customer.</p>

## Definition

```yaml
{"summary": "Retrieve a card", "description": "<p>You can always see the 10 most recent cards directly on a customer; this method lets you retrieve details about a specific card stored on the customer.</p>", "operationId": "GetCustomersCustomerCardsId", "parameters": [{"name": "customer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/card"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}, "deprecated": true}
```
