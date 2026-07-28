---
title: Retrieve a Customer's PaymentMethod
page_id: operation-get-v1-customers-customer-payment-methods-payment-method-dad03dfb
path: operations/untagged
description: <p>Retrieves a PaymentMethod object for a given Customer.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/customers/{customer}/payment_methods/{payment_method}
operation_ids:
    - GetCustomersCustomerPaymentMethodsPaymentMethod
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a Customer's PaymentMethod

`GET /v1/customers/{customer}/payment_methods/{payment_method}`

Operation ID: `GetCustomersCustomerPaymentMethodsPaymentMethod`

<p>Retrieves a PaymentMethod object for a given Customer.</p>

## Definition

```yaml
{"summary": "Retrieve a Customer's PaymentMethod", "description": "<p>Retrieves a PaymentMethod object for a given Customer.</p>", "operationId": "GetCustomersCustomerPaymentMethodsPaymentMethod", "parameters": [{"name": "customer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "payment_method", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_method"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
