---
title: workers-observability_filter_node
page_id: schema-workers-observability-filter-node-b11b7e93
path: schemas
description: 'Supports nested groups via kind: ''group''.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers-observability_filter_node

Supports nested groups via kind: 'group'.

```yaml
{"description": "Supports nested groups via kind: 'group'.", "type": "object", "anyOf": [{"properties": {"filterCombination": {"type": "string", "enum": ["and", "or", "AND", "OR"]}, "filters": {"type": "array", "items": {"$ref": "#/components/schemas/workers-observability_filter_node"}, "minItems": 1}, "kind": {"type": "string", "enum": ["group"]}}, "required": ["kind", "filterCombination", "filters"], "type": "object"}, {"$ref": "#/components/schemas/workers-observability_filter_leaf"}]}
```
