---
title: Report payment attempt guaranteed
page_id: operation-post-v1-payment-records-id-report-payment-attempt-guaranteed-d601bea0
path: operations/untagged
description: |-
    <p>Report that the most recent payment attempt on the specified Payment Record
     was guaranteed.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/payment_records/{id}/report_payment_attempt_guaranteed
operation_ids:
    - PostPaymentRecordsIdReportPaymentAttemptGuaranteed
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Report payment attempt guaranteed

`POST /v1/payment_records/{id}/report_payment_attempt_guaranteed`

Operation ID: `PostPaymentRecordsIdReportPaymentAttemptGuaranteed`

<p>Report that the most recent payment attempt on the specified Payment Record
 was guaranteed.</p>

## Definition

```yaml
{"summary": "Report payment attempt guaranteed", "description": "<p>Report that the most recent payment attempt on the specified Payment Record\n was guaranteed.</p>", "operationId": "PostPaymentRecordsIdReportPaymentAttemptGuaranteed", "parameters": [{"name": "id", "in": "path", "description": "The ID of the Payment Record.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["guaranteed_at"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "guaranteed_at": {"type": "integer", "description": "When the reported payment was guaranteed. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_record"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
