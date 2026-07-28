---
title: workers_durable_object_transferred_export
page_id: schema-workers-durable-object-transferred-export-297e51a4
path: schemas
description: |-
    A `transferred` tombstone (source side of a two-phase transfer):
    hands ownership of the provisioned namespace to another script in
    the same account, named by `transferred_to`. The target must have
    already deployed a matching `expecting-transfer` entry. The source
    class may stay in code during the rollout window (an info notice
    is emitted). `storage`, `renamed_to` and `transfer_from` are not
    allowed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_durable_object_transferred_export

A `transferred` tombstone (source side of a two-phase transfer):
hands ownership of the provisioned namespace to another script in
the same account, named by `transferred_to`. The target must have
already deployed a matching `expecting-transfer` entry. The source
class may stay in code during the rollout window (an info notice
is emitted). `storage`, `renamed_to` and `transfer_from` are not
allowed.

```yaml
{"description": "A `transferred` tombstone (source side of a two-phase transfer):\nhands ownership of the provisioned namespace to another script in\nthe same account, named by `transferred_to`. The target must have\nalready deployed a matching `expecting-transfer` entry. The source\nclass may stay in code during the rollout window (an info notice\nis emitted). `storage`, `renamed_to` and `transfer_from` are not\nallowed.\n", "type": "object", "properties": {"state": {"description": "Tombstone that transfers the namespace to another script.", "type": "string", "enum": ["transferred"]}, "transferred_to": {"description": "The destination script name. Must be in the same account and\nthe same dispatch-namespace context (or both non-dispatch).\nCross-dispatch-namespace transfers are rejected. Write-only:\nnever present in GET responses.\n", "type": "string", "example": "target-worker", "maxLength": 128, "writeOnly": true}, "type": {"description": "Marks this entry as a Durable Object export.", "type": "string", "enum": ["durable-object"]}}, "additionalProperties": false, "required": ["type", "state", "transferred_to"]}
```
