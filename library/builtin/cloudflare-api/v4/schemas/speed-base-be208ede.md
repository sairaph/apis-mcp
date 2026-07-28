---
title: speed_base
page_id: schema-speed-base-be208ede
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# speed_base

```yaml
{"properties": {"editable": {"description": "Whether or not this setting can be modified for this zone (based on your Cloudflare plan level).", "type": "boolean", "default": true, "enum": [true, false], "readOnly": true, "x-auditable": true}, "id": {"description": "Identifier of the zone setting.", "type": "string", "example": "development_mode", "x-auditable": true}, "modified_on": {"description": "last time this setting was modified.", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00.12345Z", "nullable": true, "readOnly": true, "x-auditable": true}, "value": {"description": "Current value of the zone setting.", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}}
```
