---
title: nsc_Cni
page_id: schema-nsc-cni-b4746e13
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# nsc_Cni

```yaml
{"type": "object", "properties": {"account": {"$ref": "#/components/schemas/nsc_AccountTag"}, "bgp": {"$ref": "#/components/schemas/nsc_BgpControl"}, "cust_ip": {"description": "Customer end of the point-to-point link\n\nThis should always be inside the same prefix as `p2p_ip`.", "type": "string", "format": "A.B.C.D/N", "example": "192.168.3.4/31"}, "id": {"type": "string", "format": "uuid"}, "interconnect": {"description": "Interconnect identifier hosting this CNI", "type": "string"}, "magic": {"$ref": "#/components/schemas/nsc_MagicSettings"}, "p2p_ip": {"description": "Cloudflare end of the point-to-point link", "type": "string", "format": "A.B.C.D/N", "example": "192.168.3.4/31"}}, "required": ["id", "interconnect", "account", "p2p_ip", "cust_ip", "magic"]}
```
