---
title: GET /v1/charges/{charge}/dispute
page_id: operation-get-v1-charges-charge-dispute-7a3cbfae
path: operations/untagged
description: <p>Retrieve a dispute for a specified charge.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/charges/{charge}/dispute
operation_ids:
    - GetChargesChargeDispute
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# GET /v1/charges/{charge}/dispute

`GET /v1/charges/{charge}/dispute`

Operation ID: `GetChargesChargeDispute`

<p>Retrieve a dispute for a specified charge.</p>

## Definition

```yaml
{"description": "<p>Retrieve a dispute for a specified charge.</p>", "operationId": "GetChargesChargeDispute", "parameters": [{"name": "charge", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dispute"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
