---
title: Report a payment attempt
page_id: operation-post-v1-payment-records-id-report-payment-attempt-8fe6d1f2
path: operations/untagged
description: |-
    <p>Report a new payment attempt on the specified Payment Record. A new payment
     attempt can only be specified if all other payment attempts are canceled or failed.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/payment_records/{id}/report_payment_attempt
operation_ids:
    - PostPaymentRecordsIdReportPaymentAttempt
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Report a payment attempt

`POST /v1/payment_records/{id}/report_payment_attempt`

Operation ID: `PostPaymentRecordsIdReportPaymentAttempt`

<p>Report a new payment attempt on the specified Payment Record. A new payment
 attempt can only be specified if all other payment attempts are canceled or failed.</p>

## Definition

```yaml
{"summary": "Report a payment attempt", "description": "<p>Report a new payment attempt on the specified Payment Record. A new payment\n attempt can only be specified if all other payment attempts are canceled or failed.</p>", "operationId": "PostPaymentRecordsIdReportPaymentAttempt", "parameters": [{"name": "id", "in": "path", "description": "The ID of the Payment Record.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["initiated_at"], "type": "object", "properties": {"description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the object. Often useful for displaying to users."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "failed": {"title": "failed", "required": ["failed_at"], "type": "object", "properties": {"failed_at": {"type": "integer", "format": "unix-time"}}, "description": "Information about the payment attempt failure."}, "guaranteed": {"title": "guaranteed", "required": ["guaranteed_at"], "type": "object", "properties": {"guaranteed_at": {"type": "integer", "format": "unix-time"}}, "description": "Information about the payment attempt guarantee."}, "initiated_at": {"type": "integer", "description": "When the reported payment was initiated. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "outcome": {"type": "string", "description": "The outcome of the reported payment.", "enum": ["failed", "guaranteed"]}, "payment_method_details": {"title": "payment_method_details", "type": "object", "properties": {"billing_details": {"title": "billing_details", "type": "object", "properties": {"address": {"title": "address", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "line1": {"maxLength": 5000, "type": "string"}, "line2": {"maxLength": 5000, "type": "string"}, "postal_code": {"maxLength": 5000, "type": "string"}, "state": {"maxLength": 5000, "type": "string"}}}, "email": {"type": "string"}, "name": {"maxLength": 5000, "type": "string"}, "phone": {"type": "string"}}}, "custom": {"title": "custom", "type": "object", "properties": {"display_name": {"maxLength": 5000, "type": "string"}, "type": {"maxLength": 5000, "type": "string"}}}, "payment_method": {"maxLength": 5000, "type": "string"}, "type": {"type": "string", "enum": ["custom"], "x-stripeBypassValidation": true}}, "description": "Information about the Payment Method debited for this payment."}, "shipping_details": {"title": "shipping_details", "type": "object", "properties": {"address": {"title": "address", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "line1": {"maxLength": 5000, "type": "string"}, "line2": {"maxLength": 5000, "type": "string"}, "postal_code": {"maxLength": 5000, "type": "string"}, "state": {"maxLength": 5000, "type": "string"}}}, "name": {"maxLength": 5000, "type": "string"}, "phone": {"type": "string"}}, "description": "Shipping information for this payment."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "failed": {"style": "deepObject", "explode": true}, "guaranteed": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "payment_method_details": {"style": "deepObject", "explode": true}, "shipping_details": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_record"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
