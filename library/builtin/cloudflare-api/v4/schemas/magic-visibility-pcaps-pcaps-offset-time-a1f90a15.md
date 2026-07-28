---
title: magic-visibility-pcaps_pcaps_offset_time
page_id: schema-magic-visibility-pcaps-pcaps-offset-time-a1f90a15
path: schemas
description: The RFC 3339 offset timestamp from which to query backwards for packets. Must be within the last 24h. When this field is empty, defaults to time of request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic-visibility-pcaps_pcaps_offset_time

The RFC 3339 offset timestamp from which to query backwards for packets. Must be within the last 24h. When this field is empty, defaults to time of request.

```yaml
{"description": "The RFC 3339 offset timestamp from which to query backwards for packets. Must be within the last 24h. When this field is empty, defaults to time of request.", "type": "string", "format": "date-time", "example": "2020-01-01T08:00:00Z", "x-auditable": true}
```
