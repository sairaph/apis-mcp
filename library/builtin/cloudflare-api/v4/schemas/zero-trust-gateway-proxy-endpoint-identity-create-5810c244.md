---
title: zero-trust-gateway_proxy-endpoint-identity-create
page_id: schema-zero-trust-gateway-proxy-endpoint-identity-create-5810c244
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_proxy-endpoint-identity-create

```yaml
{"type": "object", "properties": {"kind": {"description": "The proxy endpoint kind", "type": "string", "example": "identity", "enum": ["identity"], "x-auditable": true}, "name": {"$ref": "#/components/schemas/zero-trust-gateway_name-6"}}, "required": ["kind", "name"]}
```
