---
title: zero-trust-gateway_action
page_id: schema-zero-trust-gateway-action-85daef70
path: schemas
description: Specify the action to perform when the associated traffic, identity, and device posture expressions either absent or evaluate to `true`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_action

Specify the action to perform when the associated traffic, identity, and device posture expressions either absent or evaluate to `true`.

```yaml
{"description": "Specify the action to perform when the associated traffic, identity, and device posture expressions either absent or evaluate to `true`.", "type": "string", "example": "allow", "enum": ["on", "off", "allow", "block", "scan", "noscan", "safesearch", "ytrestricted", "isolate", "noisolate", "override", "l4_override", "egress", "resolve", "quarantine", "redirect"], "x-auditable": true}
```
