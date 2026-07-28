---
title: load-balancing_location_strategy
page_id: schema-load-balancing-location-strategy-f2a4ecbf
path: schemas
description: Controls location-based steering for non-proxied requests. See `steering_policy` to learn how steering is affected.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_location_strategy

Controls location-based steering for non-proxied requests. See `steering_policy` to learn how steering is affected.

```yaml
{"description": "Controls location-based steering for non-proxied requests. See `steering_policy` to learn how steering is affected.", "type": "object", "properties": {"mode": {"description": "Determines the authoritative location when ECS is not preferred, does not exist in the request, or its GeoIP lookup is unsuccessful.\n- `\"pop\"`: Use the Cloudflare PoP location.\n- `\"resolver_ip\"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is unsuccessful, use the Cloudflare PoP location.", "type": "string", "example": "resolver_ip", "default": "pop", "enum": ["pop", "resolver_ip"], "x-auditable": true}, "prefer_ecs": {"description": "Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the authoritative location.\n- `\"always\"`: Always prefer ECS.\n- `\"never\"`: Never prefer ECS.\n- `\"proximity\"`: Prefer ECS only when `steering_policy=\"proximity\"`.\n- `\"geo\"`: Prefer ECS only when `steering_policy=\"geo\"`.", "type": "string", "example": "always", "default": "proximity", "enum": ["always", "never", "proximity", "geo"], "x-auditable": true}}}
```
