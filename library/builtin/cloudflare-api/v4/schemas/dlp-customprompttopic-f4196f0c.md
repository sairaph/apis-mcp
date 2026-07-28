---
title: dlp_CustomPromptTopic
page_id: schema-dlp-customprompttopic-f4196f0c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_CustomPromptTopic

```yaml
{"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "description": {"type": "string", "nullable": true}, "enabled": {"type": "boolean", "deprecated": true}, "id": {"type": "string", "format": "uuid"}, "name": {"type": "string"}, "profile_id": {"type": "string", "format": "uuid", "deprecated": true, "nullable": true}, "topic": {"type": "string", "maxLength": 50, "minLength": 2}, "updated_at": {"type": "string", "format": "date-time"}}, "required": ["id", "name", "created_at", "updated_at", "enabled", "topic"]}
```
