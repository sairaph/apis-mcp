---
title: zones_automatic_https_rewrites-2
page_id: schema-zones-automatic-https-rewrites-2-df819e0e
path: schemas
description: Enable the Automatic HTTPS Rewrites feature for this zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_automatic_https_rewrites-2

Enable the Automatic HTTPS Rewrites feature for this zone.

```yaml
{"description": "Enable the Automatic HTTPS Rewrites feature for this zone.", "default": "off", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "automatic_https_rewrites", "enum": ["automatic_https_rewrites"]}, "value": {"$ref": "#/components/schemas/zones_automatic_https_rewrites_value"}}}], "title": "Zone Enable Automatic HTTPS Rewrites"}
```
