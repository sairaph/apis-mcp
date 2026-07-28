---
title: Create a Payment Evaluation
page_id: operation-post-v1-radar-payment-evaluations-a2930082
path: operations/untagged
description: <p>Request a Radar API fraud risk score from Stripe for a payment before sending it for external processor authorization.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/radar/payment_evaluations
operation_ids:
    - PostRadarPaymentEvaluations
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a Payment Evaluation

`POST /v1/radar/payment_evaluations`

Operation ID: `PostRadarPaymentEvaluations`

<p>Request a Radar API fraud risk score from Stripe for a payment before sending it for external processor authorization.</p>

## Definition

```yaml
{"summary": "Create a Payment Evaluation", "description": "<p>Request a Radar API fraud risk score from Stripe for a payment before sending it for external processor authorization.</p>", "operationId": "PostRadarPaymentEvaluations", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["customer_details", "payment_details"], "type": "object", "properties": {"client_device_metadata_details": {"title": "client_device_metadata_wrapper", "required": ["radar_session"], "type": "object", "properties": {"radar_session": {"maxLength": 5000, "type": "string"}}, "description": "Details about the Client Device Metadata to associate with the payment evaluation."}, "customer_details": {"title": "customer_details", "type": "object", "properties": {"customer": {"maxLength": 5000, "type": "string"}, "customer_account": {"maxLength": 5000, "type": "string"}, "email": {"type": "string"}, "name": {"maxLength": 5000, "type": "string"}, "phone": {"type": "string"}}, "description": "Details about the customer associated with the payment evaluation."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "payment_details": {"title": "payment_details", "required": ["amount", "currency", "payment_method_details"], "type": "object", "properties": {"amount": {"type": "integer"}, "currency": {"type": "string", "format": "currency"}, "description": {"maxLength": 5000, "type": "string"}, "money_movement_details": {"title": "money_movement_details", "required": ["money_movement_type"], "type": "object", "properties": {"card": {"title": "money_movement_card_additional_data", "type": "object", "properties": {"customer_presence": {"type": "string", "enum": ["off_session", "on_session"]}, "payment_type": {"type": "string", "enum": ["one_off", "recurring", "setup_one_off", "setup_recurring"]}}}, "money_movement_type": {"type": "string", "enum": ["card"]}}}, "payment_method_details": {"title": "payment_method_details", "required": ["payment_method"], "type": "object", "properties": {"billing_details": {"title": "billing_details", "type": "object", "properties": {"address": {"title": "address", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "line1": {"maxLength": 5000, "type": "string"}, "line2": {"maxLength": 5000, "type": "string"}, "postal_code": {"maxLength": 5000, "type": "string"}, "state": {"maxLength": 5000, "type": "string"}}}, "email": {"type": "string"}, "name": {"maxLength": 5000, "type": "string"}, "phone": {"type": "string"}}}, "payment_method": {"maxLength": 5000, "type": "string"}}}, "shipping_details": {"title": "shipping_details", "type": "object", "properties": {"address": {"title": "address", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "line1": {"maxLength": 5000, "type": "string"}, "line2": {"maxLength": 5000, "type": "string"}, "postal_code": {"maxLength": 5000, "type": "string"}, "state": {"maxLength": 5000, "type": "string"}}}, "name": {"maxLength": 5000, "type": "string"}, "phone": {"type": "string"}}}, "statement_descriptor": {"maxLength": 5000, "type": "string"}}, "description": "Details about the payment."}}, "additionalProperties": false}, "encoding": {"client_device_metadata_details": {"style": "deepObject", "explode": true}, "customer_details": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "payment_details": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/radar.payment_evaluation"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
