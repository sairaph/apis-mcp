---
title: load-balancing_pool_set
page_id: schema-load-balancing-pool-set-c3b3c508
path: schemas
description: One entry in a load balancer's `pool_sets`. Pool sets are evaluated in array order; the first whose `match` succeeds applies its `overrides` (or `fixed_response`), and evaluation stops there.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_pool_set

One entry in a load balancer's `pool_sets`. Pool sets are evaluated in array order; the first whose `match` succeeds applies its `overrides` (or `fixed_response`), and evaluation stops there.

```yaml
{"description": "One entry in a load balancer's `pool_sets`. Pool sets are evaluated in array order; the first whose `match` succeeds applies its `overrides` (or `fixed_response`), and evaluation stops there.", "type": "object", "properties": {"disabled": {"description": "Disable this specific pool set. It will no longer be evaluated.", "type": "boolean", "default": false, "x-auditable": true}, "fixed_response": {"$ref": "#/components/schemas/load-balancing_fixed_response"}, "match": {"$ref": "#/components/schemas/load-balancing_match"}, "name": {"description": "A human-readable name for this pool set.", "type": "string", "example": "wnam-primary", "maxLength": 200, "x-auditable": true}, "overrides": {"$ref": "#/components/schemas/load-balancing_pool_set_overrides"}}, "example": {"match": {"topology": {"regions": ["WNAM"]}}, "name": "wnam-primary", "overrides": {"fallback_pool": "9290f38c5d07c2e2f4df57b1f61d4196", "pools": ["17b5962d775c646f3f9725cbc7a53df4"], "steering_policy": "random"}}, "x-stainless-terraform-configurability": "computed_optional"}
```
