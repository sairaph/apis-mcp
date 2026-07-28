---
title: mcn_resource_diff
page_id: schema-mcn-resource-diff-cdacbfcb
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_resource_diff

```yaml
{"type": "object", "properties": {"diff": {"$ref": "#/components/schemas/mcn_yaml_diff"}, "keys_require_replace": {"type": "array", "items": {"type": "string"}}, "monthly_cost_estimate_diff": {"$ref": "#/components/schemas/mcn_cost_diff"}, "planned_action": {"$ref": "#/components/schemas/mcn_planned_action"}, "resource": {"$ref": "#/components/schemas/mcn_resource_preview"}}, "required": ["resource", "diff", "planned_action", "keys_require_replace", "monthly_cost_estimate_diff"]}
```
