---
title: zones_always_use_https
page_id: schema-zones-always-use-https-ad98bd7b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_always_use_https

```yaml
{"type": "object", "properties": {"id": {"description": "If enabled, any `http://`` URL is converted to `https://` through a\n301 redirect.\n", "type": "string", "enum": ["always_use_https"], "x-auditable": true}}, "title": "Always Use HTTPS", "x-stainless-skip": ["terraform"]}
```
