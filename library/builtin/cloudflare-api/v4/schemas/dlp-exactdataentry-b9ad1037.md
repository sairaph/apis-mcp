---
title: dlp_ExactDataEntry
page_id: schema-dlp-exactdataentry-b9ad1037
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_ExactDataEntry

```yaml
{"type": "object", "properties": {"case_sensitive": {"description": "Only applies to custom word lists.\nDetermines if the words should be matched in a case-sensitive manner\nCannot be set to false if secret is true", "type": "boolean"}, "created_at": {"type": "string", "format": "date-time"}, "description": {"description": "The optional description of the exact data entry.", "type": "string", "nullable": true}, "enabled": {"type": "boolean"}, "id": {"type": "string", "format": "uuid"}, "name": {"type": "string"}, "secret": {"type": "boolean"}, "updated_at": {"type": "string", "format": "date-time"}}, "required": ["id", "name", "created_at", "updated_at", "enabled", "secret", "case_sensitive"]}
```
