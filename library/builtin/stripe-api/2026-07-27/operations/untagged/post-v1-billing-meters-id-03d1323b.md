---
title: Update a billing meter
page_id: operation-post-v1-billing-meters-id-e981aa4d
path: operations/untagged
description: <p>Updates a billing meter.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/billing/meters/{id}
operation_ids:
    - PostBillingMetersId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a billing meter

`POST /v1/billing/meters/{id}`

Operation ID: `PostBillingMetersId`

<p>Updates a billing meter.</p>

## Definition

```yaml
{"summary": "Update a billing meter", "description": "<p>Updates a billing meter.</p>", "operationId": "PostBillingMetersId", "parameters": [{"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"display_name": {"maxLength": 250, "type": "string", "description": "The meter’s name. Not visible to the customer."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billing.meter"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
