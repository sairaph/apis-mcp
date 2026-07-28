---
title: tunnel_subnet_capacity
page_id: schema-tunnel-subnet-capacity-978e2b2a
path: schemas
description: IP capacity information for the subnet.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_subnet_capacity

IP capacity information for the subnet.

```yaml
{"description": "IP capacity information for the subnet.", "type": "object", "properties": {"total": {"description": "Total number of assignable IPs in the subnet.", "type": "integer", "example": 254}, "used": {"description": "Number of assigned IPs in the subnet.", "type": "integer", "example": 42}}}
```
