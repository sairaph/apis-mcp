---
title: dlp_SkipConfig
page_id: schema-dlp-skipconfig-5d4db4bf
path: schemas
description: Content types to exclude from context analysis and return all matches.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_SkipConfig

Content types to exclude from context analysis and return all matches.

```yaml
{"description": "Content types to exclude from context analysis and return all matches.", "type": "object", "properties": {"files": {"description": "If the content type is a file, skip context analysis and return all matches.", "type": "boolean"}}, "required": ["files"]}
```
