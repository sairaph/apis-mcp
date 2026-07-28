---
title: zones_string_constraint
page_id: schema-zones-string-constraint-545d4c7e
path: schemas
description: String constraint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_string_constraint

String constraint.

```yaml
{"description": "String constraint.", "type": "object", "properties": {"operator": {"description": "The matches operator can use asterisks and pipes as wildcard and 'or' operators.", "default": "contains", "enum": ["matches", "contains", "equals", "not_equal", "not_contain"], "x-auditable": true}, "value": {"description": "The value to apply the operator to.", "type": "string", "x-auditable": true}}, "required": ["operator", "value"]}
```
