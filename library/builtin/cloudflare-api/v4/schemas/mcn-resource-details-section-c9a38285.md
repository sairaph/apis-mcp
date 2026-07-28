---
title: mcn_resource_details_section
page_id: schema-mcn-resource-details-section-c9a38285
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_resource_details_section

```yaml
{"type": "object", "properties": {"help_text": {"type": "string"}, "hidden_items": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_resource_details_section_item"}}, "name": {"type": "string"}, "visible_items": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_resource_details_section_item"}}}, "required": ["name", "visible_items", "hidden_items"]}
```
