---
title: zones_base
page_id: schema-zones-base-19fe863d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_base

```yaml
{"properties": {"editable": {"description": "Whether or not this setting can be modified for this zone (based on your Cloudflare plan level).", "type": "boolean", "default": true, "enum": [true, false], "readOnly": true}, "id": {"description": "Identifier of the zone setting.", "type": "string", "example": "development_mode"}, "modified_on": {"description": "last time this setting was modified.", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00.12345Z", "nullable": true, "readOnly": true}, "value": {"description": "Current value of the zone setting.", "example": "on"}}, "required": ["id", "value"]}
```
