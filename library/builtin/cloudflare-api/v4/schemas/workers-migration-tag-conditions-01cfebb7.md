---
title: workers_migration_tag_conditions
page_id: schema-workers-migration-tag-conditions-01cfebb7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_migration_tag_conditions

```yaml
{"type": "object", "properties": {"new_tag": {"description": "Tag to set as the latest migration tag.", "type": "string", "example": "v2", "writeOnly": true, "x-auditable": true}, "old_tag": {"description": "Tag used to verify against the latest migration tag for this Worker. If they don't match, the upload is rejected.", "type": "string", "example": "v1", "writeOnly": true, "x-auditable": true}}}
```
