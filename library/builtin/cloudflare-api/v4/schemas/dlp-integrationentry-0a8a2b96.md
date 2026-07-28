---
title: dlp_IntegrationEntry
page_id: schema-dlp-integrationentry-0a8a2b96
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_IntegrationEntry

```yaml
{"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "enabled": {"type": "boolean"}, "id": {"type": "string", "format": "uuid"}, "name": {"type": "string"}, "profile_id": {"type": "string", "format": "uuid", "nullable": true, "x-stainless-terraform-configurability": "computed_optional"}, "updated_at": {"type": "string", "format": "date-time"}}, "required": ["id", "name", "created_at", "updated_at", "enabled"]}
```
