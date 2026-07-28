---
title: dlp_ContextAwareness
page_id: schema-dlp-contextawareness-44da2f10
path: schemas
description: Scan the context of predefined entries to only return matches surrounded by keywords.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_ContextAwareness

Scan the context of predefined entries to only return matches surrounded by keywords.

```yaml
{"description": "Scan the context of predefined entries to only return matches surrounded by keywords.", "type": "object", "properties": {"enabled": {"description": "If true, scan the context of predefined entries to only return matches surrounded by keywords.", "type": "boolean"}, "skip": {"$ref": "#/components/schemas/dlp_SkipConfig"}}, "deprecated": true, "required": ["enabled", "skip"]}
```
