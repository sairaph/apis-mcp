---
title: dlp_CustomEntry
page_id: schema-dlp-customentry-55fcf48e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_CustomEntry

```yaml
{"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "description": {"type": "string", "nullable": true}, "enabled": {"type": "boolean", "deprecated": true}, "id": {"type": "string", "format": "uuid"}, "name": {"type": "string"}, "pattern": {"$ref": "#/components/schemas/dlp_Pattern"}, "profile_id": {"type": "string", "format": "uuid", "deprecated": true, "nullable": true}, "updated_at": {"type": "string", "format": "date-time"}}, "required": ["id", "name", "created_at", "updated_at", "pattern", "enabled"]}
```
