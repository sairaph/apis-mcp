---
title: Delete a value list item
page_id: operation-delete-v1-radar-value-list-items-item-3d057103
path: operations/untagged
description: <p>Deletes a <code>ValueListItem</code> object, removing it from its parent value list.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/radar/value_list_items/{item}
operation_ids:
    - DeleteRadarValueListItemsItem
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Delete a value list item

`DELETE /v1/radar/value_list_items/{item}`

Operation ID: `DeleteRadarValueListItemsItem`

<p>Deletes a <code>ValueListItem</code> object, removing it from its parent value list.</p>

## Definition

```yaml
{"summary": "Delete a value list item", "description": "<p>Deletes a <code>ValueListItem</code> object, removing it from its parent value list.</p>", "operationId": "DeleteRadarValueListItemsItem", "parameters": [{"name": "item", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/deleted_radar.value_list_item"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
