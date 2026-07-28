---
title: r2_lifecycle-storage-transition
page_id: schema-r2-lifecycle-storage-transition-89aa6d43
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_lifecycle-storage-transition

```yaml
{"type": "object", "properties": {"condition": {"oneOf": [{"$ref": "#/components/schemas/r2_lifecycle-age-condition"}, {"$ref": "#/components/schemas/r2_lifecycle-date-condition"}]}, "storageClass": {"type": "string", "enum": ["InfrequentAccess"], "x-auditable": true}}, "required": ["condition", "storageClass"]}
```
