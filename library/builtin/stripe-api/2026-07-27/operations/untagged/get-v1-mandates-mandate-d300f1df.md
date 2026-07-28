---
title: Retrieve a Mandate
page_id: operation-get-v1-mandates-mandate-820f0cce
path: operations/untagged
description: <p>Retrieves a Mandate object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/mandates/{mandate}
operation_ids:
    - GetMandatesMandate
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a Mandate

`GET /v1/mandates/{mandate}`

Operation ID: `GetMandatesMandate`

<p>Retrieves a Mandate object.</p>

## Definition

```yaml
{"summary": "Retrieve a Mandate", "description": "<p>Retrieves a Mandate object.</p>", "operationId": "GetMandatesMandate", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "mandate", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mandate"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
