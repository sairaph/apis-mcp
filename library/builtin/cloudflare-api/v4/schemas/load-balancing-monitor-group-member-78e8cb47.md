---
title: load-balancing_monitor-group-member
page_id: schema-load-balancing-monitor-group-member-78e8cb47
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_monitor-group-member

```yaml
{"type": "object", "properties": {"created_at": {"description": "The timestamp of when the monitor was added to the group", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00.12345Z", "readOnly": true}, "enabled": {"description": "Whether this monitor is enabled in the group", "type": "boolean", "example": true}, "monitor_id": {"$ref": "#/components/schemas/load-balancing_monitor_id"}, "monitoring_only": {"description": "Whether this monitor is used for monitoring only (does not affect pool health)", "type": "boolean", "example": false}, "must_be_healthy": {"description": "Whether this monitor must be healthy for the pool to be considered healthy", "type": "boolean", "example": true}, "updated_at": {"description": "The timestamp of when the monitor group member was last updated", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00.12345Z", "readOnly": true}}, "required": ["monitor_id", "enabled", "monitoring_only", "must_be_healthy"]}
```
