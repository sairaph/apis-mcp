---
title: Update a refund
page_id: operation-post-v1-refunds-refund-ea8dbf9d
path: operations/untagged
description: |-
    <p>Updates the refund that you specify by setting the values of the passed parameters. Any parameters that you don’t provide remain unchanged.</p>

    <p>This request only accepts <code>metadata</code> as an argument.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/refunds/{refund}
operation_ids:
    - PostRefundsRefund
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a refund

`POST /v1/refunds/{refund}`

Operation ID: `PostRefundsRefund`

<p>Updates the refund that you specify by setting the values of the passed parameters. Any parameters that you don’t provide remain unchanged.</p>

<p>This request only accepts <code>metadata</code> as an argument.</p>

## Definition

```yaml
{"summary": "Update a refund", "description": "<p>Updates the refund that you specify by setting the values of the passed parameters. Any parameters that you don’t provide remain unchanged.</p>\n\n<p>This request only accepts <code>metadata</code> as an argument.</p>", "operationId": "PostRefundsRefund", "parameters": [{"name": "refund", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/refund"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
