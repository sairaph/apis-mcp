---
title: dlp_DataTagCategoryUpdate
page_id: schema-dlp-datatagcategoryupdate-c5b0343c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_DataTagCategoryUpdate

```yaml
{"type": "object", "properties": {"description": {"type": "string", "nullable": true}, "name": {"type": "string", "nullable": true}, "tags": {"description": "The desired final state of tags.\n- `None` (omitted): no tag changes.\n- `Some([])`: delete all tags.\n- `Some([...])`: desired final set + order.", "type": "array", "items": {"$ref": "#/components/schemas/dlp_DataTagCategoryTagUpdate"}, "nullable": true, "x-stainless-skip": ["terraform"]}}}
```
