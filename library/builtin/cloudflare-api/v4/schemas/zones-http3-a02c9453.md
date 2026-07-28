---
title: zones_http3
page_id: schema-zones-http3-a02c9453
path: schemas
description: HTTP3 enabled for this zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_http3

HTTP3 enabled for this zone.

```yaml
{"description": "HTTP3 enabled for this zone.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "http3", "enum": ["http3"]}, "value": {"$ref": "#/components/schemas/zones_http3_value"}}}], "title": "HTTP3 Value"}
```
