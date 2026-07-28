---
title: zero-trust-gateway_proxy-endpoints
page_id: schema-zero-trust-gateway-proxy-endpoints-1b96a123
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_proxy-endpoints

```yaml
{"discriminator": {"mapping": {"identity": "#/components/schemas/zero-trust-gateway_proxy-endpoint-identity", "ip": "#/components/schemas/zero-trust-gateway_proxy-endpoint-ip"}, "propertyName": "kind"}, "oneOf": [{"$ref": "#/components/schemas/zero-trust-gateway_proxy-endpoint-ip"}, {"$ref": "#/components/schemas/zero-trust-gateway_proxy-endpoint-identity"}]}
```
