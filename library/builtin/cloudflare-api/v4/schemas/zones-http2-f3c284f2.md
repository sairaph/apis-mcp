---
title: zones_http2
page_id: schema-zones-http2-f3c284f2
path: schemas
description: HTTP2 enabled for this zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_http2

HTTP2 enabled for this zone.

```yaml
{"description": "HTTP2 enabled for this zone.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "http2", "enum": ["http2"]}, "value": {"$ref": "#/components/schemas/zones_http2_value"}}}], "title": "HTTP2 Value"}
```
