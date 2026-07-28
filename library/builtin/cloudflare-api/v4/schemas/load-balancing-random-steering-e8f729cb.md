---
title: load-balancing_random_steering
page_id: schema-load-balancing-random-steering-e8f729cb
path: schemas
description: |-
    Configures pool weights.
    - `steering_policy="random"`: A random pool is selected with probability proportional to pool weights.
    - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each pool's outstanding requests.
    - `steering_policy="least_connections"`: Use pool weights to scale each pool's open connections.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_random_steering

Configures pool weights.
- `steering_policy="random"`: A random pool is selected with probability proportional to pool weights.
- `steering_policy="least_outstanding_requests"`: Use pool weights to scale each pool's outstanding requests.
- `steering_policy="least_connections"`: Use pool weights to scale each pool's open connections.

```yaml
{"description": "Configures pool weights.\n- `steering_policy=\"random\"`: A random pool is selected with probability proportional to pool weights.\n- `steering_policy=\"least_outstanding_requests\"`: Use pool weights to scale each pool's outstanding requests.\n- `steering_policy=\"least_connections\"`: Use pool weights to scale each pool's open connections.", "type": "object", "properties": {"default_weight": {"description": "The default weight for pools in the load balancer that are not specified in the pool_weights map.", "type": "number", "example": 0.2, "default": 1, "maximum": 1, "minimum": 0, "multipleOf": 0.1, "x-auditable": true}, "pool_weights": {"description": "A mapping of pool IDs to custom weights. The weight is relative to other pools in the load balancer.", "type": "object", "example": {"9290f38c5d07c2e2f4df57b1f61d4196": 0.5, "de90f38ced07c2e2f4df50b1f61d4194": 0.3}, "additionalProperties": {"type": "number", "x-auditable": true}}}}
```
