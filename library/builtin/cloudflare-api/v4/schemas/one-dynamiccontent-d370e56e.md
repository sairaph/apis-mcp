---
title: one_DynamicContent
page_id: schema-one-dynamiccontent-d370e56e
path: schemas
description: Dynamic content for instruction/form_input steps.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# one_DynamicContent

Dynamic content for instruction/form_input steps.

```yaml
{"description": "Dynamic content for instruction/form_input steps.", "type": "object", "properties": {"label": {"description": "Display label.", "type": "string"}, "type": {"description": "Content type.\n\n* `copy_block` - copy_block\n* `external_link` - external_link", "type": "string", "enum": ["copy_block", "external_link"]}, "url_template": {"description": "URL template with {{ variable }} interpolation (for external_link).", "type": "string", "nullable": true}, "value_from": {"description": "Field path to get value from (for copy_block).", "type": "string", "nullable": true}}, "required": ["label", "type"]}
```
