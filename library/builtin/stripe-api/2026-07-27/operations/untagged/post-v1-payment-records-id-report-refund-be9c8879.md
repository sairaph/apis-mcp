---
title: Report a refund
page_id: operation-post-v1-payment-records-id-report-refund-f134c026
path: operations/untagged
description: |-
    <p>Report that the most recent payment attempt on the specified Payment Record
     was refunded.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/payment_records/{id}/report_refund
operation_ids:
    - PostPaymentRecordsIdReportRefund
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Report a refund

`POST /v1/payment_records/{id}/report_refund`

Operation ID: `PostPaymentRecordsIdReportRefund`

<p>Report that the most recent payment attempt on the specified Payment Record
 was refunded.</p>

## Definition

```yaml
{"summary": "Report a refund", "description": "<p>Report that the most recent payment attempt on the specified Payment Record\n was refunded.</p>", "operationId": "PostPaymentRecordsIdReportRefund", "parameters": [{"name": "id", "in": "path", "description": "The ID of the Payment Record.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["outcome", "processor_details"], "type": "object", "properties": {"amount": {"title": "amount", "required": ["currency", "value"], "type": "object", "properties": {"currency": {"type": "string", "format": "currency"}, "value": {"type": "integer"}}, "description": "A positive integer in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) representing how much of this payment to refund. Can refund only up to the remaining, unrefunded amount of the payment."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "initiated_at": {"type": "integer", "description": "When the reported refund was initiated. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "outcome": {"type": "string", "description": "The outcome of the reported refund.", "enum": ["refunded"], "x-stripeBypassValidation": true}, "processor_details": {"title": "processor_details", "required": ["type"], "type": "object", "properties": {"custom": {"title": "custom", "required": ["refund_reference"], "type": "object", "properties": {"refund_reference": {"maxLength": 5000, "type": "string"}}}, "type": {"type": "string", "enum": ["custom"]}}, "description": "Processor information for this refund."}, "refunded": {"title": "refunded", "required": ["refunded_at"], "type": "object", "properties": {"refunded_at": {"type": "integer", "format": "unix-time"}}, "description": "Information about the payment attempt refund."}}, "additionalProperties": false}, "encoding": {"amount": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "processor_details": {"style": "deepObject", "explode": true}, "refunded": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_record"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
