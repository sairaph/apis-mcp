---
title: radar.value_list_item
page_id: schema-radar-value-list-item-d4e6ad7b
path: schemas
description: |-
    Value list items allow you to add specific values to a given Radar value list, which can then be used in rules.

    Related guide: [Managing list items](https://docs.stripe.com/radar/lists#managing-list-items)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# radar.value_list_item

Value list items allow you to add specific values to a given Radar value list, which can then be used in rules.

Related guide: [Managing list items](https://docs.stripe.com/radar/lists#managing-list-items)

```yaml
{"title": "RadarListListItem", "required": ["created", "created_by", "id", "livemode", "object", "value", "value_list"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "created_by": {"maxLength": 5000, "type": "string", "description": "The name or email address of the user who added this item to the value list."}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["radar.value_list_item"]}, "value": {"maxLength": 5000, "type": "string", "description": "The value of the item."}, "value_list": {"maxLength": 5000, "type": "string", "description": "The identifier of the value list this item belongs to."}}, "description": "Value list items allow you to add specific values to a given Radar value list, which can then be used in rules.\n\nRelated guide: [Managing list items](https://docs.stripe.com/radar/lists#managing-list-items)", "x-expandableFields": [], "x-resourceId": "radar.value_list_item"}
```
