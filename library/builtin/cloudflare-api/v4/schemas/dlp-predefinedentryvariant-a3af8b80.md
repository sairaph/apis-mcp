---
title: dlp_PredefinedEntryVariant
page_id: schema-dlp-predefinedentryvariant-a3af8b80
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_PredefinedEntryVariant

```yaml
{"oneOf": [{"description": "A Predefined AI prompt classification topic entry.", "properties": {"description": {"description": "A customer-facing explanation of what this predefined AI prompt topic represents.", "type": "string", "nullable": true}, "topic_type": {"$ref": "#/components/schemas/dlp_PromptTopicType"}, "type": {"type": "string", "enum": ["PromptTopic"]}}, "required": ["topic_type", "type"], "type": "object"}, {"description": "A general predefined entry.", "properties": {"description": {"description": "A customer-facing explanation of what this predefined entry represents.", "type": "string", "nullable": true}, "type": {"type": "string", "enum": ["General"]}}, "required": ["type"], "type": "object"}]}
```
