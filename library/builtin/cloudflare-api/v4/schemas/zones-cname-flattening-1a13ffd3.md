---
title: zones_cname_flattening
page_id: schema-zones-cname-flattening-1a13ffd3
path: schemas
description: Whether or not cname flattening is on.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_cname_flattening

Whether or not cname flattening is on.

```yaml
{"description": "Whether or not cname flattening is on.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "How to flatten the cname destination.", "enum": ["cname_flattening"]}, "value": {"$ref": "#/components/schemas/zones_cname_flattening_value"}}}], "deprecated": true, "title": "Cloudflare CNAME Flattening", "x-stainless-deprecation-message": "This zone setting is deprecated; please use the DNS Settings route instead. More information at https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#2025-03-21"}
```
