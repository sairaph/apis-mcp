---
title: workers_worker_export
page_id: schema-workers-worker-export-8ddeb7ca
path: schemas
description: |-
    A named Worker entrypoint export (`type: worker`). Worker
    entrypoints are always live (`state: created`) and carry no
    storage or lifecycle fields. The optional `cache` block overrides
    the Worker's global `cache_options.enabled` for this entrypoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_worker_export

A named Worker entrypoint export (`type: worker`). Worker
entrypoints are always live (`state: created`) and carry no
storage or lifecycle fields. The optional `cache` block overrides
the Worker's global `cache_options.enabled` for this entrypoint.

```yaml
{"description": "A named Worker entrypoint export (`type: worker`). Worker\nentrypoints are always live (`state: created`) and carry no\nstorage or lifecycle fields. The optional `cache` block overrides\nthe Worker's global `cache_options.enabled` for this entrypoint.\n", "type": "object", "properties": {"cache": {"description": "Cache override for this entrypoint. Overrides the Worker's\nglobal `cache_options.enabled` for this entrypoint only.\n", "allOf": [{"$ref": "#/components/schemas/workers_entrypoint_cache_options"}]}, "state": {"description": "Live export. May be omitted; defaults to `created`.", "type": "string", "default": "created", "enum": ["created"]}, "type": {"description": "Marks this entry as a Worker entrypoint export.", "type": "string", "enum": ["worker"]}}, "additionalProperties": false, "required": ["type"]}
```
