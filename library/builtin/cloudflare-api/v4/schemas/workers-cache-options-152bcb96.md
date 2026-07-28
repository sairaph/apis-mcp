---
title: workers_cache_options
page_id: schema-workers-cache-options-152bcb96
path: schemas
description: |-
    Global CacheW configuration for the Worker. When caching is on,
    the platform provisions a `cloudflare.app` zone for the Worker.
    A `type: worker` entry in the `exports` map can override this
    value for a single entrypoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_cache_options

Global CacheW configuration for the Worker. When caching is on,
the platform provisions a `cloudflare.app` zone for the Worker.
A `type: worker` entry in the `exports` map can override this
value for a single entrypoint.

```yaml
{"description": "Global CacheW configuration for the Worker. When caching is on,\nthe platform provisions a `cloudflare.app` zone for the Worker.\nA `type: worker` entry in the `exports` map can override this\nvalue for a single entrypoint.\n", "type": "object", "properties": {"cross_version_cache": {"description": "Whether cached responses are shared across Worker version\nuploads. This is independent of `enabled`. It can stay true\nwhile caching is off, so the preference survives turning\ncaching off and back on.\n", "type": "boolean", "example": true, "default": false, "x-auditable": true}, "enabled": {"description": "Whether caching is enabled for this Worker.", "type": "boolean", "example": true, "default": false, "x-auditable": true}}, "additionalProperties": false, "required": ["enabled"]}
```
