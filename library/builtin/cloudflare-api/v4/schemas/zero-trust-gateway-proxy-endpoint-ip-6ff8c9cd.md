---
title: zero-trust-gateway_proxy-endpoint-ip
page_id: schema-zero-trust-gateway-proxy-endpoint-ip-6ff8c9cd
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_proxy-endpoint-ip

```yaml
{"type": "object", "properties": {"created_at": {"$ref": "#/components/schemas/zero-trust-gateway_read_only_timestamp"}, "id": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-3"}, "ips": {"$ref": "#/components/schemas/zero-trust-gateway_ips"}, "kind": {"description": "The proxy endpoint kind", "type": "string", "example": "ip", "enum": ["ip"], "x-auditable": true}, "name": {"$ref": "#/components/schemas/zero-trust-gateway_name-6"}, "subdomain": {"$ref": "#/components/schemas/zero-trust-gateway_subdomain-2"}, "updated_at": {"$ref": "#/components/schemas/zero-trust-gateway_read_only_timestamp"}}, "required": ["name", "ips"]}
```
