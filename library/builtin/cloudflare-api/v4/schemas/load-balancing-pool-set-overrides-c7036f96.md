---
title: load-balancing_pool_set_overrides
page_id: schema-load-balancing-pool-set-overrides-c7036f96
path: schemas
description: 'The behavior a pool set applies when its `match` succeeds. A strict subset of a rule''s `overrides`: a pool set replaces the topology wholesale with a flat pool list (`pools`), so only the declarative pool-routing fields plus `fallback_pool` and `steering_policy` are settable. All fields are optional.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_pool_set_overrides

The behavior a pool set applies when its `match` succeeds. A strict subset of a rule's `overrides`: a pool set replaces the topology wholesale with a flat pool list (`pools`), so only the declarative pool-routing fields plus `fallback_pool` and `steering_policy` are settable. All fields are optional.

```yaml
{"description": "The behavior a pool set applies when its `match` succeeds. A strict subset of a rule's `overrides`: a pool set replaces the topology wholesale with a flat pool list (`pools`), so only the declarative pool-routing fields plus `fallback_pool` and `steering_policy` are settable. All fields are optional.", "type": "object", "properties": {"fallback_pool": {"$ref": "#/components/schemas/load-balancing_fallback_pool"}, "pool_default_weight": {"$ref": "#/components/schemas/load-balancing_pool_default_weight"}, "pool_weights": {"$ref": "#/components/schemas/load-balancing_pool_weights"}, "pools": {"$ref": "#/components/schemas/load-balancing_pools"}, "steering_policy": {"$ref": "#/components/schemas/load-balancing_steering_policy"}}, "x-stainless-terraform-configurability": "computed_optional"}
```
