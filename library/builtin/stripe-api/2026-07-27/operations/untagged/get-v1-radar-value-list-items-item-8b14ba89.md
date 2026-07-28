---
title: Retrieve a value list item
page_id: operation-get-v1-radar-value-list-items-item-24ea9d27
path: operations/untagged
description: <p>Retrieves a <code>ValueListItem</code> object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/radar/value_list_items/{item}
operation_ids:
    - GetRadarValueListItemsItem
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a value list item

`GET /v1/radar/value_list_items/{item}`

Operation ID: `GetRadarValueListItemsItem`

<p>Retrieves a <code>ValueListItem</code> object.</p>

## Definition

```yaml
{"summary": "Retrieve a value list item", "description": "<p>Retrieves a <code>ValueListItem</code> object.</p>", "operationId": "GetRadarValueListItemsItem", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "item", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/radar.value_list_item"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
