---
title: mcn_resource_details_section_item
page_id: schema-mcn-resource-details-section-item-369bbd7f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_resource_details_section_item

```yaml
{"type": "object", "properties": {"helpText": {"type": "string"}, "name": {"type": "string"}, "value": {"type": "object", "discriminator": {"propertyName": "item_type"}, "oneOf": [{"$ref": "#/components/schemas/mcn_string_item"}, {"$ref": "#/components/schemas/mcn_yaml_item"}, {"$ref": "#/components/schemas/mcn_yaml_diff_item"}, {"$ref": "#/components/schemas/mcn_resource_preview_item"}, {"$ref": "#/components/schemas/mcn_list_item"}]}}}
```
