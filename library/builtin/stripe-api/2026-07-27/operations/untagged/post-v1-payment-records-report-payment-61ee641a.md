---
title: Report a payment
page_id: operation-post-v1-payment-records-report-payment-2e420148
path: operations/untagged
description: |-
    <p>Report a new Payment Record. You may report a Payment Record as it is
     initialized and later report updates through the other report_* methods, or report Payment
     Records in a terminal state directly, through this method.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/payment_records/report_payment
operation_ids:
    - PostPaymentRecordsReportPayment
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Report a payment

`POST /v1/payment_records/report_payment`

Operation ID: `PostPaymentRecordsReportPayment`

<p>Report a new Payment Record. You may report a Payment Record as it is
 initialized and later report updates through the other report_* methods, or report Payment
 Records in a terminal state directly, through this method.</p>

## Definition

```yaml
{"summary": "Report a payment", "description": "<p>Report a new Payment Record. You may report a Payment Record as it is\n initialized and later report updates through the other report_* methods, or report Payment\n Records in a terminal state directly, through this method.</p>", "operationId": "PostPaymentRecordsReportPayment", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["amount_requested", "initiated_at", "payment_method_details"], "type": "object", "properties": {"amount_requested": {"title": "amount", "required": ["currency", "value"], "type": "object", "properties": {"currency": {"type": "string", "format": "currency"}, "value": {"type": "integer"}}, "description": "The amount you initially requested for this payment."}, "customer_details": {"title": "customer_details", "type": "object", "properties": {"customer": {"maxLength": 5000, "type": "string"}, "email": {"type": "string"}, "name": {"maxLength": 5000, "type": "string"}, "phone": {"type": "string"}}, "description": "Customer information for this payment."}, "customer_presence": {"type": "string", "description": "Indicates whether the customer was present in your checkout flow during this payment.", "enum": ["off_session", "on_session"]}, "description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the object. Often useful for displaying to users."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "failed": {"title": "failed", "required": ["failed_at"], "type": "object", "properties": {"failed_at": {"type": "integer", "format": "unix-time"}}, "description": "Information about the payment attempt failure."}, "guaranteed": {"title": "guaranteed", "required": ["guaranteed_at"], "type": "object", "properties": {"guaranteed_at": {"type": "integer", "format": "unix-time"}}, "description": "Information about the payment attempt guarantee."}, "initiated_at": {"type": "integer", "description": "When the reported payment was initiated. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "outcome": {"type": "string", "description": "The outcome of the reported payment.", "enum": ["failed", "guaranteed"]}, "payment_method_details": {"title": "payment_method_details", "type": "object", "properties": {"billing_details": {"title": "billing_details", "type": "object", "properties": {"address": {"title": "address", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "line1": {"maxLength": 5000, "type": "string"}, "line2": {"maxLength": 5000, "type": "string"}, "postal_code": {"maxLength": 5000, "type": "string"}, "state": {"maxLength": 5000, "type": "string"}}}, "email": {"type": "string"}, "name": {"maxLength": 5000, "type": "string"}, "phone": {"type": "string"}}}, "custom": {"title": "custom", "type": "object", "properties": {"display_name": {"maxLength": 5000, "type": "string"}, "type": {"maxLength": 5000, "type": "string"}}}, "payment_method": {"maxLength": 5000, "type": "string"}, "type": {"type": "string", "enum": ["custom"], "x-stripeBypassValidation": true}}, "description": "Information about the Payment Method debited for this payment."}, "processor_details": {"title": "processor_details", "required": ["type"], "type": "object", "properties": {"custom": {"title": "custom", "required": ["payment_reference"], "type": "object", "properties": {"payment_reference": {"maxLength": 5000, "type": "string"}}}, "type": {"type": "string", "enum": ["custom"]}}, "description": "Processor information for this payment."}, "shipping_details": {"title": "shipping_details", "type": "object", "properties": {"address": {"title": "address", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "line1": {"maxLength": 5000, "type": "string"}, "line2": {"maxLength": 5000, "type": "string"}, "postal_code": {"maxLength": 5000, "type": "string"}, "state": {"maxLength": 5000, "type": "string"}}}, "name": {"maxLength": 5000, "type": "string"}, "phone": {"type": "string"}}, "description": "Shipping information for this payment."}}, "additionalProperties": false}, "encoding": {"amount_requested": {"style": "deepObject", "explode": true}, "customer_details": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "failed": {"style": "deepObject", "explode": true}, "guaranteed": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "payment_method_details": {"style": "deepObject", "explode": true}, "processor_details": {"style": "deepObject", "explode": true}, "shipping_details": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_record"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
