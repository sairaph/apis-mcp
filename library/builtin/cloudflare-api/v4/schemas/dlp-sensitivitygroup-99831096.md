---
title: dlp_SensitivityGroup
page_id: schema-dlp-sensitivitygroup-99831096
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_SensitivityGroup

```yaml
{"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "description": {"type": "string", "nullable": true}, "id": {"type": "string", "format": "uuid"}, "levels": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_SensitivityLevel"}}, "name": {"type": "string"}, "template_id": {"type": "string", "format": "uuid", "nullable": true}, "updated_at": {"type": "string", "format": "date-time"}}, "required": ["id", "name", "created_at", "updated_at", "levels"]}
```
