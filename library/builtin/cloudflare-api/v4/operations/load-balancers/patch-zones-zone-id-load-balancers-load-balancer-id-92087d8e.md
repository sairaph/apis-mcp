---
title: Patch Load Balancer
page_id: operation-patch-zones-zone-id-load-balancers-load-balancer-id-0a6f326d
path: operations/load-balancers
description: Apply changes to an existing load balancer, overwriting the supplied properties.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/load_balancers/{load_balancer_id}
operation_ids:
    - load-balancers-patch-load-balancer
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Load Balancer

`PATCH /zones/{zone_id}/load_balancers/{load_balancer_id}`

Operation ID: `load-balancers-patch-load-balancer`

Apply changes to an existing load balancer, overwriting the supplied properties.

## Definition

```yaml
{"operationId": "load-balancers-patch-load-balancer", "summary": "Patch Load Balancer", "description": "Apply changes to an existing load balancer, overwriting the supplied properties.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-identifier"}}, {"name": "load_balancer_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"adaptive_routing": {"$ref": "#/components/schemas/load-balancing_adaptive_routing"}, "country_pools": {"$ref": "#/components/schemas/load-balancing_country_pools"}, "default_pools": {"$ref": "#/components/schemas/load-balancing_default_pools"}, "description": {"$ref": "#/components/schemas/load-balancing_components-schemas-description"}, "enabled": {"$ref": "#/components/schemas/load-balancing_components-schemas-enabled"}, "fallback_pool": {"$ref": "#/components/schemas/load-balancing_fallback_pool"}, "location_strategy": {"$ref": "#/components/schemas/load-balancing_location_strategy"}, "name": {"$ref": "#/components/schemas/load-balancing_components-schemas-name"}, "pop_pools": {"$ref": "#/components/schemas/load-balancing_pop_pools"}, "proxied": {"$ref": "#/components/schemas/load-balancing_proxied"}, "random_steering": {"$ref": "#/components/schemas/load-balancing_random_steering"}, "region_pools": {"$ref": "#/components/schemas/load-balancing_region_pools"}, "rules": {"$ref": "#/components/schemas/load-balancing_rules"}, "session_affinity": {"$ref": "#/components/schemas/load-balancing_session_affinity"}, "session_affinity_attributes": {"$ref": "#/components/schemas/load-balancing_session_affinity_attributes"}, "session_affinity_ttl": {"$ref": "#/components/schemas/load-balancing_session_affinity_ttl"}, "steering_policy": {"$ref": "#/components/schemas/load-balancing_steering_policy"}, "ttl": {"$ref": "#/components/schemas/load-balancing_ttl"}}}}}}, "responses": {"200": {"description": "Patch Load Balancer response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-single_response"}}}}, "4XX": {"description": "Patch Load Balancer response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-single_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancers"], "x-api-token-group": ["Load Balancers Write"]}
```
