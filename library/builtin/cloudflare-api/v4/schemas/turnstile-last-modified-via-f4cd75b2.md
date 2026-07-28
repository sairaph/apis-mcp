---
title: turnstile_last_modified_via
page_id: schema-turnstile-last-modified-via-f4cd75b2
path: schemas
description: |-
    Origin of the most recent mutation (create, update, delete, or
    secret rotation). Server-derived; not client-settable. Omitted for
    widgets last mutated before this field existed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# turnstile_last_modified_via

Origin of the most recent mutation (create, update, delete, or
secret rotation). Server-derived; not client-settable. Omitted for
widgets last mutated before this field existed.

```yaml
{"description": "Origin of the most recent mutation (create, update, delete, or\nsecret rotation). Server-derived; not client-settable. Omitted for\nwidgets last mutated before this field existed.\n", "type": "string", "example": "dashboard", "enum": ["wrangler", "dashboard", "spin", "api", "unknown"], "readOnly": true, "x-auditable": true}
```
