---
title: zones_0rtt
page_id: schema-zones-0rtt-15dffcc2
path: schemas
description: 0-RTT session resumption enabled for this zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_0rtt

0-RTT session resumption enabled for this zone.

```yaml
{"description": "0-RTT session resumption enabled for this zone.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "0rtt", "enum": ["0rtt"]}, "value": {"$ref": "#/components/schemas/zones_0rtt_value"}}}], "title": "0-RTT Value"}
```
