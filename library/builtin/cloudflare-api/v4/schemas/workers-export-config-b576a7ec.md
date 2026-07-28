---
title: workers_export_config
page_id: schema-workers-export-config-b576a7ec
path: schemas
description: |-
    A single entry in the `exports` map, keyed by export name (a
    `WorkerEntrypoint` class name, a Durable Object class name, or
    `default` for the Worker's default export). The `type`
    discriminator selects the top-level shape: `worker` entrypoint
    entries may carry `cache` configuration, while `durable-object`
    entries are further refined by the optional `state` field
    (default `created`). Tombstone states (`deleted`, `renamed`,
    `transferred`) express destructive lifecycle operations
    declaratively; `expecting-transfer` is the live target side of a
    transfer. The server validates the exact per-(type, state) field
    combinations; fields not listed for a variant are rejected.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_export_config

A single entry in the `exports` map, keyed by export name (a
`WorkerEntrypoint` class name, a Durable Object class name, or
`default` for the Worker's default export). The `type`
discriminator selects the top-level shape: `worker` entrypoint
entries may carry `cache` configuration, while `durable-object`
entries are further refined by the optional `state` field
(default `created`). Tombstone states (`deleted`, `renamed`,
`transferred`) express destructive lifecycle operations
declaratively; `expecting-transfer` is the live target side of a
transfer. The server validates the exact per-(type, state) field
combinations; fields not listed for a variant are rejected.

```yaml
{"description": "A single entry in the `exports` map, keyed by export name (a\n`WorkerEntrypoint` class name, a Durable Object class name, or\n`default` for the Worker's default export). The `type`\ndiscriminator selects the top-level shape: `worker` entrypoint\nentries may carry `cache` configuration, while `durable-object`\nentries are further refined by the optional `state` field\n(default `created`). Tombstone states (`deleted`, `renamed`,\n`transferred`) express destructive lifecycle operations\ndeclaratively; `expecting-transfer` is the live target side of a\ntransfer. The server validates the exact per-(type, state) field\ncombinations; fields not listed for a variant are rejected.\n", "discriminator": {"mapping": {"durable-object": "#/components/schemas/workers_durable_object_export_variants", "worker": "#/components/schemas/workers_worker_export"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/workers_worker_export"}, {"$ref": "#/components/schemas/workers_durable_object_export_variants"}]}
```
