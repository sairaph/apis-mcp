---
title: vuln_scanner_update-target-environment-request
page_id: schema-vuln-scanner-update-target-environment-request-c94765e7
path: schemas
description: Full replacement.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_update-target-environment-request

Full replacement.

```yaml
{"description": "Full replacement.", "type": "object", "properties": {"description": {"description": "Optional description.", "type": "string", "example": "Main production environment", "nullable": true}, "name": {"description": "Human-readable name.", "type": "string", "example": "Production Zone"}, "target": {"$ref": "#/components/schemas/vuln_scanner_target-type"}}, "required": ["name", "target"]}
```
