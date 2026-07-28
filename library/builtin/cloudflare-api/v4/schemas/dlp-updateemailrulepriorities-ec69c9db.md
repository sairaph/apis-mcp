---
title: dlp_UpdateEmailRulePriorities
page_id: schema-dlp-updateemailrulepriorities-ec69c9db
path: schemas
description: |-
    Used to update multiple email rule priorities as an atomic action,
    to support patterns such as swapping the priorities of two email rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_UpdateEmailRulePriorities

Used to update multiple email rule priorities as an atomic action,
to support patterns such as swapping the priorities of two email rules.

```yaml
{"description": "Used to update multiple email rule priorities as an atomic action,\nto support patterns such as swapping the priorities of two email rules.", "type": "object", "properties": {"new_priorities": {"type": "object", "additionalProperties": {"format": "int32", "minimum": 0, "type": "integer"}}}, "required": ["new_priorities"]}
```
