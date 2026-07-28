---
title: Update a file link
page_id: operation-post-v1-file-links-link-5fc70133
path: operations/untagged
description: <p>Updates an existing file link object. Expired links can no longer be updated.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/file_links/{link}
operation_ids:
    - PostFileLinksLink
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a file link

`POST /v1/file_links/{link}`

Operation ID: `PostFileLinksLink`

<p>Updates an existing file link object. Expired links can no longer be updated.</p>

## Definition

```yaml
{"summary": "Update a file link", "description": "<p>Updates an existing file link object. Expired links can no longer be updated.</p>", "operationId": "PostFileLinksLink", "parameters": [{"name": "link", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "expires_at": {"description": "A future timestamp after which the link will no longer be usable, or `now` to expire the link immediately.", "anyOf": [{"maxLength": 5000, "type": "string", "enum": ["now"]}, {"type": "integer", "format": "unix-time"}, {"type": "string", "enum": [""]}]}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "expires_at": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/file_link"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
