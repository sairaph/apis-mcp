---
title: dlp_SensitivityGroupTemplate
page_id: schema-dlp-sensitivitygrouptemplate-2d7681c0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_SensitivityGroupTemplate

```yaml
{"type": "object", "properties": {"description": {"type": "string"}, "id": {"type": "string", "format": "uuid"}, "levels": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_SensitivityLevelTemplate"}}, "name": {"type": "string"}}, "required": ["id", "name", "description", "levels"]}
```
