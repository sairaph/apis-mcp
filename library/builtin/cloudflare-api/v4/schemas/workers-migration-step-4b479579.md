---
title: workers_migration_step
page_id: schema-workers-migration-step-4b479579
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_migration_step

```yaml
{"type": "object", "properties": {"deleted_classes": {"description": "A list of classes to delete Durable Object namespaces from.", "type": "array", "items": {"type": "string", "x-auditable": true}, "writeOnly": true}, "new_classes": {"description": "A list of classes to create Durable Object namespaces from.", "type": "array", "items": {"type": "string", "x-auditable": true}, "writeOnly": true}, "new_sqlite_classes": {"description": "A list of classes to create Durable Object namespaces with SQLite from.", "type": "array", "items": {"type": "string", "x-auditable": true}, "writeOnly": true}, "renamed_classes": {"description": "A list of classes with Durable Object namespaces that were renamed.", "type": "array", "items": {"properties": {"from": {"type": "string", "writeOnly": true, "x-auditable": true}, "to": {"type": "string", "writeOnly": true, "x-auditable": true}}, "type": "object"}, "writeOnly": true}, "transferred_classes": {"description": "A list of transfers for Durable Object namespaces from a different Worker and class to a class defined in this Worker.", "type": "array", "items": {"properties": {"from": {"type": "string", "writeOnly": true, "x-auditable": true}, "from_script": {"type": "string", "writeOnly": true, "x-auditable": true}, "to": {"type": "string", "writeOnly": true, "x-auditable": true}}, "type": "object"}, "writeOnly": true}}}
```
