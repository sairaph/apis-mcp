---
title: vuln_scanner_bola-body-response-bytes
page_id: schema-vuln-scanner-bola-body-response-bytes-8575f73d
path: schemas
description: Body received but unable to read as UTF-8. Raw bytes, base64-encoded.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-body-response-bytes

Body received but unable to read as UTF-8. Raw bytes, base64-encoded.

```yaml
{"description": "Body received but unable to read as UTF-8. Raw bytes, base64-encoded.", "type": "object", "properties": {"contents": {"type": "string"}, "kind": {"type": "string", "enum": ["bytes"]}, "truncated": {"type": "boolean"}}, "required": ["kind", "contents", "truncated"]}
```
