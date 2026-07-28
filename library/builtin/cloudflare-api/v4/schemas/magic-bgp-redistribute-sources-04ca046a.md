---
title: magic_bgp_redistribute_sources
page_id: schema-magic-bgp-redistribute-sources-04ca046a
path: schemas
description: Per-source toggles controlling which route sources are redistributed into BGP. Each property enables redistribution for one route source.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_bgp_redistribute_sources

Per-source toggles controlling which route sources are redistributed into BGP. Each property enables redistribution for one route source.

```yaml
{"description": "Per-source toggles controlling which route sources are redistributed into BGP. Each property enables redistribution for one route source.", "type": "object", "properties": {"static:wan": {"description": "Redistribute static WAN routes into BGP", "type": "boolean"}}, "required": ["static:wan"]}
```
