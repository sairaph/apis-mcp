---
title: Find a Tax Association
page_id: operation-get-v1-tax-associations-find-16754f0c
path: operations/untagged
description: <p>Finds a tax association object by PaymentIntent id.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/tax/associations/find
operation_ids:
    - GetTaxAssociationsFind
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Find a Tax Association

`GET /v1/tax/associations/find`

Operation ID: `GetTaxAssociationsFind`

<p>Finds a tax association object by PaymentIntent id.</p>

## Definition

```yaml
{"summary": "Find a Tax Association", "description": "<p>Finds a tax association object by PaymentIntent id.</p>", "operationId": "GetTaxAssociationsFind", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "payment_intent", "in": "query", "description": "Valid [PaymentIntent](https://docs.stripe.com/api/payment_intents/object) id", "required": true, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tax.association"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
