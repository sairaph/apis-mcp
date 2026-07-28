---
title: addressing_on_demand_enabled
page_id: schema-addressing-on-demand-enabled-b61a1dd9
path: schemas
description: Whether advertisement of the prefix to the Internet may be dynamically enabled or disabled.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# addressing_on_demand_enabled

Whether advertisement of the prefix to the Internet may be dynamically enabled or disabled.

```yaml
{"description": "Whether advertisement of the prefix to the Internet may be dynamically enabled or disabled.", "type": "boolean", "example": true, "deprecated": true, "x-auditable": true, "x-stainless-deprecation-message": "Prefer the [BGP Prefixes API](https://developers.cloudflare.com/api/resources/addressing/subresources/prefixes/subresources/bgp_prefixes/) instead, which allows for advertising multiple BGP routes within a single IP Prefix."}
```
