---
title: Reactivate a billing meter
page_id: operation-post-v1-billing-meters-id-reactivate-d7c4bfbe
path: operations/untagged
description: <p>When a meter is reactivated, events for this meter can be accepted and you can attach the meter to a price.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/billing/meters/{id}/reactivate
operation_ids:
    - PostBillingMetersIdReactivate
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Reactivate a billing meter

`POST /v1/billing/meters/{id}/reactivate`

Operation ID: `PostBillingMetersIdReactivate`

<p>When a meter is reactivated, events for this meter can be accepted and you can attach the meter to a price.</p>

## Definition

```yaml
{"summary": "Reactivate a billing meter", "description": "<p>When a meter is reactivated, events for this meter can be accepted and you can attach the meter to a price.</p>", "operationId": "PostBillingMetersIdReactivate", "parameters": [{"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billing.meter"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
