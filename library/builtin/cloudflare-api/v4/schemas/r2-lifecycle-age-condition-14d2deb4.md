---
title: r2_lifecycle-age-condition
page_id: schema-r2-lifecycle-age-condition-14d2deb4
path: schemas
description: Condition for lifecycle transitions to apply after an object reaches an age in seconds.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_lifecycle-age-condition

Condition for lifecycle transitions to apply after an object reaches an age in seconds.

```yaml
{"description": "Condition for lifecycle transitions to apply after an object reaches an age in seconds.", "type": "object", "properties": {"maxAge": {"type": "integer", "x-auditable": true}, "type": {"type": "string", "enum": ["Age"], "x-auditable": true}}, "required": ["type", "maxAge"]}
```
