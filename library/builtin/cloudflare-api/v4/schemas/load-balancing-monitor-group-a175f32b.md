---
title: load-balancing_monitor-group
page_id: schema-load-balancing-monitor-group-a175f32b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_monitor-group

```yaml
{"type": "object", "properties": {"created_on": {"description": "The timestamp of when the monitor group was created", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00.12345Z", "readOnly": true}, "description": {"description": "A short description of the monitor group", "type": "string", "example": "Primary datacenter monitors"}, "id": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor_group_id"}], "readOnly": true}, "members": {"description": "List of monitors in this group", "type": "array", "items": {"$ref": "#/components/schemas/load-balancing_monitor-group-member"}, "x-stainless-collection-type": "set"}, "modified_on": {"description": "The timestamp of when the monitor group was last updated", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00.12345Z", "readOnly": true}}, "required": ["id", "description", "members"]}
```
