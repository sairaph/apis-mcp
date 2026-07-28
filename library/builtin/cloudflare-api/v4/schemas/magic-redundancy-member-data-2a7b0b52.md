---
title: magic_redundancy_member_data
page_id: schema-magic-redundancy-member-data-2a7b0b52
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_redundancy_member_data

```yaml
{"type": "object", "properties": {"data": {"description": "Full tunnel object as returned by the corresponding endpoint", "type": "object"}, "type": {"description": "Tunnel type", "type": "string", "enum": ["gre", "ipsec", "cni"]}}, "required": ["type", "data"]}
```
