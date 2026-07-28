---
title: dlp_NewDataClass
page_id: schema-dlp-newdataclass-01dca00c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_NewDataClass

```yaml
{"type": "object", "properties": {"data_tags": {"type": "array", "items": {"format": "uuid", "type": "string"}}, "description": {"type": "string", "nullable": true}, "expression": {"type": "string"}, "name": {"type": "string"}, "sensitivity_levels": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_SensitivityLevelRef"}}}, "required": ["name", "expression", "sensitivity_levels", "data_tags"]}
```
