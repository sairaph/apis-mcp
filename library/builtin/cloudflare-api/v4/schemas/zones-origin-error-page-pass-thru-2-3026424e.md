---
title: zones_origin_error_page_pass_thru-2
page_id: schema-zones-origin-error-page-pass-thru-2-3026424e
path: schemas
description: Cloudflare will proxy customer error pages on any 502,504 errors on origin server instead of showing a default Cloudflare error page. This does not apply to 522 errors and is limited to Enterprise Zones.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_origin_error_page_pass_thru-2

Cloudflare will proxy customer error pages on any 502,504 errors on origin server instead of showing a default Cloudflare error page. This does not apply to 522 errors and is limited to Enterprise Zones.

```yaml
{"description": "Cloudflare will proxy customer error pages on any 502,504 errors on origin server instead of showing a default Cloudflare error page. This does not apply to 522 errors and is limited to Enterprise Zones.", "default": "off", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "origin_error_page_pass_thru", "enum": ["origin_error_page_pass_thru"]}, "value": {"$ref": "#/components/schemas/zones_origin_error_page_pass_thru_value"}}}], "title": "Error Pages On"}
```
