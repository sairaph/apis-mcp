---
title: vuln_scanner_bola-test-error
page_id: schema-vuln-scanner-bola-test-error-d01f218e
path: schemas
description: Error that occurred during a test.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-test-error

Error that occurred during a test.

```yaml
{"description": "Error that occurred during a test.", "type": "object", "properties": {"description": {"description": "Human-readable error description.", "type": "string"}, "error_code": {"description": "Numeric error code identifying the class of error, if available.", "type": "integer", "format": "uint32", "minimum": 0, "nullable": true}}, "required": ["description"]}
```
