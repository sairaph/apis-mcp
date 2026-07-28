---
title: Report payment attempt informational
page_id: operation-post-v1-payment-records-id-report-payment-attempt-informational-8c4917b4
path: operations/untagged
description: <p>Report informational updates on the specified Payment Record.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/payment_records/{id}/report_payment_attempt_informational
operation_ids:
    - PostPaymentRecordsIdReportPaymentAttemptInformational
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Report payment attempt informational

`POST /v1/payment_records/{id}/report_payment_attempt_informational`

Operation ID: `PostPaymentRecordsIdReportPaymentAttemptInformational`

<p>Report informational updates on the specified Payment Record.</p>

## Definition

```yaml
{"summary": "Report payment attempt informational", "description": "<p>Report informational updates on the specified Payment Record.</p>", "operationId": "PostPaymentRecordsIdReportPaymentAttemptInformational", "parameters": [{"name": "id", "in": "path", "description": "The ID of the Payment Record.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"customer_details": {"title": "customer_details", "type": "object", "properties": {"customer": {"maxLength": 5000, "type": "string"}, "email": {"type": "string"}, "name": {"maxLength": 5000, "type": "string"}, "phone": {"type": "string"}}, "description": "Customer information for this payment."}, "description": {"description": "An arbitrary string attached to the object. Often useful for displaying to users.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"type": "string", "enum": [""]}]}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "shipping_details": {"description": "Shipping information for this payment.", "anyOf": [{"title": "shipping_details", "type": "object", "properties": {"address": {"title": "address", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "line1": {"maxLength": 5000, "type": "string"}, "line2": {"maxLength": 5000, "type": "string"}, "postal_code": {"maxLength": 5000, "type": "string"}, "state": {"maxLength": 5000, "type": "string"}}}, "name": {"maxLength": 5000, "type": "string"}, "phone": {"type": "string"}}}, {"type": "string", "enum": [""]}]}}, "additionalProperties": false}, "encoding": {"customer_details": {"style": "deepObject", "explode": true}, "description": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "shipping_details": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_record"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
