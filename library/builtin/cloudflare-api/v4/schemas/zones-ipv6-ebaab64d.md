---
title: zones_ipv6
page_id: schema-zones-ipv6-ebaab64d
path: schemas
description: Enable IPv6 on all subdomains that are Cloudflare enabled.  (https://support.cloudflare.com/hc/en-us/articles/200168586).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_ipv6

Enable IPv6 on all subdomains that are Cloudflare enabled.  (https://support.cloudflare.com/hc/en-us/articles/200168586).

```yaml
{"description": "Enable IPv6 on all subdomains that are Cloudflare enabled.  (https://support.cloudflare.com/hc/en-us/articles/200168586).", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "ipv6", "enum": ["ipv6"]}, "value": {"$ref": "#/components/schemas/zones_ipv6_value"}}}], "title": "IPv6"}
```
