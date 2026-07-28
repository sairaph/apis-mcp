---
title: workers_multiple_step_migrations
page_id: schema-workers-multiple-step-migrations-fea7923c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_multiple_step_migrations

```yaml
{"allOf": [{"$ref": "#/components/schemas/workers_migration_tag_conditions"}, {"properties": {"steps": {"description": "Migrations to apply in order.", "type": "array", "items": {"$ref": "#/components/schemas/workers_migration_step"}, "writeOnly": true}}, "type": "object"}]}
```
