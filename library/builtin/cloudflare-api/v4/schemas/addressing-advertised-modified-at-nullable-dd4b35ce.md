---
title: addressing_advertised_modified_at_nullable
page_id: schema-addressing-advertised-modified-at-nullable-dd4b35ce
path: schemas
description: Last time the advertisement status was changed. This field is only not 'null' if on demand is enabled.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# addressing_advertised_modified_at_nullable

Last time the advertisement status was changed. This field is only not 'null' if on demand is enabled.

```yaml
{"description": "Last time the advertisement status was changed. This field is only not 'null' if on demand is enabled.", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00.12345Z", "deprecated": true, "nullable": true, "x-auditable": true, "x-stainless-deprecation-message": "Prefer the [BGP Prefixes API](https://developers.cloudflare.com/api/resources/addressing/subresources/prefixes/subresources/bgp_prefixes/) instead, which allows for advertising multiple BGP routes within a single IP Prefix."}
```
