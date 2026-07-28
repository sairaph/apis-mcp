---
title: turnstile_clearance_level
page_id: schema-turnstile-clearance-level-11b3fa8f
path: schemas
description: |-
    If Turnstile is embedded on a Cloudflare site and the widget should grant challenge clearance,
    this setting can determine the clearance level to be set
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# turnstile_clearance_level

If Turnstile is embedded on a Cloudflare site and the widget should grant challenge clearance,
this setting can determine the clearance level to be set

```yaml
{"description": "If Turnstile is embedded on a Cloudflare site and the widget should grant challenge clearance,\nthis setting can determine the clearance level to be set\n", "type": "string", "example": "interactive", "enum": ["no_clearance", "jschallenge", "managed", "interactive"], "x-auditable": true}
```
