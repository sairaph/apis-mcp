---
title: aaa_history-status
page_id: schema-aaa-history-status-f26ab2b3
path: schemas
description: |-
    Indicates the quality of the resource identification used to derive the history.
    - `exact`: Resource was identified by the resource URI.
    - `approximate`: Resource was identified without the resource URI.
    - `unavailable`: The source audit log entry did not contain enough information to identify the resource; result is empty.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_history-status

Indicates the quality of the resource identification used to derive the history.
- `exact`: Resource was identified by the resource URI.
- `approximate`: Resource was identified without the resource URI.
- `unavailable`: The source audit log entry did not contain enough information to identify the resource; result is empty.

```yaml
{"description": "Indicates the quality of the resource identification used to derive the history.\n- `exact`: Resource was identified by the resource URI.\n- `approximate`: Resource was identified without the resource URI.\n- `unavailable`: The source audit log entry did not contain enough information to identify the resource; result is empty.\n", "type": "string", "example": "exact", "enum": ["exact", "approximate", "unavailable"]}
```
