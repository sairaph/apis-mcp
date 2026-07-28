---
title: load-balancing_topology_match
page_id: schema-load-balancing-topology-match-1159e427
path: schemas
description: 'Matches requests by location. Set any combination of `pops`, `countries`, and `regions` (at least one is required); a request matches when its value appears in any populated list (e.g. `regions: ["WNAM"]` with `countries: ["US"]` matches a request in either WNAM or the US).'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_topology_match

Matches requests by location. Set any combination of `pops`, `countries`, and `regions` (at least one is required); a request matches when its value appears in any populated list (e.g. `regions: ["WNAM"]` with `countries: ["US"]` matches a request in either WNAM or the US).

```yaml
{"description": "Matches requests by location. Set any combination of `pops`, `countries`, and `regions` (at least one is required); a request matches when its value appears in any populated list (e.g. `regions: [\"WNAM\"]` with `countries: [\"US\"]` matches a request in either WNAM or the US).", "type": "object", "properties": {"countries": {"description": "A list of ISO 3166-1 alpha-2 country codes. Matches when the request's country is in this list.", "type": "array", "items": {"type": "string", "x-auditable": true}, "maxItems": 1000}, "pops": {"description": "A list of Cloudflare PoP codes. Matches when the request's PoP is in this list.", "type": "array", "items": {"type": "string", "x-auditable": true}, "maxItems": 1000}, "regions": {"description": "A list of Cloudflare region codes (e.g. `WNAM`, `ENAM`, `WEU`). Matches when the request's region is in this list.", "type": "array", "items": {"type": "string", "x-auditable": true}, "maxItems": 1000}}, "x-stainless-terraform-configurability": "computed_optional"}
```
