---
title: load-balancing_pools
page_id: schema-load-balancing-pools-39803496
path: schemas
description: A flat, ordered list of pool IDs to route the matched audience to. Replaces the resolved topology with exactly these pools. Mutually exclusive with `fixed_response`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_pools

A flat, ordered list of pool IDs to route the matched audience to. Replaces the resolved topology with exactly these pools. Mutually exclusive with `fixed_response`.

```yaml
{"description": "A flat, ordered list of pool IDs to route the matched audience to. Replaces the resolved topology with exactly these pools. Mutually exclusive with `fixed_response`.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["17b5962d775c646f3f9725cbc7a53df4"], "x-stainless-terraform-configurability": "computed_optional"}
```
