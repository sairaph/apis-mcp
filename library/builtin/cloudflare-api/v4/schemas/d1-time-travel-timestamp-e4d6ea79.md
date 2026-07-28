---
title: d1_time-travel-timestamp
page_id: schema-d1-time-travel-timestamp-e4d6ea79
path: schemas
description: An ISO 8601 timestamp used for time travel operations. The database will be restored to the nearest available bookmark at or before this timestamp.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# d1_time-travel-timestamp

An ISO 8601 timestamp used for time travel operations. The database will be restored to the nearest available bookmark at or before this timestamp.

```yaml
{"description": "An ISO 8601 timestamp used for time travel operations. The database will be restored to the nearest available bookmark at or before this timestamp.", "type": "string", "format": "date-time", "example": "2024-01-15T12:00:00Z", "x-auditable": true}
```
