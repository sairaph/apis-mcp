---
title: magic-transit_colo_result
page_id: schema-magic-transit-colo-result-30abd7f9
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic-transit_colo_result

```yaml
{"type": "object", "properties": {"colo": {"$ref": "#/components/schemas/magic-transit_colo"}, "error": {"$ref": "#/components/schemas/magic-transit_error"}, "hops": {"type": "array", "items": {"$ref": "#/components/schemas/magic-transit_hop_result"}}, "target_summary": {"$ref": "#/components/schemas/magic-transit_target_summary"}, "traceroute_time_ms": {"$ref": "#/components/schemas/magic-transit_traceroute_time_ms"}}}
```
