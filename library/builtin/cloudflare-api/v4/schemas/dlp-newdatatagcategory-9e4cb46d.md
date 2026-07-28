---
title: dlp_NewDataTagCategory
page_id: schema-dlp-newdatatagcategory-9e4cb46d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_NewDataTagCategory

```yaml
{"type": "object", "properties": {"description": {"type": "string", "nullable": true}, "name": {"type": "string"}, "tags": {"description": "Tags to create with the category. Mutually exclusive with `template_id`.", "type": "array", "items": {"$ref": "#/components/schemas/dlp_NewDataTag"}, "x-stainless-skip": ["terraform"]}, "template_id": {"type": "string", "format": "uuid", "nullable": true, "x-stainless-skip": ["terraform"]}}, "required": ["name"]}
```
