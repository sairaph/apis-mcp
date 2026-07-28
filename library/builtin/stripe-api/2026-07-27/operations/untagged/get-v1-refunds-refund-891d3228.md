---
title: Retrieve a refund
page_id: operation-get-v1-refunds-refund-ee739289
path: operations/untagged
description: <p>Retrieves the details of an existing refund.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/refunds/{refund}
operation_ids:
    - GetRefundsRefund
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a refund

`GET /v1/refunds/{refund}`

Operation ID: `GetRefundsRefund`

<p>Retrieves the details of an existing refund.</p>

## Definition

```yaml
{"summary": "Retrieve a refund", "description": "<p>Retrieves the details of an existing refund.</p>", "operationId": "GetRefundsRefund", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "refund", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/refund"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
