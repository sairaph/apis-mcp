---
title: rum_ruleset
page_id: schema-rum-ruleset-e21b9e12
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rum_ruleset

```yaml
{"type": "object", "properties": {"enabled": {"description": "Whether the ruleset is enabled.", "type": "boolean", "example": true}, "id": {"$ref": "#/components/schemas/rum_ruleset_identifier"}, "zone_name": {"type": "string", "example": "example.com"}, "zone_tag": {"$ref": "#/components/schemas/rum_zone_tag"}}}
```
