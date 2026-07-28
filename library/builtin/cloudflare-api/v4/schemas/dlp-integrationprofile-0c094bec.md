---
title: dlp_IntegrationProfile
page_id: schema-dlp-integrationprofile-0c094bec
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_IntegrationProfile

```yaml
{"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "description": {"description": "The description of the profile.", "type": "string", "nullable": true}, "entries": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_Entry"}, "deprecated": true}, "id": {"type": "string", "format": "uuid"}, "name": {"type": "string"}, "shared_entries": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_Entry"}, "x-stainless-terraform-configurability": "computed"}, "updated_at": {"type": "string", "format": "date-time"}}, "required": ["id", "name", "entries", "shared_entries", "created_at", "updated_at"]}
```
