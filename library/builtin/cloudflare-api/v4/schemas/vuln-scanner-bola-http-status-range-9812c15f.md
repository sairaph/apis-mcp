---
title: vuln_scanner_bola-http-status-range
page_id: schema-vuln-scanner-bola-http-status-range-9812c15f
path: schemas
description: Range of HTTP status codes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-http-status-range

Range of HTTP status codes.

```yaml
{"description": "Range of HTTP status codes.", "type": "object", "properties": {"max": {"description": "Maximum (inclusive) status code of the range.", "type": "integer", "maximum": 65535, "minimum": 0}, "min": {"description": "Minimum (inclusive) status code of the range.", "type": "integer", "maximum": 65535, "minimum": 0}}, "required": ["min", "max"]}
```
