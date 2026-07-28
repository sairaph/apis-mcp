---
title: zero-trust-gateway_endpoints
page_id: schema-zero-trust-gateway-endpoints-044817e0
path: schemas
description: Configure the destination endpoints for this location.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_endpoints

Configure the destination endpoints for this location.

```yaml
{"description": "Configure the destination endpoints for this location.", "type": "object", "properties": {"doh": {"$ref": "#/components/schemas/zero-trust-gateway_doh_endpoint"}, "dot": {"$ref": "#/components/schemas/zero-trust-gateway_dot_endpoint"}, "ipv4": {"$ref": "#/components/schemas/zero-trust-gateway_ipv4_endpoint"}, "ipv6": {"$ref": "#/components/schemas/zero-trust-gateway_ipv6_endpoint"}}, "nullable": true, "required": ["ipv4", "ipv6", "doh", "dot"], "x-stainless-terraform-configurability": "optional"}
```
