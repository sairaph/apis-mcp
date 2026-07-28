---
title: vuln_scanner_bola-verdict
page_id: schema-vuln-scanner-bola-verdict-446c8873
path: schemas
description: A verdict. `ok` means the scan passed, `warning` means the scan detected issues, `inconclusive` means errors prevented the scanner from reaching an accurate verdict.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-verdict

A verdict. `ok` means the scan passed, `warning` means the scan detected issues, `inconclusive` means errors prevented the scanner from reaching an accurate verdict.

```yaml
{"description": "A verdict. `ok` means the scan passed, `warning` means the scan detected issues, `inconclusive` means errors prevented the scanner from reaching an accurate verdict.", "type": "string", "enum": ["ok", "warning", "inconclusive"]}
```
