---
title: r2_lock-rule-age-condition
page_id: schema-r2-lock-rule-age-condition-0e9237db
path: schemas
description: Condition to apply a lock rule to an object for how long in seconds.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_lock-rule-age-condition

Condition to apply a lock rule to an object for how long in seconds.

```yaml
{"description": "Condition to apply a lock rule to an object for how long in seconds.", "type": "object", "properties": {"maxAgeSeconds": {"type": "integer", "example": 100, "x-auditable": true}, "type": {"type": "string", "enum": ["Age"], "x-auditable": true}}, "required": ["type", "maxAgeSeconds"]}
```
