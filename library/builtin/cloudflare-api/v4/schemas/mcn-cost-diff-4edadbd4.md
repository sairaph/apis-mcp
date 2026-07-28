---
title: mcn_cost_diff
page_id: schema-mcn-cost-diff-4edadbd4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_cost_diff

```yaml
{"type": "object", "properties": {"currency": {"type": "string", "x-auditable": true}, "current_monthly_cost": {"type": "number", "format": "double"}, "diff": {"type": "number", "format": "double"}, "proposed_monthly_cost": {"type": "number", "format": "double"}}, "required": ["current_monthly_cost", "proposed_monthly_cost", "diff", "currency"]}
```
