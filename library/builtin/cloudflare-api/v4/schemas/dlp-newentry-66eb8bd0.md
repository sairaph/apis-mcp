---
title: dlp_NewEntry
page_id: schema-dlp-newentry-66eb8bd0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_NewEntry

```yaml
{"type": "object", "properties": {"description": {"type": "string", "nullable": true}, "enabled": {"type": "boolean"}, "name": {"type": "string"}, "pattern": {"$ref": "#/components/schemas/dlp_Pattern"}, "profile_id": {"type": "string", "format": "uuid"}}, "required": ["name", "pattern", "enabled"]}
```
