---
title: r2_lifecycle-date-condition
page_id: schema-r2-lifecycle-date-condition-19d50431
path: schemas
description: Condition for lifecycle transitions to apply on a specific date.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_lifecycle-date-condition

Condition for lifecycle transitions to apply on a specific date.

```yaml
{"description": "Condition for lifecycle transitions to apply on a specific date.", "type": "object", "properties": {"date": {"type": "string", "format": "date-time", "x-auditable": true}, "type": {"type": "string", "enum": ["Date"], "x-auditable": true}}, "required": ["type", "date"]}
```
