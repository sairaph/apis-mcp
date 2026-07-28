---
title: vuln_scanner_bola-test-request
page_id: schema-vuln-scanner-bola-test-request-e3438d05
path: schemas
description: HTTP request that was made.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-test-request

HTTP request that was made.

```yaml
{"description": "HTTP request that was made.", "type": "object", "properties": {"body": {"description": "Request body, if any.", "type": "object", "nullable": true}, "credential_set": {"description": "Credential set that was used.", "allOf": [{"$ref": "#/components/schemas/vuln_scanner_bola-test-credential-set"}]}, "header_names": {"description": "Names of headers that were sent.", "type": "array", "items": {"type": "string"}}, "method": {"description": "HTTP method.", "allOf": [{"$ref": "#/components/schemas/vuln_scanner_bola-method"}]}, "url": {"description": "Exact and full URL (including host, query parameters) that was requested.", "type": "string", "format": "uri"}, "variable_captures": {"description": "Variable captures requested for this step.", "type": "array", "items": {"$ref": "#/components/schemas/vuln_scanner_bola-variable-capture"}}}, "required": ["method", "url", "credential_set", "header_names", "variable_captures"]}
```
