---
title: Report payment attempt canceled
page_id: operation-post-v1-payment-records-id-report-payment-attempt-canceled-fb823edd
path: operations/untagged
description: |-
    <p>Report that the most recent payment attempt on the specified Payment Record
     was canceled.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/payment_records/{id}/report_payment_attempt_canceled
operation_ids:
    - PostPaymentRecordsIdReportPaymentAttemptCanceled
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Report payment attempt canceled

`POST /v1/payment_records/{id}/report_payment_attempt_canceled`

Operation ID: `PostPaymentRecordsIdReportPaymentAttemptCanceled`

<p>Report that the most recent payment attempt on the specified Payment Record
 was canceled.</p>

## Definition

```yaml
{"summary": "Report payment attempt canceled", "description": "<p>Report that the most recent payment attempt on the specified Payment Record\n was canceled.</p>", "operationId": "PostPaymentRecordsIdReportPaymentAttemptCanceled", "parameters": [{"name": "id", "in": "path", "description": "The ID of the Payment Record.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["canceled_at"], "type": "object", "properties": {"canceled_at": {"type": "integer", "description": "When the reported payment was canceled. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_record"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
