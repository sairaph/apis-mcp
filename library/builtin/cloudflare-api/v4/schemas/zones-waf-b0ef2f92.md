---
title: zones_waf
page_id: schema-zones-waf-b0ef2f92
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_waf

```yaml
{"type": "object", "properties": {"id": {"description": "Turn on or off [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).\nYou cannot enable or disable individual WAF managed rules via Page Rules.\n", "type": "string", "example": "waf", "enum": ["waf"], "x-auditable": true}, "value": {"description": "The status of WAF managed rules (previous version).\n", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}, "title": "WAF", "x-stainless-skip": ["terraform"]}
```
