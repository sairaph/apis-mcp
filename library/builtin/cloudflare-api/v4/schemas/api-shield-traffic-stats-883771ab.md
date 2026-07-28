---
title: api-shield_traffic_stats
page_id: schema-api-shield-traffic-stats-883771ab
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_traffic_stats

```yaml
{"type": "object", "properties": {"traffic_stats": {"type": "object", "properties": {"last_updated": {"$ref": "#/components/schemas/api-shield_timestamp-2"}, "period_seconds": {"description": "The period in seconds these statistics were computed over", "type": "integer", "example": 3600, "readOnly": true, "x-auditable": true}, "requests": {"description": "The average number of requests seen during this period", "type": "number", "format": "float", "example": 1987.06, "readOnly": true, "x-auditable": true}}, "required": ["period_seconds", "requests", "last_updated"]}}, "readOnly": true}
```
