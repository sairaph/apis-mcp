---
title: vuln_scanner_bola-test
page_id: schema-vuln-scanner-bola-test-1026610b
path: schemas
description: Result of a single test.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-test

Result of a single test.

```yaml
{"description": "Result of a single test.", "type": "object", "properties": {"preflight_errors": {"description": "Errors that prevented step execution.", "type": "array", "items": {"$ref": "#/components/schemas/vuln_scanner_bola-test-error"}}, "steps": {"description": "Steps that were executed.", "type": "array", "items": {"$ref": "#/components/schemas/vuln_scanner_bola-test-step"}}, "verdict": {"description": "Verdict of this single test.", "allOf": [{"$ref": "#/components/schemas/vuln_scanner_bola-verdict"}]}}, "required": ["verdict", "steps"]}
```
