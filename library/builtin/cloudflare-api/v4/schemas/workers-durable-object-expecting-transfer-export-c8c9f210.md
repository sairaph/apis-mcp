---
title: workers_durable_object_expecting_transfer_export
page_id: schema-workers-durable-object-expecting-transfer-export-c8c9f210
path: schemas
description: |-
    The target side of a two-phase transfer (`state:
    expecting-transfer`). Declares that this script expects to receive
    a namespace for this class from the `transfer_from` script. This
    is a live entry, not a tombstone: bindings resolve through the
    source's namespace until the source commits with a `transferred`
    tombstone. `storage` and `transfer_from` are required; `renamed_to`
    and `transferred_to` are not allowed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_durable_object_expecting_transfer_export

The target side of a two-phase transfer (`state:
expecting-transfer`). Declares that this script expects to receive
a namespace for this class from the `transfer_from` script. This
is a live entry, not a tombstone: bindings resolve through the
source's namespace until the source commits with a `transferred`
tombstone. `storage` and `transfer_from` are required; `renamed_to`
and `transferred_to` are not allowed.

```yaml
{"description": "The target side of a two-phase transfer (`state:\nexpecting-transfer`). Declares that this script expects to receive\na namespace for this class from the `transfer_from` script. This\nis a live entry, not a tombstone: bindings resolve through the\nsource's namespace until the source commits with a `transferred`\ntombstone. `storage` and `transfer_from` are required; `renamed_to`\nand `transferred_to` are not allowed.\n", "type": "object", "properties": {"container": {"description": "Name of the container (declared in the upload's\n`metadata.containers`) that backs this Durable Object once the\ntransfer settles. Valid only on live entries.\n", "type": "string", "example": "my-container", "maxLength": 128}, "state": {"description": "Target side of a two-phase transfer.", "type": "string", "enum": ["expecting-transfer"]}, "storage": {"$ref": "#/components/schemas/workers_export_storage"}, "transfer_from": {"description": "The source script name to receive the namespace from. Must be\nin the same account and dispatch-namespace context. Present on\nreads for `expecting-transfer` entries.\n", "type": "string", "example": "source-worker", "maxLength": 128}, "type": {"description": "Marks this entry as a Durable Object export.", "type": "string", "enum": ["durable-object"]}}, "additionalProperties": false, "required": ["type", "state", "storage", "transfer_from"]}
```
