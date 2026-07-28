---
title: load-balancing_origin-health
page_id: schema-load-balancing-origin-health-bcd82f9c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_origin-health

```yaml
{"type": "object", "properties": {"ip": {"type": "object", "properties": {"failure_reason": {"description": "Failure reason.", "type": "string", "example": "No failure reasons", "x-auditable": true}, "healthy": {"description": "Origin health status.", "type": "boolean", "example": true, "x-auditable": true}, "response_code": {"description": "Response code from origin health check.", "type": "number", "example": 200, "x-auditable": true}, "rtt": {"description": "Origin RTT (Round Trip Time) response.", "type": "string", "example": "201.5ms", "x-auditable": true}}}}}
```
