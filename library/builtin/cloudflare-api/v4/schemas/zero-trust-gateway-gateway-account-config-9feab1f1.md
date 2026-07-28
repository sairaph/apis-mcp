---
title: zero-trust-gateway_gateway_account_config
page_id: schema-zero-trust-gateway-gateway-account-config-9feab1f1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_gateway_account_config

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_api-response-single"}, {"properties": {"result": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_gateway-account-settings"}, {"properties": {"created_at": {"$ref": "#/components/schemas/zero-trust-gateway_read_only_timestamp"}, "updated_at": {"$ref": "#/components/schemas/zero-trust-gateway_read_only_timestamp"}}, "type": "object"}]}}, "type": "object"}]}
```
