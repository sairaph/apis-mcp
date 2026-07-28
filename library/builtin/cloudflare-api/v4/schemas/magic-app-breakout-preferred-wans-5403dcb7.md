---
title: magic_app_breakout_preferred_wans
page_id: schema-magic-app-breakout-preferred-wans-5403dcb7
path: schemas
description: WAN interfaces to prefer over default WANs, highest-priority first. Can only be specified for breakout rules (breakout must be true).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_app_breakout_preferred_wans

WAN interfaces to prefer over default WANs, highest-priority first. Can only be specified for breakout rules (breakout must be true).

```yaml
{"description": "WAN interfaces to prefer over default WANs, highest-priority first. Can only be specified for breakout rules (breakout must be true).", "type": "array", "items": {"$ref": "#/components/schemas/magic_identifier"}}
```
