---
title: zones_polish
page_id: schema-zones-polish-b63764b4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_polish

```yaml
{"type": "object", "properties": {"id": {"description": "Apply options from the Polish feature of the Cloudflare Speed app.\n", "type": "string", "example": "polish", "enum": ["polish"], "x-auditable": true}, "value": {"description": "The level of Polish you want applied to your origin.\n", "type": "string", "example": "lossless", "enum": ["off", "lossless", "lossy"], "x-auditable": true}}, "title": "Polish", "x-stainless-skip": ["terraform"]}
```
