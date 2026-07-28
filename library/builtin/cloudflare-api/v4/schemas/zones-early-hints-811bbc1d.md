---
title: zones_early_hints
page_id: schema-zones-early-hints-811bbc1d
path: schemas
description: When enabled, Cloudflare will attempt to speed up overall page loads by serving `103` responses with `Link` headers from the final response. Refer to [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for more information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_early_hints

When enabled, Cloudflare will attempt to speed up overall page loads by serving `103` responses with `Link` headers from the final response. Refer to [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for more information.

```yaml
{"description": "When enabled, Cloudflare will attempt to speed up overall page loads by serving `103` responses with `Link` headers from the final response. Refer to [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for more information.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "early_hints", "enum": ["early_hints"]}, "value": {"$ref": "#/components/schemas/zones_early_hints_value"}}}], "title": "Early Hints"}
```
