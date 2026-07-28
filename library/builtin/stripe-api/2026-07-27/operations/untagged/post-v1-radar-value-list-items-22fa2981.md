---
title: Create a value list item
page_id: operation-post-v1-radar-value-list-items-ca3df51d
path: operations/untagged
description: <p>Creates a new <code>ValueListItem</code> object, which is added to the specified parent value list.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/radar/value_list_items
operation_ids:
    - PostRadarValueListItems
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a value list item

`POST /v1/radar/value_list_items`

Operation ID: `PostRadarValueListItems`

<p>Creates a new <code>ValueListItem</code> object, which is added to the specified parent value list.</p>

## Definition

```yaml
{"summary": "Create a value list item", "description": "<p>Creates a new <code>ValueListItem</code> object, which is added to the specified parent value list.</p>", "operationId": "PostRadarValueListItems", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["value", "value_list"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "value": {"maxLength": 800, "type": "string", "description": "The value of the item (whose type must match the type of the parent value list)."}, "value_list": {"maxLength": 5000, "type": "string", "description": "The identifier of the value list which the created item will be added to."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/radar.value_list_item"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
