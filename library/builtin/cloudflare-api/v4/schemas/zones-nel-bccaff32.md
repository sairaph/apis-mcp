---
title: zones_nel
page_id: schema-zones-nel-bccaff32
path: schemas
description: Enable Network Error Logging reporting on your zone. (Beta)
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_nel

Enable Network Error Logging reporting on your zone. (Beta)

```yaml
{"description": "Enable Network Error Logging reporting on your zone. (Beta) ", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "Zone setting identifier.", "example": "nel", "enum": ["nel"]}, "value": {"$ref": "#/components/schemas/zones_nel_value"}}}], "title": "Network Error Logging"}
```
