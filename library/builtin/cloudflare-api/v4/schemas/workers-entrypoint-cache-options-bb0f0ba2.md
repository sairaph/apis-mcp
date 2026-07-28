---
title: workers_entrypoint_cache_options
page_id: schema-workers-entrypoint-cache-options-bb0f0ba2
path: schemas
description: |-
    CacheW configuration for a single `type: worker` entrypoint.
    When present it overrides the Worker's global
    `cache_options.enabled` for that entrypoint, so a Worker can turn
    caching on globally and opt one entrypoint out, or the reverse.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_entrypoint_cache_options

CacheW configuration for a single `type: worker` entrypoint.
When present it overrides the Worker's global
`cache_options.enabled` for that entrypoint, so a Worker can turn
caching on globally and opt one entrypoint out, or the reverse.

```yaml
{"description": "CacheW configuration for a single `type: worker` entrypoint.\nWhen present it overrides the Worker's global\n`cache_options.enabled` for that entrypoint, so a Worker can turn\ncaching on globally and opt one entrypoint out, or the reverse.\n", "type": "object", "properties": {"enabled": {"description": "Whether caching is enabled for this entrypoint.", "type": "boolean", "example": true, "x-auditable": true}}, "additionalProperties": false, "required": ["enabled"]}
```
