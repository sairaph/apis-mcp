---
title: dos_Prefix
page_id: schema-dos-prefix-434e88dd
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dos_Prefix

```yaml
{"type": "object", "properties": {"comment": {"description": "A comment describing the prefix.", "type": "string", "x-auditable": true}, "created_on": {"description": "The creation timestamp of the prefix.", "type": "string", "format": "date-time", "x-auditable": true}, "excluded": {"description": "Whether to exclude the prefix from protection.", "type": "boolean", "x-auditable": true}, "id": {"description": "The unique ID of the prefix.", "type": "string", "x-auditable": true}, "modified_on": {"description": "The last modification timestamp of the prefix.", "type": "string", "format": "date-time", "x-auditable": true}, "prefix": {"description": "The prefix in CIDR format.", "type": "string", "example": "192.0.2.0/24", "x-auditable": true}}, "additionalProperties": false, "required": ["id", "prefix", "comment", "excluded", "created_on", "modified_on"]}
```
