---
title: resource-tagging_tag_key_summary
page_id: schema-resource-tagging-tag-key-summary-68e6682c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_tag_key_summary

```yaml
{"type": "object", "properties": {"key": {"description": "A tag key.", "type": "string", "example": "environment"}, "values": {"description": "All distinct values for this tag key.", "type": "array", "items": {"type": "string"}, "example": ["production", "staging"]}}, "required": ["key", "values"]}
```
