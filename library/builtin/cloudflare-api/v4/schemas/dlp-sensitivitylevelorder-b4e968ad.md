---
title: dlp_SensitivityLevelOrder
page_id: schema-dlp-sensitivitylevelorder-b4e968ad
path: schemas
description: |-
    The ordered list of level IDs for a sensitivity group.
    Used to get and set the ordering of levels independently of level attributes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_SensitivityLevelOrder

The ordered list of level IDs for a sensitivity group.
Used to get and set the ordering of levels independently of level attributes.

```yaml
{"description": "The ordered list of level IDs for a sensitivity group.\nUsed to get and set the ordering of levels independently of level attributes.", "type": "object", "properties": {"level_ids": {"type": "array", "items": {"format": "uuid", "type": "string"}}}, "required": ["level_ids"]}
```
