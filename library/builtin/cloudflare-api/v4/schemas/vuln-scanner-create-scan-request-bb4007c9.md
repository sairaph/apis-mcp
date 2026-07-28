---
title: vuln_scanner_create-scan-request
page_id: schema-vuln-scanner-create-scan-request-bb4007c9
path: schemas
description: |-
    Create a new vulnerability scan. The `scan_type` discriminator
    selects the scan variant and its required context fields.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_create-scan-request

Create a new vulnerability scan. The `scan_type` discriminator
selects the scan variant and its required context fields.

```yaml
{"description": "Create a new vulnerability scan. The `scan_type` discriminator\nselects the scan variant and its required context fields.\n", "discriminator": {"mapping": {"bola": "#/components/schemas/vuln_scanner_create-bola-scan-request"}, "propertyName": "scan_type"}, "oneOf": [{"$ref": "#/components/schemas/vuln_scanner_create-bola-scan-request"}]}
```
