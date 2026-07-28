---
title: infra_IPInfo
page_id: schema-infra-ipinfo-464a4ea7
path: schemas
description: The IPv4/IPv6 address that identifies where to reach a target
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# infra_IPInfo

The IPv4/IPv6 address that identifies where to reach a target

```yaml
{"description": "The IPv4/IPv6 address that identifies where to reach a target", "type": "object", "properties": {"ipv4": {"description": "The target's IPv4 address", "type": "object", "properties": {"ip_addr": {"description": "IP address of the target", "type": "string", "example": "187.26.29.249", "x-auditable": true}, "virtual_network_id": {"description": "(optional) Private virtual network identifier for the target. If omitted, the default virtual network ID will be used.", "type": "string", "format": "uuid", "example": "c77b744e-acc8-428f-9257-6878c046ed55", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}}}, "ipv6": {"description": "The target's IPv6 address", "type": "object", "properties": {"ip_addr": {"description": "IP address of the target", "type": "string", "example": "64c0:64e8:f0b4:8dbf:7104:72b0:ec8f:f5e0", "x-auditable": true}, "virtual_network_id": {"description": "(optional) Private virtual network identifier for the target. If omitted, the default virtual network ID will be used.", "type": "string", "format": "uuid", "example": "c77b744e-acc8-428f-9257-6878c046ed55", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}}}}}
```
