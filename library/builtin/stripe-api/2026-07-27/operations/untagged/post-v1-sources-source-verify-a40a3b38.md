---
title: POST /v1/sources/{source}/verify
page_id: operation-post-v1-sources-source-verify-31f4f428
path: operations/untagged
description: <p>Verify a given source.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/sources/{source}/verify
operation_ids:
    - PostSourcesSourceVerify
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# POST /v1/sources/{source}/verify

`POST /v1/sources/{source}/verify`

Operation ID: `PostSourcesSourceVerify`

<p>Verify a given source.</p>

## Definition

```yaml
{"description": "<p>Verify a given source.</p>", "operationId": "PostSourcesSourceVerify", "parameters": [{"name": "source", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["values"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "values": {"type": "array", "description": "The values needed to verify the source.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "values": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/source"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
