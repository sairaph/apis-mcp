---
title: dos_NewPrefix
page_id: schema-dos-newprefix-597e8434
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dos_NewPrefix

```yaml
{"type": "object", "properties": {"comment": {"description": "A comment describing the prefix.", "type": "string", "x-auditable": true}, "excluded": {"description": "Whether to exclude the prefix from protection.", "type": "boolean", "x-auditable": true}, "prefix": {"description": "The prefix to add in CIDR format.", "type": "string", "example": "192.0.2.0/24", "x-auditable": true}}, "additionalProperties": false, "required": ["prefix", "comment", "excluded"]}
```
