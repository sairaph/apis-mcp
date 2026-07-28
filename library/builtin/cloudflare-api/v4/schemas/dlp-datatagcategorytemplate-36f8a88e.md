---
title: dlp_DataTagCategoryTemplate
page_id: schema-dlp-datatagcategorytemplate-36f8a88e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_DataTagCategoryTemplate

```yaml
{"type": "object", "properties": {"description": {"type": "string"}, "id": {"type": "string", "format": "uuid"}, "name": {"type": "string"}, "tags": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_DataTagTemplate"}}}, "required": ["id", "name", "description", "tags"]}
```
