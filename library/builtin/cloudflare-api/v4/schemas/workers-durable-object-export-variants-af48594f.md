---
title: workers_durable_object_export_variants
page_id: schema-workers-durable-object-export-variants-af48594f
path: schemas
description: |-
    Any Durable Object export entry (`type: durable-object`), across
    all lifecycle states. The optional `state` field selects the
    variant (default `created`); the tombstone states (`deleted`,
    `renamed`, `transferred`) and the live transfer target
    (`expecting-transfer`) are the alternatives. The server validates
    the exact per-state field combinations; fields not listed for a
    variant are rejected.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_durable_object_export_variants

Any Durable Object export entry (`type: durable-object`), across
all lifecycle states. The optional `state` field selects the
variant (default `created`); the tombstone states (`deleted`,
`renamed`, `transferred`) and the live transfer target
(`expecting-transfer`) are the alternatives. The server validates
the exact per-state field combinations; fields not listed for a
variant are rejected.

```yaml
{"description": "Any Durable Object export entry (`type: durable-object`), across\nall lifecycle states. The optional `state` field selects the\nvariant (default `created`); the tombstone states (`deleted`,\n`renamed`, `transferred`) and the live transfer target\n(`expecting-transfer`) are the alternatives. The server validates\nthe exact per-state field combinations; fields not listed for a\nvariant are rejected.\n", "oneOf": [{"$ref": "#/components/schemas/workers_durable_object_export"}, {"$ref": "#/components/schemas/workers_durable_object_deleted_export"}, {"$ref": "#/components/schemas/workers_durable_object_renamed_export"}, {"$ref": "#/components/schemas/workers_durable_object_transferred_export"}, {"$ref": "#/components/schemas/workers_durable_object_expecting_transfer_export"}]}
```
