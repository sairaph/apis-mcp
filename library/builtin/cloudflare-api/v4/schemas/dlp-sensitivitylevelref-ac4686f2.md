---
title: dlp_SensitivityLevelRef
page_id: schema-dlp-sensitivitylevelref-ac4686f2
path: schemas
description: A reference pairing a sensitivity group with a specific level within that group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_SensitivityLevelRef

A reference pairing a sensitivity group with a specific level within that group.

```yaml
{"description": "A reference pairing a sensitivity group with a specific level within that group.", "type": "object", "properties": {"group_id": {"type": "string", "format": "uuid"}, "level_id": {"type": "string", "format": "uuid"}}, "required": ["group_id", "level_id"]}
```
