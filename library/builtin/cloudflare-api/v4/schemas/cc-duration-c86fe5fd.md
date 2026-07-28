---
title: cc_Duration
page_id: schema-cc-duration-c86fe5fd
path: schemas
description: |-
    Duration string. From Go documentation:
      A string representing the duration in the form "3d1h3m". Leading zero units are omitted.
      As a special case, durations less than one second format use a smaller unit (milli-, micro-, or nanoseconds)
      to ensure that the leading digit is non-zero.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_Duration

Duration string. From Go documentation:
  A string representing the duration in the form "3d1h3m". Leading zero units are omitted.
  As a special case, durations less than one second format use a smaller unit (milli-, micro-, or nanoseconds)
  to ensure that the leading digit is non-zero.

```yaml
{"description": "Duration string. From Go documentation:\n  A string representing the duration in the form \"3d1h3m\". Leading zero units are omitted.\n  As a special case, durations less than one second format use a smaller unit (milli-, micro-, or nanoseconds)\n  to ensure that the leading digit is non-zero.\n", "type": "string", "example": "30s"}
```
