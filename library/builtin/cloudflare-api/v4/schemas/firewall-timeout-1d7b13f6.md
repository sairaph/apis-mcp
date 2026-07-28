---
title: firewall_timeout
page_id: schema-firewall-timeout-1d7b13f6
path: schemas
description: |-
    The time in seconds during which Cloudflare will perform the mitigation action. Must be an integer value greater than or equal to the period.
    Notes: If "mode" is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the zone's Challenge Passage time and you should not provide this value.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_timeout

The time in seconds during which Cloudflare will perform the mitigation action. Must be an integer value greater than or equal to the period.
Notes: If "mode" is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the zone's Challenge Passage time and you should not provide this value.

```yaml
{"description": "The time in seconds during which Cloudflare will perform the mitigation action. Must be an integer value greater than or equal to the period.\nNotes: If \"mode\" is \"challenge\", \"managed_challenge\", or \"js_challenge\", Cloudflare will use the zone's Challenge Passage time and you should not provide this value.", "type": "number", "example": 86400, "maximum": 86400, "minimum": 1, "x-auditable": true}
```
