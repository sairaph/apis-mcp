---
title: load-balancing_weight
page_id: schema-load-balancing-weight-73684eeb
path: schemas
description: |-
    The weight of this origin relative to other origins in the pool. Based on the configured weight the total traffic is distributed among origins within the pool.
    - `origin_steering.policy="least_outstanding_requests"`: Use weight to scale the origin's outstanding requests.
    - `origin_steering.policy="least_connections"`: Use weight to scale the origin's open connections.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_weight

The weight of this origin relative to other origins in the pool. Based on the configured weight the total traffic is distributed among origins within the pool.
- `origin_steering.policy="least_outstanding_requests"`: Use weight to scale the origin's outstanding requests.
- `origin_steering.policy="least_connections"`: Use weight to scale the origin's open connections.

```yaml
{"description": "The weight of this origin relative to other origins in the pool. Based on the configured weight the total traffic is distributed among origins within the pool.\n- `origin_steering.policy=\"least_outstanding_requests\"`: Use weight to scale the origin's outstanding requests.\n- `origin_steering.policy=\"least_connections\"`: Use weight to scale the origin's open connections.", "type": "number", "example": 0.6, "default": 1, "maximum": 1, "minimum": 0, "multipleOf": 0.01, "x-auditable": true}
```
