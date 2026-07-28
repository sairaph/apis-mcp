---
title: Retrieve a PaymentIntent
page_id: operation-get-v1-payment-intents-intent-6d378891
path: operations/untagged
description: |-
    <p>Retrieves the details of a PaymentIntent that has previously been created. </p>

    <p>You can retrieve a PaymentIntent client-side using a publishable key when the <code>client_secret</code> is in the query string. </p>

    <p>If you retrieve a PaymentIntent with a publishable key, it only returns a subset of properties. Refer to the <a href="#payment_intent_object">payment intent</a> object reference for more details.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/payment_intents/{intent}
operation_ids:
    - GetPaymentIntentsIntent
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a PaymentIntent

`GET /v1/payment_intents/{intent}`

Operation ID: `GetPaymentIntentsIntent`

<p>Retrieves the details of a PaymentIntent that has previously been created. </p>

<p>You can retrieve a PaymentIntent client-side using a publishable key when the <code>client_secret</code> is in the query string. </p>

<p>If you retrieve a PaymentIntent with a publishable key, it only returns a subset of properties. Refer to the <a href="#payment_intent_object">payment intent</a> object reference for more details.</p>

## Definition

```yaml
{"summary": "Retrieve a PaymentIntent", "description": "<p>Retrieves the details of a PaymentIntent that has previously been created. </p>\n\n<p>You can retrieve a PaymentIntent client-side using a publishable key when the <code>client_secret</code> is in the query string. </p>\n\n<p>If you retrieve a PaymentIntent with a publishable key, it only returns a subset of properties. Refer to the <a href=\"#payment_intent_object\">payment intent</a> object reference for more details.</p>", "operationId": "GetPaymentIntentsIntent", "parameters": [{"name": "client_secret", "in": "query", "description": "The client secret of the PaymentIntent. We require it if you use a publishable key to retrieve the source.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "intent", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_intent"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
