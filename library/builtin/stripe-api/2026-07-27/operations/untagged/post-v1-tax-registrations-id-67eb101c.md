---
title: Update a registration
page_id: operation-post-v1-tax-registrations-id-caffa199
path: operations/untagged
description: |-
    <p>Updates an existing Tax <code>Registration</code> object.</p>

    <p>A registration cannot be deleted after it has been created. If you wish to end a registration you may do so by setting <code>expires_at</code>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/tax/registrations/{id}
operation_ids:
    - PostTaxRegistrationsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a registration

`POST /v1/tax/registrations/{id}`

Operation ID: `PostTaxRegistrationsId`

<p>Updates an existing Tax <code>Registration</code> object.</p>

<p>A registration cannot be deleted after it has been created. If you wish to end a registration you may do so by setting <code>expires_at</code>.</p>

## Definition

```yaml
{"summary": "Update a registration", "description": "<p>Updates an existing Tax <code>Registration</code> object.</p>\n\n<p>A registration cannot be deleted after it has been created. If you wish to end a registration you may do so by setting <code>expires_at</code>.</p>", "operationId": "PostTaxRegistrationsId", "parameters": [{"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"active_from": {"description": "Time at which the registration becomes active. It can be either `now` to indicate the current time, or a timestamp measured in seconds since the Unix epoch.", "anyOf": [{"maxLength": 5000, "type": "string", "enum": ["now"]}, {"type": "integer", "format": "unix-time"}]}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "expires_at": {"description": "If set, the registration stops being active at this time. If not set, the registration will be active indefinitely. It can be either `now` to indicate the current time, or a timestamp measured in seconds since the Unix epoch.", "anyOf": [{"maxLength": 5000, "type": "string", "enum": ["now"]}, {"type": "integer", "format": "unix-time"}, {"type": "string", "enum": [""]}]}}, "additionalProperties": false}, "encoding": {"active_from": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "expires_at": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tax.registration"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
