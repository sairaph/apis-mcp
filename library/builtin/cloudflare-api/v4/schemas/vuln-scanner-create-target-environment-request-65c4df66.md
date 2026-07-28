---
title: vuln_scanner_create-target-environment-request
page_id: schema-vuln-scanner-create-target-environment-request-65c4df66
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_create-target-environment-request

```yaml
{"type": "object", "properties": {"description": {"description": "Optional description.", "type": "string", "example": "Main production environment", "nullable": true}, "name": {"description": "Human-readable name.", "type": "string", "example": "Production Zone"}, "target": {"$ref": "#/components/schemas/vuln_scanner_target-type"}}, "required": ["name", "target"]}
```
