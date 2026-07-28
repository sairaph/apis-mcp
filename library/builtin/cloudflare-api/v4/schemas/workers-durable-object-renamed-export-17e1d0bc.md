---
title: workers_durable_object_renamed_export
page_id: schema-workers-durable-object-renamed-export-17e1d0bc
path: schemas
description: |-
    A `renamed` tombstone: rewrites the provisioned namespace's class
    name from this map key to `renamed_to`. The source class may stay
    in code during the rollout window (an info notice is emitted).
    `storage`, `transferred_to` and `transfer_from` are not allowed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_durable_object_renamed_export

A `renamed` tombstone: rewrites the provisioned namespace's class
name from this map key to `renamed_to`. The source class may stay
in code during the rollout window (an info notice is emitted).
`storage`, `transferred_to` and `transfer_from` are not allowed.

```yaml
{"description": "A `renamed` tombstone: rewrites the provisioned namespace's class\nname from this map key to `renamed_to`. The source class may stay\nin code during the rollout window (an info notice is emitted).\n`storage`, `transferred_to` and `transfer_from` are not allowed.\n", "type": "object", "properties": {"renamed_to": {"description": "The destination class name. Must differ from the source class\n(the map key) and must be declared as a live (`created`) entry\nin the same `exports` map. Write-only: never present in GET\nresponses.\n", "type": "string", "example": "NewName", "maxLength": 128, "writeOnly": true}, "state": {"description": "Tombstone that renames the namespace's class.", "type": "string", "enum": ["renamed"]}, "type": {"description": "Marks this entry as a Durable Object export.", "type": "string", "enum": ["durable-object"]}}, "additionalProperties": false, "required": ["type", "state", "renamed_to"]}
```
