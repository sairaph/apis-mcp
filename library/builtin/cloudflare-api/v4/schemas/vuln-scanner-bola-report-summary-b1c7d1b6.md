---
title: vuln_scanner_bola-report-summary
page_id: schema-vuln-scanner-bola-report-summary-b1c7d1b6
path: schemas
description: Overall report summary.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-report-summary

Overall report summary.

```yaml
{"description": "Overall report summary.", "type": "object", "properties": {"verdict": {"description": "Overall verdict of the vulnerability scan.", "allOf": [{"$ref": "#/components/schemas/vuln_scanner_bola-verdict"}]}}, "required": ["verdict"]}
```
