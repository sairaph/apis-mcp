---
title: load-balancing_pool_sets
page_id: schema-load-balancing-pool-sets-3cc41da3
path: schemas
description: 'An optional list of pool sets, evaluated in array order with first match wins. Pool sets are independent from the standard steering fields (`region_pools` / `country_pools` / `pop_pools` / `default_pools` / `steering_policy` / `random_steering` / `fallback_pool` / `rules`). On a PATCH, an empty array (`pool_sets: []`) clears all pool sets, while omitting the field leaves existing pool sets unchanged.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_pool_sets

An optional list of pool sets, evaluated in array order with first match wins. Pool sets are independent from the standard steering fields (`region_pools` / `country_pools` / `pop_pools` / `default_pools` / `steering_policy` / `random_steering` / `fallback_pool` / `rules`). On a PATCH, an empty array (`pool_sets: []`) clears all pool sets, while omitting the field leaves existing pool sets unchanged.

```yaml
{"description": "An optional list of pool sets, evaluated in array order with first match wins. Pool sets are independent from the standard steering fields (`region_pools` / `country_pools` / `pop_pools` / `default_pools` / `steering_policy` / `random_steering` / `fallback_pool` / `rules`). On a PATCH, an empty array (`pool_sets: []`) clears all pool sets, while omitting the field leaves existing pool sets unchanged.", "type": "array", "items": {"$ref": "#/components/schemas/load-balancing_pool_set"}, "x-stainless-terraform-configurability": "computed_optional"}
```
