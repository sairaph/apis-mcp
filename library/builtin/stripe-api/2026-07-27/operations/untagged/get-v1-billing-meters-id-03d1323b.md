---
title: Retrieve a billing meter
page_id: operation-get-v1-billing-meters-id-131f7def
path: operations/untagged
description: <p>Retrieves a billing meter given an ID.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/billing/meters/{id}
operation_ids:
    - GetBillingMetersId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a billing meter

`GET /v1/billing/meters/{id}`

Operation ID: `GetBillingMetersId`

<p>Retrieves a billing meter given an ID.</p>

## Definition

```yaml
{"summary": "Retrieve a billing meter", "description": "<p>Retrieves a billing meter given an ID.</p>", "operationId": "GetBillingMetersId", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billing.meter"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
