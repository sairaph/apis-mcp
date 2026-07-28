---
title: zero-trust-gateway_gateway-account-egress-cidr
page_id: schema-zero-trust-gateway-gateway-account-egress-cidr-e3be7c6d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_gateway-account-egress-cidr

```yaml
{"type": "object", "properties": {"geolocation": {"description": "Specify the geographic location of this CIDR pair.", "type": "object", "properties": {"city": {"description": "Specify the city of this egress IP.", "type": "string"}, "country": {"description": "Specify the country of this egress IP.", "type": "string"}}}, "ipv4": {"description": "Specify the IPv4 address of this egress CIDR pair.", "type": "string"}, "ipv4_colo_name": {"description": "Specify the colocation from which this IPv4 address egresses.", "type": "string"}, "ipv6_cidr": {"description": "Specify the IPv6 network address of this egress CIDR pair.", "type": "string"}}, "required": ["ipv4", "ipv4_colo_name", "ipv6_cidr", "geolocation"]}
```
