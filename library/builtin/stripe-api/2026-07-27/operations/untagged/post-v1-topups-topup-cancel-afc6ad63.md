---
title: Cancel a top-up
page_id: operation-post-v1-topups-topup-cancel-85812a44
path: operations/untagged
description: <p>Cancels a top-up. Only pending top-ups can be canceled.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/topups/{topup}/cancel
operation_ids:
    - PostTopupsTopupCancel
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Cancel a top-up

`POST /v1/topups/{topup}/cancel`

Operation ID: `PostTopupsTopupCancel`

<p>Cancels a top-up. Only pending top-ups can be canceled.</p>

## Definition

```yaml
{"summary": "Cancel a top-up", "description": "<p>Cancels a top-up. Only pending top-ups can be canceled.</p>", "operationId": "PostTopupsTopupCancel", "parameters": [{"name": "topup", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/topup"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
