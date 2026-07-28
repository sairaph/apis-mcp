---
title: magic_interface_address_interconnect
page_id: schema-magic-interface-address-interconnect-ff54515c
path: schemas
description: |-
    The IPv4 interface address for the interconnect. For MPLS Interconnects,
    use a /30 or /31 prefix. For GRE Interconnects, a /29, /30, or /31 prefix
    may be used. A /29 prefix is only allowed for v1.5 interconnects,
    and the address must be the .3 host of the subnet (the fourth address
    overall; the network address is not usable). Select the subnet from RFC 1918
    or the approved link-local ranges.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_interface_address_interconnect

The IPv4 interface address for the interconnect. For MPLS Interconnects,
use a /30 or /31 prefix. For GRE Interconnects, a /29, /30, or /31 prefix
may be used. A /29 prefix is only allowed for v1.5 interconnects,
and the address must be the .3 host of the subnet (the fourth address
overall; the network address is not usable). Select the subnet from RFC 1918
or the approved link-local ranges.

```yaml
{"description": "The IPv4 interface address for the interconnect. For MPLS Interconnects,\nuse a /30 or /31 prefix. For GRE Interconnects, a /29, /30, or /31 prefix\nmay be used. A /29 prefix is only allowed for v1.5 interconnects,\nand the address must be the .3 host of the subnet (the fourth address\noverall; the network address is not usable). Select the subnet from RFC 1918\nor the approved link-local ranges.\n", "type": "string", "example": "192.0.2.3/29", "x-auditable": true}
```
