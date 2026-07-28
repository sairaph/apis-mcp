---
title: ai-audit_content_signal
page_id: schema-ai-audit-content-signal-427c8c38
path: schemas
description: Content signal directives from robots.txt.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# ai-audit_content_signal

Content signal directives from robots.txt.

```yaml
{"description": "Content signal directives from robots.txt.", "type": "object", "properties": {"ai-input": {"description": "Whether AI input usage is permitted.", "type": "string", "enum": ["yes", "no"]}, "ai-train": {"description": "Whether AI training is permitted.", "type": "string", "enum": ["yes", "no"]}, "search": {"description": "Whether search indexing is permitted.", "type": "string", "enum": ["yes", "no"]}}}
```
