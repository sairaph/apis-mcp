---
title: zero-trust-gateway_ips
page_id: schema-zero-trust-gateway-ips-215e4a03
path: schemas
description: Specify the list of CIDRs to restrict ingress connections.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_ips

Specify the list of CIDRs to restrict ingress connections.

```yaml
{"description": "Specify the list of CIDRs to restrict ingress connections.", "type": "array", "items": {"description": "Specify an IPv4 or IPv6 CIDR. Limit IPv6 to a maximum of /109 and IPv4 to a maximum of /25.", "example": "192.0.2.1/32", "type": "string", "x-auditable": true}}
```
