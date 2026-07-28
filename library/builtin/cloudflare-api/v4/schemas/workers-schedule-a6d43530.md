---
title: workers_schedule
page_id: schema-workers-schedule-a6d43530
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_schedule

```yaml
{"type": "object", "properties": {"created_on": {"type": "string", "readOnly": true, "x-auditable": true}, "cron": {"type": "string", "example": "*/30 * * * *", "x-auditable": true}, "modified_on": {"type": "string", "readOnly": true, "x-auditable": true}}, "required": ["cron"]}
```
