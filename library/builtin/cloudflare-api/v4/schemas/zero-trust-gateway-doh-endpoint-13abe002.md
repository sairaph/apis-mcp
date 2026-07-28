---
title: zero-trust-gateway_doh_endpoint
page_id: schema-zero-trust-gateway-doh-endpoint-13abe002
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_doh_endpoint

```yaml
{"type": "object", "properties": {"enabled": {"description": "Indicate whether the DOH endpoint is enabled for this location.", "type": "boolean", "example": true, "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "networks": {"$ref": "#/components/schemas/zero-trust-gateway_ip_networks"}, "require_token": {"description": "Specify whether the DOH endpoint requires user identity authentication.", "type": "boolean", "example": true, "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}}}
```
