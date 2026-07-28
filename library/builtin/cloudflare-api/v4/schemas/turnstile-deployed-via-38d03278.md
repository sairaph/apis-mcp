---
title: turnstile_deployed_via
page_id: schema-turnstile-deployed-via-38d03278
path: schemas
description: |-
    Origin that created this widget, recorded at creation time and
    immutable afterward. Server-derived from the create request; not
    client-settable. Omitted from the response for widgets created
    before this field existed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# turnstile_deployed_via

Origin that created this widget, recorded at creation time and
immutable afterward. Server-derived from the create request; not
client-settable. Omitted from the response for widgets created
before this field existed.

```yaml
{"description": "Origin that created this widget, recorded at creation time and\nimmutable afterward. Server-derived from the create request; not\nclient-settable. Omitted from the response for widgets created\nbefore this field existed.\n", "type": "string", "example": "wrangler", "enum": ["wrangler", "dashboard", "spin", "api", "unknown"], "readOnly": true, "x-auditable": true}
```
