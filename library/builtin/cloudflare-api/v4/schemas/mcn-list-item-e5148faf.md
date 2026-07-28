---
title: mcn_list_item
page_id: schema-mcn-list-item-e5148faf
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_list_item

```yaml
{"type": "object", "properties": {"item_type": {"type": "string"}, "list": {"type": "array", "items": {"discriminator": {"propertyName": "item_type"}, "oneOf": [{"$ref": "#/components/schemas/mcn_string_item"}, {"$ref": "#/components/schemas/mcn_resource_preview_item"}], "type": "object"}}}, "required": ["item_type", "list"]}
```
