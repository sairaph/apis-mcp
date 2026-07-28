---
title: Retrieve a value list
page_id: operation-get-v1-radar-value-lists-value-list-b969f0b7
path: operations/untagged
description: <p>Retrieves a <code>ValueList</code> object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/radar/value_lists/{value_list}
operation_ids:
    - GetRadarValueListsValueList
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a value list

`GET /v1/radar/value_lists/{value_list}`

Operation ID: `GetRadarValueListsValueList`

<p>Retrieves a <code>ValueList</code> object.</p>

## Definition

```yaml
{"summary": "Retrieve a value list", "description": "<p>Retrieves a <code>ValueList</code> object.</p>", "operationId": "GetRadarValueListsValueList", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "value_list", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/radar.value_list"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
