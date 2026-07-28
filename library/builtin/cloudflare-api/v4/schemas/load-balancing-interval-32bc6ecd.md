---
title: load-balancing_interval
page_id: schema-load-balancing-interval-32bc6ecd
path: schemas
description: The interval between each health check. Shorter intervals may improve failover time, but will increase load on the origins as we check from multiple locations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_interval

The interval between each health check. Shorter intervals may improve failover time, but will increase load on the origins as we check from multiple locations.

```yaml
{"description": "The interval between each health check. Shorter intervals may improve failover time, but will increase load on the origins as we check from multiple locations.", "type": "integer", "default": 60, "x-auditable": true}
```
