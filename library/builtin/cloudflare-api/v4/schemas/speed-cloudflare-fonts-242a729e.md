---
title: speed_cloudflare_fonts
page_id: schema-speed-cloudflare-fonts-242a729e
path: schemas
description: |-
    Enhance your website's font delivery with Cloudflare Fonts. Deliver Google Hosted fonts from your own domain,
    boost performance, and enhance user privacy. Refer to the Cloudflare Fonts documentation for more information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# speed_cloudflare_fonts

Enhance your website's font delivery with Cloudflare Fonts. Deliver Google Hosted fonts from your own domain,
boost performance, and enhance user privacy. Refer to the Cloudflare Fonts documentation for more information.

```yaml
{"description": "Enhance your website's font delivery with Cloudflare Fonts. Deliver Google Hosted fonts from your own domain,\nboost performance, and enhance user privacy. Refer to the Cloudflare Fonts documentation for more information.\n", "allOf": [{"$ref": "#/components/schemas/speed_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "type": "string", "example": "fonts", "enum": ["fonts"], "x-auditable": true}, "value": {"$ref": "#/components/schemas/speed_cloudflare_fonts_value"}}, "type": "object"}], "title": "Cloudflare Fonts"}
```
