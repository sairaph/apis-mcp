---
title: workers_durable_object_deleted_export
page_id: schema-workers-durable-object-deleted-export-7cea9244
path: schemas
description: |-
    A `deleted` tombstone: retires the provisioned namespace for this
    class and all of its data. The class must be absent from the
    uploaded code and no other Worker in the account may bind to the
    namespace, otherwise the deploy is rejected. No other fields are
    allowed. Deletion is irreversible.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_durable_object_deleted_export

A `deleted` tombstone: retires the provisioned namespace for this
class and all of its data. The class must be absent from the
uploaded code and no other Worker in the account may bind to the
namespace, otherwise the deploy is rejected. No other fields are
allowed. Deletion is irreversible.

```yaml
{"description": "A `deleted` tombstone: retires the provisioned namespace for this\nclass and all of its data. The class must be absent from the\nuploaded code and no other Worker in the account may bind to the\nnamespace, otherwise the deploy is rejected. No other fields are\nallowed. Deletion is irreversible.\n", "type": "object", "properties": {"state": {"description": "Tombstone that deletes the namespace.", "type": "string", "enum": ["deleted"]}, "type": {"description": "Marks this entry as a Durable Object export.", "type": "string", "enum": ["durable-object"]}}, "additionalProperties": false, "required": ["type", "state"]}
```
