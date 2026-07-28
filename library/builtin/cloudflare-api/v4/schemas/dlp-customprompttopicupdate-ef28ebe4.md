---
title: dlp_CustomPromptTopicUpdate
page_id: schema-dlp-customprompttopicupdate-ef28ebe4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_CustomPromptTopicUpdate

```yaml
{"type": "object", "properties": {"description": {"type": "string", "nullable": true}, "enabled": {"type": "boolean"}, "name": {"type": "string"}, "topic": {"type": "string", "maxLength": 50, "minLength": 2}}, "required": ["name", "enabled", "topic"]}
```
