---
title: r2_rule
page_id: schema-r2-rule-760dc3d3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_rule

```yaml
{"type": "object", "properties": {"actions": {"description": "Array of R2 object actions that will trigger notifications.", "type": "array", "items": {"$ref": "#/components/schemas/r2_r2-action"}, "example": ["PutObject", "CopyObject"], "uniqueItems": true}, "description": {"description": "A description that can be used to identify the event notification rule after creation.", "type": "string", "example": "Notifications from source bucket to queue", "x-auditable": true}, "prefix": {"description": "Notifications will be sent only for objects with this prefix.", "type": "string", "example": "img/", "x-auditable": true}, "suffix": {"description": "Notifications will be sent only for objects with this suffix.", "type": "string", "example": ".jpeg", "x-auditable": true}}, "required": ["actions"]}
```
