---
title: dos_NewInfraPrefix
page_id: schema-dos-newinfraprefix-3e834a7e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dos_NewInfraPrefix

```yaml
{"type": "object", "properties": {"comment": {"description": "An comment describing the allowlist prefix.", "type": "string", "x-auditable": true}, "enabled": {"description": "Whether to enable the allowlist prefix into effect.", "type": "boolean", "x-auditable": true}, "prefix": {"description": "The allowlist prefix to add in CIDR format.", "type": "string", "x-auditable": true}}, "additionalProperties": false, "required": ["prefix", "comment", "enabled"]}
```
