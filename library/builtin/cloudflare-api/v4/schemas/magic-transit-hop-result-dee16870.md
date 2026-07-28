---
title: magic-transit_hop_result
page_id: schema-magic-transit-hop-result-dee16870
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic-transit_hop_result

```yaml
{"type": "object", "properties": {"nodes": {"description": "An array of node objects.", "type": "array", "items": {"$ref": "#/components/schemas/magic-transit_node_result"}}, "packets_lost": {"$ref": "#/components/schemas/magic-transit_packets_lost"}, "packets_sent": {"$ref": "#/components/schemas/magic-transit_packets_sent"}, "packets_ttl": {"$ref": "#/components/schemas/magic-transit_packets_ttl"}}}
```
