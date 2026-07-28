---
title: cloudforce-one_AccountExemptions
page_id: schema-cloudforce-one-accountexemptions-e25a0150
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one_AccountExemptions

```yaml
{"type": "object", "properties": {"namespace": {"type": "array", "items": {"type": "string"}, "example": ["^test-.*$"]}, "tag_match": {"type": "array", "items": {"type": "string"}, "example": ["^staging-.*$"]}, "worker_name": {"type": "array", "items": {"type": "string"}, "example": ["^demo-.*$"]}}, "required": ["namespace", "tag_match", "worker_name"]}
```
