---
title: vuln_scanner_bola-body-response-json
page_id: schema-vuln-scanner-bola-body-response-json-11cf36b7
path: schemas
description: Body received as valid JSON.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-body-response-json

Body received as valid JSON.

```yaml
{"description": "Body received as valid JSON.", "type": "object", "properties": {"contents": {"type": "string"}, "kind": {"type": "string", "enum": ["json"]}, "truncated": {"type": "boolean"}}, "required": ["kind", "contents", "truncated"]}
```
