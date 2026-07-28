---
title: zones_respect_strong_etag
page_id: schema-zones-respect-strong-etag-08a54c02
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_respect_strong_etag

```yaml
{"type": "object", "properties": {"id": {"description": "Turn on or off byte-for-byte equivalency checks between the\nCloudflare cache and the origin server.\n", "type": "string", "enum": ["respect_strong_etag"], "x-auditable": true}, "value": {"description": "The status of Respect Strong ETags\n", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}, "title": "Respect Strong ETags", "x-stainless-skip": ["terraform"]}
```
