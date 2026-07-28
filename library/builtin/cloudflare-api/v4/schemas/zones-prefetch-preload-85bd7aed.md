---
title: zones_prefetch_preload
page_id: schema-zones-prefetch-preload-85bd7aed
path: schemas
description: Cloudflare will prefetch any URLs that are included in the response headers. This is limited to Enterprise Zones.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_prefetch_preload

Cloudflare will prefetch any URLs that are included in the response headers. This is limited to Enterprise Zones.

```yaml
{"description": "Cloudflare will prefetch any URLs that are included in the response headers. This is limited to Enterprise Zones.", "default": "off", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "prefetch_preload", "enum": ["prefetch_preload"]}, "value": {"$ref": "#/components/schemas/zones_prefetch_preload_value"}}}], "title": "Prefetch preload"}
```
