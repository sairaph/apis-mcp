---
title: zones_mirage
page_id: schema-zones-mirage-d421b0cd
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_mirage

```yaml
{"type": "object", "properties": {"id": {"description": "Cloudflare Mirage reduces bandwidth used by images in mobile browsers.\nIt can accelerate loading of image-heavy websites on very slow mobile connections and HTTP/1.\n", "type": "string", "example": "mirage", "enum": ["mirage"], "x-auditable": true}, "value": {"description": "The status of Mirage.\n", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}, "deprecated": true, "title": "Mirage", "x-stainless-deprecation-message": "Mirage is deprecated. This functionality is no longer supported. \nSee https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#2025-11-03 for further details.\n", "x-stainless-skip": ["terraform"]}
```
