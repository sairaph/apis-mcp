---
title: r2_lock-rule-date-condition
page_id: schema-r2-lock-rule-date-condition-97779b16
path: schemas
description: Condition to apply a lock rule to an object until a specific date.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_lock-rule-date-condition

Condition to apply a lock rule to an object until a specific date.

```yaml
{"description": "Condition to apply a lock rule to an object until a specific date.", "type": "object", "properties": {"date": {"type": "string", "format": "date-time", "x-auditable": true}, "type": {"type": "string", "enum": ["Date"], "x-auditable": true}}, "required": ["type", "date"]}
```
