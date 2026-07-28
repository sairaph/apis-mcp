---
title: vuln_scanner_delete-scan-response
page_id: schema-vuln-scanner-delete-scan-response-24e01666
path: schemas
description: Successful scan deletion result.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_delete-scan-response

Successful scan deletion result.

```yaml
{"description": "Successful scan deletion result.", "type": "object", "properties": {"id": {"description": "ID of the deleted scan.", "type": "string", "format": "uuid", "x-auditable": true}}, "required": ["id"]}
```
