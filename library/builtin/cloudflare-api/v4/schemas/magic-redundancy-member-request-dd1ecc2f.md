---
title: magic_redundancy_member_request
page_id: schema-magic-redundancy-member-request-dd1ecc2f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_redundancy_member_request

```yaml
{"type": "object", "properties": {"id": {"description": "UUID of the tunnel or interconnect", "type": "string", "example": "a1b2c3d4e5f647890a1b2c3d4e5f6789", "maxLength": 36, "pattern": "^[a-fA-F0-9\\-]{32,36}$"}, "type": {"description": "Tunnel type: gre, ipsec, or cni", "type": "string", "enum": ["gre", "ipsec", "cni"]}}, "required": ["id", "type"]}
```
