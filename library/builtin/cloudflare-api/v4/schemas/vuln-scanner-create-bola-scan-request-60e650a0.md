---
title: vuln_scanner_create-bola-scan-request
page_id: schema-vuln-scanner-create-bola-scan-request-60e650a0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_create-bola-scan-request

```yaml
{"type": "object", "properties": {"credential_sets": {"$ref": "#/components/schemas/vuln_scanner_bola-credential-sets"}, "open_api": {"description": "OpenAPI schema definition for the API under test. The scanner\nuses this to discover endpoints and construct requests.\n", "type": "string"}, "scan_type": {"type": "string", "enum": ["bola"]}, "target_environment_id": {"description": "The target environment to scan.", "type": "string", "format": "uuid"}}, "required": ["target_environment_id", "scan_type", "open_api", "credential_sets"]}
```
