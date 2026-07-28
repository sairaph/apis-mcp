---
title: zones_browser_check
page_id: schema-zones-browser-check-44f531b9
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_browser_check

```yaml
{"type": "object", "properties": {"id": {"description": "Inspect the visitor's browser for headers commonly associated with\nspammers and certain bots.\n", "type": "string", "enum": ["browser_check"], "x-auditable": true}, "value": {"description": "The status of Browser Integrity Check.\n", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}, "title": "Browser Integrity Check", "x-stainless-skip": ["terraform"]}
```
