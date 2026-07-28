---
title: workers_exports_config_map
page_id: schema-workers-exports-config-map-cac63c3a
path: schemas
description: |-
    Declarative exports for the Worker, keyed by export name. Worker
    entrypoint entries (`type: worker`) carry cache configuration for
    that entrypoint. Durable Object entries (`type: durable-object`)
    are also configured here and drive namespace reconciliation. At
    most 100 entries; class names are at most 128 characters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_exports_config_map

Declarative exports for the Worker, keyed by export name. Worker
entrypoint entries (`type: worker`) carry cache configuration for
that entrypoint. Durable Object entries (`type: durable-object`)
are also configured here and drive namespace reconciliation. At
most 100 entries; class names are at most 128 characters.

```yaml
{"description": "Declarative exports for the Worker, keyed by export name. Worker\nentrypoint entries (`type: worker`) carry cache configuration for\nthat entrypoint. Durable Object entries (`type: durable-object`)\nare also configured here and drive namespace reconciliation. At\nmost 100 entries; class names are at most 128 characters.\n", "type": "object", "example": {"Admin": {"cache": {"enabled": true}, "type": "worker"}, "Counter": {"storage": "sqlite", "type": "durable-object"}, "OldCounter": {"renamed_to": "Counter", "state": "renamed", "type": "durable-object"}, "default": {"cache": {"enabled": false}, "type": "worker"}}, "additionalProperties": {"$ref": "#/components/schemas/workers_export_config"}, "maxProperties": 100}
```
