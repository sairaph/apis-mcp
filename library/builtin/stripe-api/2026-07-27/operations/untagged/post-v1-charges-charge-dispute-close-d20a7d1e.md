---
title: POST /v1/charges/{charge}/dispute/close
page_id: operation-post-v1-charges-charge-dispute-close-9cd906f2
path: operations/untagged
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/charges/{charge}/dispute/close
operation_ids:
    - PostChargesChargeDisputeClose
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# POST /v1/charges/{charge}/dispute/close

`POST /v1/charges/{charge}/dispute/close`

Operation ID: `PostChargesChargeDisputeClose`

## Definition

```yaml
{"description": "", "operationId": "PostChargesChargeDisputeClose", "parameters": [{"name": "charge", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dispute"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
