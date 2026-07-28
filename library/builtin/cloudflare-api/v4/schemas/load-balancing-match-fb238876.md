---
title: load-balancing_match
page_id: schema-load-balancing-match-fb238876
path: schemas
description: 'Determines which requests a pool set applies to. Set `topology` to match by location or `default: true` to match all requests; the two are mutually exclusive. A pool set with no `match` matches all requests.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_match

Determines which requests a pool set applies to. Set `topology` to match by location or `default: true` to match all requests; the two are mutually exclusive. A pool set with no `match` matches all requests.

```yaml
{"description": "Determines which requests a pool set applies to. Set `topology` to match by location or `default: true` to match all requests; the two are mutually exclusive. A pool set with no `match` matches all requests.", "type": "object", "properties": {"default": {"description": "When true, matches every request. Cannot be combined with `topology`.", "type": "boolean", "default": false, "x-auditable": true}, "topology": {"$ref": "#/components/schemas/load-balancing_topology_match"}}, "x-stainless-terraform-configurability": "computed_optional"}
```
