---
title: dlp_SensitivityGroupLevelUpdate
page_id: schema-dlp-sensitivitygrouplevelupdate-5fca0dc7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_SensitivityGroupLevelUpdate

```yaml
{"allOf": [{"$ref": "#/components/schemas/dlp_SensitivityLevelUpdate"}, {"properties": {"id": {"description": "If `None` (omitted), a new level will be created. Otherwise, an existing level will\nbe updated.", "type": "string", "format": "uuid", "nullable": true}}, "type": "object"}]}
```
