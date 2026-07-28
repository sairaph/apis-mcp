---
title: firewall_schemas-action
page_id: schema-firewall-schemas-action-1444fc4f
path: schemas
description: The action to apply to a matched request. The `log` action is only available on an Enterprise plan.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_schemas-action

The action to apply to a matched request. The `log` action is only available on an Enterprise plan.

```yaml
{"description": "The action to apply to a matched request. The `log` action is only available on an Enterprise plan.", "type": "string", "example": "block", "enum": ["block", "challenge", "js_challenge", "managed_challenge", "allow", "log", "bypass"]}
```
