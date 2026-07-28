---
title: magic_routed_subnet
page_id: schema-magic-routed-subnet-bf5b1aa1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_routed_subnet

```yaml
{"type": "object", "properties": {"nat": {"$ref": "#/components/schemas/magic_nat"}, "next_hop": {"$ref": "#/components/schemas/magic_ip-address"}, "prefix": {"$ref": "#/components/schemas/magic_cidr"}}, "required": ["prefix", "next_hop"]}
```
