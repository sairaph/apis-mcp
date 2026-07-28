---
title: vuln_scanner_target-environment
page_id: schema-vuln-scanner-target-environment-a9262b8f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_target-environment

```yaml
{"type": "object", "properties": {"description": {"description": "Optional description providing additional context.", "type": "string", "example": "Main production environment", "nullable": true}, "id": {"description": "Target environment identifier.", "type": "string", "format": "uuid", "x-auditable": true}, "name": {"description": "Human-readable name.", "type": "string", "example": "Production Zone", "x-auditable": true}, "target": {"$ref": "#/components/schemas/vuln_scanner_target-type"}}, "required": ["id", "name", "target"]}
```
