---
title: moq_upstreams
page_id: schema-moq-upstreams-73a64ebd
path: schemas
description: |-
    Upstreams are external MOQT server publishers that a relay falls back
    to when it has no local publisher for a requested namespace/track.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# moq_upstreams

Upstreams are external MOQT server publishers that a relay falls back
to when it has no local publisher for a requested namespace/track.

```yaml
{"description": "Upstreams are external MOQT server publishers that a relay falls back\nto when it has no local publisher for a requested namespace/track.\n", "type": "object", "properties": {"enabled": {"type": "boolean", "default": false}, "upstreams": {"description": "Ordered list of upstream MOQT server publishers. Each entry is an\nobject (not a bare string) so per-upstream configuration can be\nadded in the future without another breaking change.\n", "type": "array", "items": {"$ref": "#/components/schemas/moq_upstream"}, "default": []}}}
```
