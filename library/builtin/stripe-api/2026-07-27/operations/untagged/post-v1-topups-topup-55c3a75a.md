---
title: Update a top-up
page_id: operation-post-v1-topups-topup-3bc4bb6e
path: operations/untagged
description: <p>Updates the metadata of a top-up. Other top-up details are not editable by design.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/topups/{topup}
operation_ids:
    - PostTopupsTopup
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a top-up

`POST /v1/topups/{topup}`

Operation ID: `PostTopupsTopup`

<p>Updates the metadata of a top-up. Other top-up details are not editable by design.</p>

## Definition

```yaml
{"summary": "Update a top-up", "description": "<p>Updates the metadata of a top-up. Other top-up details are not editable by design.</p>", "operationId": "PostTopupsTopup", "parameters": [{"name": "topup", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the object. Often useful for displaying to users."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/topup"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
