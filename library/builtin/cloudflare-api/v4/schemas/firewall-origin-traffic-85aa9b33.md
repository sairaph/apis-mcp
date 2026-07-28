---
title: firewall_origin_traffic
page_id: schema-firewall-origin-traffic-85aa9b33
path: schemas
description: |-
    When true, only the uncached traffic served from your origin servers will count towards rate limiting. In this case, any cached traffic served by Cloudflare will not count towards rate limiting. This field is optional.
    Notes: This field is deprecated. Instead, use response headers and set "origin_traffic" to "false" to avoid legacy behaviour interacting with the "response_headers" property.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_origin_traffic

When true, only the uncached traffic served from your origin servers will count towards rate limiting. In this case, any cached traffic served by Cloudflare will not count towards rate limiting. This field is optional.
Notes: This field is deprecated. Instead, use response headers and set "origin_traffic" to "false" to avoid legacy behaviour interacting with the "response_headers" property.

```yaml
{"description": "When true, only the uncached traffic served from your origin servers will count towards rate limiting. In this case, any cached traffic served by Cloudflare will not count towards rate limiting. This field is optional.\nNotes: This field is deprecated. Instead, use response headers and set \"origin_traffic\" to \"false\" to avoid legacy behaviour interacting with the \"response_headers\" property.", "type": "boolean", "x-auditable": true}
```
