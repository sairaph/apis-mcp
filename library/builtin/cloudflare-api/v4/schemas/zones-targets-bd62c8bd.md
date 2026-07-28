---
title: zones_targets
page_id: schema-zones-targets-bd62c8bd
path: schemas
description: The rule targets to evaluate on each request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_targets

The rule targets to evaluate on each request.

```yaml
{"description": "The rule targets to evaluate on each request.", "type": "array", "items": {"$ref": "#/components/schemas/zones_target"}, "example": [{"constraint": {"operator": "matches", "value": "*example.com/images/*"}, "target": "url"}], "x-stainless-skip": ["terraform"]}
```
