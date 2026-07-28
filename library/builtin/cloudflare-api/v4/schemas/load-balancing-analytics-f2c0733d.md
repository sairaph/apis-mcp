---
title: load-balancing_analytics
page_id: schema-load-balancing-analytics-f2c0733d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_analytics

```yaml
{"type": "object", "properties": {"id": {"type": "integer", "default": 1}, "origins": {"type": "array", "items": {"$ref": "#/components/schemas/load-balancing_origin-analytics"}, "example": [{"address": "198.51.100.4", "changed": true, "enabled": true, "failure_reason": "No failures", "healthy": true, "ip": "198.51.100.4", "name": "some-origin"}]}, "pool": {"type": "object", "example": {"changed": true, "healthy": true, "id": "74bc6a8b9b0dda3d651707a2928bad0c", "minimum_origins": 1, "name": "some-pool"}}, "timestamp": {"type": "string", "format": "date-time", "example": "2014-01-01T05:20:00.12345Z"}}}
```
