---
title: Immediately invalidate an ephemeral key
page_id: operation-delete-v1-ephemeral-keys-key-0b5d6d83
path: operations/untagged
description: <p>Invalidates a short-lived API key for a given resource.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/ephemeral_keys/{key}
operation_ids:
    - DeleteEphemeralKeysKey
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Immediately invalidate an ephemeral key

`DELETE /v1/ephemeral_keys/{key}`

Operation ID: `DeleteEphemeralKeysKey`

<p>Invalidates a short-lived API key for a given resource.</p>

## Definition

```yaml
{"summary": "Immediately invalidate an ephemeral key", "description": "<p>Invalidates a short-lived API key for a given resource.</p>", "operationId": "DeleteEphemeralKeysKey", "parameters": [{"name": "key", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ephemeral_key"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
