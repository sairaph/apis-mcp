---
title: firewall_bypass
page_id: schema-firewall-bypass-cfb1ec87
path: schemas
description: Criteria specifying when the current rate limit should be bypassed. You can specify that the rate limit should not apply to one or more URLs.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_bypass

Criteria specifying when the current rate limit should be bypassed. You can specify that the rate limit should not apply to one or more URLs.

```yaml
{"description": "Criteria specifying when the current rate limit should be bypassed. You can specify that the rate limit should not apply to one or more URLs.", "type": "array", "items": {"properties": {"name": {"type": "string", "example": "url", "enum": ["url"]}, "value": {"description": "The URL to bypass.", "type": "string", "example": "api.example.com/*"}}, "type": "object"}}
```
