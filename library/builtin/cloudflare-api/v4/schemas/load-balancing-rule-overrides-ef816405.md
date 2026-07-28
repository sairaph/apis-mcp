---
title: load-balancing_rule_overrides
page_id: schema-load-balancing-rule-overrides-ef816405
path: schemas
description: A collection of overrides to apply when this rule's condition (or a pool set's `match`) is true. All fields are optional.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_rule_overrides

A collection of overrides to apply when this rule's condition (or a pool set's `match`) is true. All fields are optional.

```yaml
{"description": "A collection of overrides to apply when this rule's condition (or a pool set's `match`) is true. All fields are optional.", "type": "object", "properties": {"adaptive_routing": {"$ref": "#/components/schemas/load-balancing_adaptive_routing"}, "country_pools": {"$ref": "#/components/schemas/load-balancing_country_pools"}, "default_pools": {"$ref": "#/components/schemas/load-balancing_default_pools"}, "fallback_pool": {"$ref": "#/components/schemas/load-balancing_fallback_pool"}, "location_strategy": {"$ref": "#/components/schemas/load-balancing_location_strategy"}, "pool_default_weight": {"$ref": "#/components/schemas/load-balancing_pool_default_weight"}, "pool_weights": {"$ref": "#/components/schemas/load-balancing_pool_weights"}, "pools": {"$ref": "#/components/schemas/load-balancing_pools"}, "pop_pools": {"$ref": "#/components/schemas/load-balancing_pop_pools"}, "random_steering": {"$ref": "#/components/schemas/load-balancing_random_steering"}, "region_pools": {"$ref": "#/components/schemas/load-balancing_region_pools"}, "session_affinity": {"$ref": "#/components/schemas/load-balancing_session_affinity"}, "session_affinity_attributes": {"$ref": "#/components/schemas/load-balancing_session_affinity_attributes"}, "session_affinity_ttl": {"$ref": "#/components/schemas/load-balancing_session_affinity_ttl"}, "steering_policy": {"$ref": "#/components/schemas/load-balancing_steering_policy"}, "ttl": {"$ref": "#/components/schemas/load-balancing_ttl"}}, "x-stainless-terraform-configurability": "computed_optional"}
```
