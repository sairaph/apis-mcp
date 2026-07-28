---
title: load-balancing_load-balancer
page_id: schema-load-balancing-load-balancer-b04f7d9e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_load-balancer

```yaml
{"type": "object", "properties": {"adaptive_routing": {"$ref": "#/components/schemas/load-balancing_adaptive_routing"}, "country_pools": {"$ref": "#/components/schemas/load-balancing_country_pools"}, "created_on": {"$ref": "#/components/schemas/load-balancing_timestamp"}, "default_pools": {"$ref": "#/components/schemas/load-balancing_default_pools"}, "description": {"$ref": "#/components/schemas/load-balancing_components-schemas-description"}, "enabled": {"$ref": "#/components/schemas/load-balancing_components-schemas-enabled"}, "fallback_pool": {"$ref": "#/components/schemas/load-balancing_fallback_pool"}, "id": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-identifier"}, "location_strategy": {"$ref": "#/components/schemas/load-balancing_location_strategy"}, "modified_on": {"$ref": "#/components/schemas/load-balancing_timestamp"}, "name": {"$ref": "#/components/schemas/load-balancing_components-schemas-name"}, "networks": {"$ref": "#/components/schemas/load-balancing_networks"}, "pool_sets": {"$ref": "#/components/schemas/load-balancing_pool_sets"}, "pop_pools": {"$ref": "#/components/schemas/load-balancing_pop_pools"}, "proxied": {"$ref": "#/components/schemas/load-balancing_proxied"}, "random_steering": {"$ref": "#/components/schemas/load-balancing_random_steering"}, "region_pools": {"$ref": "#/components/schemas/load-balancing_region_pools"}, "rules": {"$ref": "#/components/schemas/load-balancing_rules"}, "session_affinity": {"$ref": "#/components/schemas/load-balancing_session_affinity"}, "session_affinity_attributes": {"$ref": "#/components/schemas/load-balancing_session_affinity_attributes"}, "session_affinity_ttl": {"$ref": "#/components/schemas/load-balancing_session_affinity_ttl"}, "steering_policy": {"$ref": "#/components/schemas/load-balancing_steering_policy"}, "ttl": {"$ref": "#/components/schemas/load-balancing_ttl"}, "zone_name": {"$ref": "#/components/schemas/load-balancing_components-schemas-zone-name"}}}
```
