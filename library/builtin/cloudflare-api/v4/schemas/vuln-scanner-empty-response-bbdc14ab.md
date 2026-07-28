---
title: vuln_scanner_empty-response
page_id: schema-vuln-scanner-empty-response-bbdc14ab
path: schemas
description: |-
    Standard response envelope with a `null` result. Operations
    that produce no result body return this response.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_empty-response

Standard response envelope with a `null` result. Operations
that produce no result body return this response.

```yaml
{"description": "Standard response envelope with a `null` result. Operations\nthat produce no result body return this response.\n", "allOf": [{"$ref": "#/components/schemas/vuln_scanner_api-response-common"}, {"properties": {"result": {"type": "object", "nullable": true}, "result_info": {"type": "object", "nullable": true}}, "type": "object"}]}
```
