---
title: r2_bucket-lock-rule
page_id: schema-r2-bucket-lock-rule-a98d9017
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_bucket-lock-rule

```yaml
{"type": "object", "properties": {"condition": {"oneOf": [{"$ref": "#/components/schemas/r2_lock-rule-age-condition"}, {"$ref": "#/components/schemas/r2_lock-rule-date-condition"}, {"$ref": "#/components/schemas/r2_lock-rule-indefinite-condition"}]}, "enabled": {"description": "Whether or not this rule is in effect.", "type": "boolean", "x-auditable": true}, "id": {"description": "Unique identifier for this rule.", "type": "string", "example": "Lock all objects for 24 hours", "x-auditable": true}, "prefix": {"description": "Rule will only apply to objects/uploads in the bucket that start with the given prefix, an empty prefix can be provided to scope rule to all objects/uploads.", "type": "string", "x-auditable": true}}, "required": ["id", "condition", "enabled"]}
```
