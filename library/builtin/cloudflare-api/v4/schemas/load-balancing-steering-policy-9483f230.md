---
title: load-balancing_steering_policy
page_id: schema-load-balancing-steering-policy-9483f230
path: schemas
description: |-
    Steering Policy for this load balancer.
    - `"off"`: Use `default_pools`.
    - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied requests, the country for `country_pools` is determined by `location_strategy`.
    - `"random"`: Select a pool randomly.
    - `"dynamic_latency"`: Use round trip time to select the closest pool in default_pools (requires pool health checks).
    - `"proximity"`: Use the pools' latitude and longitude to select the closest pool using the Cloudflare PoP location for proxied requests or the location determined by `location_strategy` for non-proxied requests.
    - `"least_outstanding_requests"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of outstanding requests. Pools with more pending requests are weighted proportionately less relative to others.
    - `"least_connections"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of open connections. Pools with more open connections are weighted proportionately less relative to others. Supported for HTTP/1 and HTTP/2 connections.
    - `""`: Will map to `"geo"` if you use `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_steering_policy

Steering Policy for this load balancer.
- `"off"`: Use `default_pools`.
- `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied requests, the country for `country_pools` is determined by `location_strategy`.
- `"random"`: Select a pool randomly.
- `"dynamic_latency"`: Use round trip time to select the closest pool in default_pools (requires pool health checks).
- `"proximity"`: Use the pools' latitude and longitude to select the closest pool using the Cloudflare PoP location for proxied requests or the location determined by `location_strategy` for non-proxied requests.
- `"least_outstanding_requests"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of outstanding requests. Pools with more pending requests are weighted proportionately less relative to others.
- `"least_connections"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of open connections. Pools with more open connections are weighted proportionately less relative to others. Supported for HTTP/1 and HTTP/2 connections.
- `""`: Will map to `"geo"` if you use `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.

```yaml
{"description": "Steering Policy for this load balancer.\n- `\"off\"`: Use `default_pools`.\n- `\"geo\"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied requests, the country for `country_pools` is determined by `location_strategy`.\n- `\"random\"`: Select a pool randomly.\n- `\"dynamic_latency\"`: Use round trip time to select the closest pool in default_pools (requires pool health checks).\n- `\"proximity\"`: Use the pools' latitude and longitude to select the closest pool using the Cloudflare PoP location for proxied requests or the location determined by `location_strategy` for non-proxied requests.\n- `\"least_outstanding_requests\"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of outstanding requests. Pools with more pending requests are weighted proportionately less relative to others.\n- `\"least_connections\"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of open connections. Pools with more open connections are weighted proportionately less relative to others. Supported for HTTP/1 and HTTP/2 connections.\n- `\"\"`: Will map to `\"geo\"` if you use `region_pools`/`country_pools`/`pop_pools` otherwise `\"off\"`.", "type": "string", "example": "dynamic_latency", "default": "", "enum": ["off", "geo", "random", "dynamic_latency", "proximity", "least_outstanding_requests", "least_connections", ""], "x-auditable": true}
```
