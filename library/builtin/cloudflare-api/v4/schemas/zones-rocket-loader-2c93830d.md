---
title: zones_rocket_loader
page_id: schema-zones-rocket-loader-2c93830d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_rocket_loader

```yaml
{"type": "object", "properties": {"id": {"description": "Turn on or off Rocket Loader in the Cloudflare Speed app.\n", "type": "string", "enum": ["rocket_loader"], "x-auditable": true}, "value": {"description": "The status of Rocket Loader\n", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}, "title": "Rocket Loader", "x-stainless-skip": ["terraform"]}
```
