---
title: addressing_advertised
page_id: schema-addressing-advertised-d98be475
path: schemas
description: Prefix advertisement status to the Internet. This field is only not 'null' if on demand is enabled.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# addressing_advertised

Prefix advertisement status to the Internet. This field is only not 'null' if on demand is enabled.

```yaml
{"description": "Prefix advertisement status to the Internet. This field is only not 'null' if on demand is enabled.", "type": "boolean", "example": true, "deprecated": true, "nullable": true, "x-auditable": true, "x-stainless-deprecation-message": "Prefer the [BGP Prefixes API](https://developers.cloudflare.com/api/resources/addressing/subresources/prefixes/subresources/bgp_prefixes/) instead, which allows for advertising multiple BGP routes within a single IP Prefix."}
```
