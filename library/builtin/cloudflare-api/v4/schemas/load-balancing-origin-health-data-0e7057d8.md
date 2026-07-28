---
title: load-balancing_origin_health_data
page_id: schema-load-balancing-origin-health-data-0e7057d8
path: schemas
description: The origin ipv4/ipv6 address or domain name mapped to its health data.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_origin_health_data

The origin ipv4/ipv6 address or domain name mapped to its health data.

```yaml
{"description": "The origin ipv4/ipv6 address or domain name mapped to its health data.", "type": "object", "properties": {"failure_reason": {"type": "string", "x-auditable": true}, "healthy": {"type": "boolean", "x-auditable": true}, "response_code": {"type": "number", "x-auditable": true}, "rtt": {"type": "string", "x-auditable": true}}, "example": {"failure_reason": "No failures", "healthy": true, "response_code": 200, "rtt": "66ms"}}
```
