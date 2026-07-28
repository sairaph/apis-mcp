---
title: vuln_scanner_bola-test-step
page_id: schema-vuln-scanner-bola-test-step-344d0f8a
path: schemas
description: A single step in a test.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-test-step

A single step in a test.

```yaml
{"description": "A single step in a test.", "type": "object", "properties": {"assertions": {"description": "Assertions that were made against the received response.", "type": "array", "items": {"$ref": "#/components/schemas/vuln_scanner_bola-test-assertion"}}, "errors": {"description": "Errors the step encountered that may explain absent or incomplete fields.", "type": "array", "items": {"$ref": "#/components/schemas/vuln_scanner_bola-test-error"}}, "request": {"description": "HTTP request that was made, if any.", "type": "object", "allOf": [{"$ref": "#/components/schemas/vuln_scanner_bola-test-request"}], "nullable": true}, "response": {"description": "HTTP response that was received, if any.", "type": "object", "allOf": [{"$ref": "#/components/schemas/vuln_scanner_bola-test-response"}], "nullable": true}}, "required": ["assertions"]}
```
