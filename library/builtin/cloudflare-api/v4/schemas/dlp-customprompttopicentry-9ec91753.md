---
title: dlp_CustomPromptTopicEntry
page_id: schema-dlp-customprompttopicentry-9ec91753
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_CustomPromptTopicEntry

```yaml
{"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "description": {"description": "The optional description of the custom prompt topic entry.", "type": "string", "nullable": true}, "enabled": {"type": "boolean", "deprecated": true}, "id": {"type": "string", "format": "uuid"}, "name": {"type": "string"}, "updated_at": {"type": "string", "format": "date-time"}}, "required": ["id", "name", "created_at", "updated_at", "enabled"]}
```
