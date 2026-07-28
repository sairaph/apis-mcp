---
title: dlp_DataClass
page_id: schema-dlp-dataclass-152fd086
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_DataClass

```yaml
{"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "data_tags": {"type": "array", "items": {"format": "uuid", "type": "string"}}, "description": {"type": "string", "nullable": true}, "expression": {"type": "string"}, "id": {"type": "string", "format": "uuid"}, "name": {"type": "string"}, "sensitivity_levels": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_SensitivityLevelRef"}}, "updated_at": {"type": "string", "format": "date-time"}}, "required": ["id", "name", "expression", "sensitivity_levels", "data_tags", "created_at", "updated_at"]}
```
