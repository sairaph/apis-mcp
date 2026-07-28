---
title: workers_durable_object_export
page_id: schema-workers-durable-object-export-70de61b2
path: schemas
description: |-
    A live Durable Object export (`state: created`, the default). The
    platform auto-provisions the namespace on first deploy, matches it
    on subsequent deploys, and never mutates or deletes it as a side
    effect of a code-only change. `storage` is required; `renamed_to`,
    `transferred_to` and `transfer_from` are not allowed on a live
    entry.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_durable_object_export

A live Durable Object export (`state: created`, the default). The
platform auto-provisions the namespace on first deploy, matches it
on subsequent deploys, and never mutates or deletes it as a side
effect of a code-only change. `storage` is required; `renamed_to`,
`transferred_to` and `transfer_from` are not allowed on a live
entry.

```yaml
{"description": "A live Durable Object export (`state: created`, the default). The\nplatform auto-provisions the namespace on first deploy, matches it\non subsequent deploys, and never mutates or deletes it as a side\neffect of a code-only change. `storage` is required; `renamed_to`,\n`transferred_to` and `transfer_from` are not allowed on a live\nentry.\n", "type": "object", "properties": {"container": {"description": "Name of the container (declared in the upload's\n`metadata.containers`) that backs this Durable Object. When\nset, the namespace is container-enabled. Valid only on live\nentries.\n", "type": "string", "example": "my-container", "maxLength": 128}, "state": {"description": "Live export. May be omitted; defaults to `created`.", "type": "string", "default": "created", "enum": ["created"]}, "storage": {"$ref": "#/components/schemas/workers_export_storage"}, "type": {"description": "Marks this entry as a Durable Object export.", "type": "string", "enum": ["durable-object"]}}, "additionalProperties": false, "required": ["type", "storage"]}
```
