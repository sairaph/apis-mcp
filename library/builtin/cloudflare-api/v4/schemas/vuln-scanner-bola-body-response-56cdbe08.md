---
title: vuln_scanner_bola-body-response
page_id: schema-vuln-scanner-bola-body-response-56cdbe08
path: schemas
description: HTTP response body preview.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-body-response

HTTP response body preview.

```yaml
{"description": "HTTP response body preview.", "discriminator": {"mapping": {"bytes": "#/components/schemas/vuln_scanner_bola-body-response-bytes", "json": "#/components/schemas/vuln_scanner_bola-body-response-json", "not_found": "#/components/schemas/vuln_scanner_bola-body-response-not-found", "text": "#/components/schemas/vuln_scanner_bola-body-response-text"}, "propertyName": "kind"}, "oneOf": [{"$ref": "#/components/schemas/vuln_scanner_bola-body-response-not-found"}, {"$ref": "#/components/schemas/vuln_scanner_bola-body-response-bytes"}, {"$ref": "#/components/schemas/vuln_scanner_bola-body-response-text"}, {"$ref": "#/components/schemas/vuln_scanner_bola-body-response-json"}]}
```
