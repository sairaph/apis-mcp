---
title: dlp_SensitivityGroupUpdate
page_id: schema-dlp-sensitivitygroupupdate-1e9f23e1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_SensitivityGroupUpdate

```yaml
{"type": "object", "properties": {"description": {"type": "string", "nullable": true}, "levels": {"description": "The desired final state of levels.\n- `None` (omitted): no level changes.\n- `Some([])`: delete all levels.\n- `Some([...])`: desired final set + order.", "type": "array", "items": {"$ref": "#/components/schemas/dlp_SensitivityGroupLevelUpdate"}, "nullable": true, "x-stainless-skip": ["terraform"]}, "name": {"type": "string", "nullable": true}}}
```
