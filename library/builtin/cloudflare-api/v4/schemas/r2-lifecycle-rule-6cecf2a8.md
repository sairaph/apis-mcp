---
title: r2_lifecycle-rule
page_id: schema-r2-lifecycle-rule-6cecf2a8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_lifecycle-rule

```yaml
{"type": "object", "properties": {"abortMultipartUploadsTransition": {"description": "Transition to abort ongoing multipart uploads.", "type": "object", "properties": {"condition": {"allOf": [{"$ref": "#/components/schemas/r2_lifecycle-age-condition"}]}}}, "conditions": {"description": "Conditions that apply to all transitions of this rule.", "type": "object", "properties": {"prefix": {"description": "Transitions will only apply to objects/uploads in the bucket that start with the given prefix, an empty prefix can be provided to scope rule to all objects/uploads.", "type": "string", "x-auditable": true}}, "required": ["prefix"]}, "deleteObjectsTransition": {"description": "Transition to delete objects.", "type": "object", "properties": {"condition": {"oneOf": [{"$ref": "#/components/schemas/r2_lifecycle-age-condition"}, {"$ref": "#/components/schemas/r2_lifecycle-date-condition"}]}}}, "enabled": {"description": "Whether or not this rule is in effect.", "type": "boolean", "x-auditable": true}, "id": {"description": "Unique identifier for this rule.", "type": "string", "example": "Expire all objects older than 24 hours", "x-auditable": true}, "storageClassTransitions": {"description": "Transitions to change the storage class of objects.", "type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/r2_lifecycle-storage-transition"}]}}}, "required": ["id", "conditions", "enabled"]}
```
