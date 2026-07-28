---
title: dos_InfraPrefix
page_id: schema-dos-infraprefix-11b4cdd3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dos_InfraPrefix

```yaml
{"type": "object", "properties": {"comment": {"description": "An optional comment describing the allowlist prefix.", "type": "string", "x-auditable": true}, "created_on": {"description": "The creation timestamp of the allowlist prefix.", "type": "string", "format": "date-time", "x-auditable": true}, "enabled": {"description": "Whether to enable the allowlist prefix into effect. Defaults to false.", "type": "boolean", "x-auditable": true}, "id": {"description": "The unique ID of the allowlist prefix.", "type": "string", "x-auditable": true}, "modified_on": {"description": "The last modification timestamp of the allowlist prefix.", "type": "string", "format": "date-time", "x-auditable": true}, "prefix": {"description": "The allowlist prefix in CIDR format.", "type": "string", "x-auditable": true}}, "additionalProperties": false, "required": ["id", "prefix", "comment", "enabled", "created_on", "modified_on"]}
```
