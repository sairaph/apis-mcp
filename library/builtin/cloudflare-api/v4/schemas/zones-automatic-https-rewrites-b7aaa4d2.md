---
title: zones_automatic_https_rewrites
page_id: schema-zones-automatic-https-rewrites-b7aaa4d2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_automatic_https_rewrites

```yaml
{"type": "object", "properties": {"id": {"description": "Turn on or off Automatic HTTPS Rewrites.", "type": "string", "enum": ["automatic_https_rewrites"], "x-auditable": true}, "value": {"description": "The status of Automatic HTTPS Rewrites.\n", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}, "title": "Automatic HTTPS Rewrites", "x-stainless-skip": ["terraform"]}
```
