---
title: zero-trust-gateway_ip
page_id: schema-zero-trust-gateway-ip-6f5cd83f
path: schemas
description: Defines the automatically generated IPv6 destination IP assigned to this location. Gateway counts all DNS requests sent to this IP as requests under this location.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_ip

Defines the automatically generated IPv6 destination IP assigned to this location. Gateway counts all DNS requests sent to this IP as requests under this location.

```yaml
{"description": "Defines the automatically generated IPv6 destination IP assigned to this location. Gateway counts all DNS requests sent to this IP as requests under this location.", "type": "string", "example": "2001:0db8:85a3:0000:0000:8a2e:0370:7334", "readOnly": true, "x-auditable": true, "x-stainless-terraform-configurability": "computed"}
```
