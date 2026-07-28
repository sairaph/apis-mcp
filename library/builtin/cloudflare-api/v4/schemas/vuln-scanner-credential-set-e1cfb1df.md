---
title: vuln_scanner_credential-set
page_id: schema-vuln-scanner-credential-set-e1cfb1df
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_credential-set

```yaml
{"type": "object", "properties": {"id": {"description": "Credential set identifier.", "type": "string", "format": "uuid", "x-auditable": true}, "name": {"description": "Human-readable name.", "type": "string", "example": "Production API credentials", "x-auditable": true}}, "required": ["id", "name"]}
```
