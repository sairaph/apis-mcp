---
title: vuln_scanner_bola-body-response-text
page_id: schema-vuln-scanner-bola-body-response-text-f5af08e5
path: schemas
description: Body received as valid UTF-8 text but not valid JSON.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-body-response-text

Body received as valid UTF-8 text but not valid JSON.

```yaml
{"description": "Body received as valid UTF-8 text but not valid JSON.", "type": "object", "properties": {"contents": {"type": "string"}, "kind": {"type": "string", "enum": ["text"]}, "truncated": {"type": "boolean"}}, "required": ["kind", "contents", "truncated"]}
```
