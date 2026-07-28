---
title: dlp_NewSensitivityGroup
page_id: schema-dlp-newsensitivitygroup-9abd9164
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_NewSensitivityGroup

```yaml
{"type": "object", "properties": {"description": {"type": "string", "nullable": true}, "levels": {"description": "Levels to create with the group. Mutually exclusive with `template_id`.", "type": "array", "items": {"$ref": "#/components/schemas/dlp_NewSensitivityLevel"}, "x-stainless-skip": ["terraform"]}, "name": {"type": "string"}, "template_id": {"type": "string", "format": "uuid", "nullable": true, "x-stainless-skip": ["terraform"]}}, "required": ["name"]}
```
