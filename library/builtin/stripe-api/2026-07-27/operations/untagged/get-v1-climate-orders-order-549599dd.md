---
title: Retrieve an order
page_id: operation-get-v1-climate-orders-order-f386a286
path: operations/untagged
description: <p>Retrieves the details of a Climate order object with the given ID.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/climate/orders/{order}
operation_ids:
    - GetClimateOrdersOrder
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve an order

`GET /v1/climate/orders/{order}`

Operation ID: `GetClimateOrdersOrder`

<p>Retrieves the details of a Climate order object with the given ID.</p>

## Definition

```yaml
{"summary": "Retrieve an order", "description": "<p>Retrieves the details of a Climate order object with the given ID.</p>", "operationId": "GetClimateOrdersOrder", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "order", "in": "path", "description": "Unique identifier of the order.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/climate.order"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
