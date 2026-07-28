---
title: rum_modify-rules-request
page_id: schema-rum-modify-rules-request-dcbdcda0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rum_modify-rules-request

```yaml
{"type": "object", "properties": {"delete_rules": {"description": "A list of rule identifiers to delete.", "type": "array", "items": {"$ref": "#/components/schemas/rum_rule_identifier"}}, "rules": {"description": "A list of rules to create or update.", "type": "array", "items": {"properties": {"host": {"type": "string", "example": "example.com"}, "id": {"$ref": "#/components/schemas/rum_rule_identifier"}, "inclusive": {"type": "boolean", "example": true}, "is_paused": {"type": "boolean", "example": false}, "paths": {"type": "array", "items": {"type": "string"}, "example": ["*"]}}, "type": "object"}}}}
```
