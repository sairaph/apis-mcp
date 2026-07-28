---
title: dlp_EntryConfidence
page_id: schema-dlp-entryconfidence-4f9e3e6b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_EntryConfidence

```yaml
{"type": "object", "properties": {"ai_context_available": {"description": "Indicates whether this entry has AI remote service validation.", "type": "boolean"}, "available": {"description": "Indicates whether this entry has any form of validation that is not an AI remote service.", "type": "boolean"}}, "required": ["available", "ai_context_available"]}
```
