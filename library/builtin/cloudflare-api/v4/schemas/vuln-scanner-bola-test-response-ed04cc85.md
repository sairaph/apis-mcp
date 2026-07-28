---
title: vuln_scanner_bola-test-response
page_id: schema-vuln-scanner-bola-test-response-ed04cc85
path: schemas
description: HTTP response that was received.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-test-response

HTTP response that was received.

```yaml
{"description": "HTTP response that was received.", "type": "object", "properties": {"body": {"description": "HTTP response body.", "allOf": [{"$ref": "#/components/schemas/vuln_scanner_bola-body-response"}]}, "header_names": {"description": "Names of headers that were received.", "type": "array", "items": {"type": "string"}}, "status": {"description": "HTTP status code.", "type": "integer", "maximum": 65535, "minimum": 0}, "status_text": {"description": "HTTP status text, if available for the status code.", "type": "string", "nullable": true}}, "required": ["status", "header_names", "body"]}
```
