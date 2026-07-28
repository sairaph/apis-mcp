---
title: vuln_scanner_bola-assertion-kind
page_id: schema-vuln-scanner-bola-assertion-kind-9718d778
path: schemas
description: The kind of assertion to make.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-assertion-kind

The kind of assertion to make.

```yaml
{"description": "The kind of assertion to make.", "oneOf": [{"description": "Assert that an HTTP status code is within a range, e.g. to assert success for the 2xx range, or expected failure for the 4xx/5xx range.", "properties": {"parameters": {"$ref": "#/components/schemas/vuln_scanner_bola-http-status-range"}, "type": {"type": "string", "enum": ["http_status_within_range"]}}, "required": ["type", "parameters"], "type": "object"}]}
```
