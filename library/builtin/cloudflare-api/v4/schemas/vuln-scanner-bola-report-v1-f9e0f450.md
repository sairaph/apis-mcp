---
title: vuln_scanner_bola-report-v1
page_id: schema-vuln-scanner-bola-report-v1-f9e0f450
path: schemas
description: Version 1 of the BOLA vulnerability scan report.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-report-v1

Version 1 of the BOLA vulnerability scan report.

```yaml
{"description": "Version 1 of the BOLA vulnerability scan report.", "type": "object", "properties": {"summary": {"description": "Summary of all steps and findings.", "allOf": [{"$ref": "#/components/schemas/vuln_scanner_bola-report-summary"}]}, "tests": {"description": "List of tests that were run.", "type": "array", "items": {"$ref": "#/components/schemas/vuln_scanner_bola-test"}}}, "required": ["summary", "tests"]}
```
