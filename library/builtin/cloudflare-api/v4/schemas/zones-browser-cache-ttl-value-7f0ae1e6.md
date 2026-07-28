---
title: zones_browser_cache_ttl_value
page_id: schema-zones-browser-cache-ttl-value-7f0ae1e6
path: schemas
description: |-
    Value of the zone setting in seconds.
    Minimum values by plan:
    - Free: 1 second
    - Pro: 1 second
    - Business: 1 second
    - Enterprise: 1 second
    Setting a TTL of 0 is equivalent to selecting `Respect Existing Headers` and is allowed for all plans.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_browser_cache_ttl_value

Value of the zone setting in seconds.
Minimum values by plan:
- Free: 1 second
- Pro: 1 second
- Business: 1 second
- Enterprise: 1 second
Setting a TTL of 0 is equivalent to selecting `Respect Existing Headers` and is allowed for all plans.

```yaml
{"description": "Value of the zone setting in seconds.\nMinimum values by plan:\n- Free: 1 second\n- Pro: 1 second\n- Business: 1 second\n- Enterprise: 1 second\nSetting a TTL of 0 is equivalent to selecting `Respect Existing Headers` and is allowed for all plans.", "type": "integer", "default": 14400, "maximum": 31536000, "minimum": 0}
```
