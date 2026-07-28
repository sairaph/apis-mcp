---
title: Retrieve a Payment Attempt Record
page_id: operation-get-v1-payment-attempt-records-id-c4d054a9
path: operations/untagged
description: <p>Retrieves a Payment Attempt Record with the given ID</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/payment_attempt_records/{id}
operation_ids:
    - GetPaymentAttemptRecordsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a Payment Attempt Record

`GET /v1/payment_attempt_records/{id}`

Operation ID: `GetPaymentAttemptRecordsId`

<p>Retrieves a Payment Attempt Record with the given ID</p>

## Definition

```yaml
{"summary": "Retrieve a Payment Attempt Record", "description": "<p>Retrieves a Payment Attempt Record with the given ID</p>", "operationId": "GetPaymentAttemptRecordsId", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "id", "in": "path", "description": "The ID of the Payment Attempt Record.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_attempt_record"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
