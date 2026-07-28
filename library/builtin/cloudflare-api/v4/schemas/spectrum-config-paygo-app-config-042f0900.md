---
title: spectrum-config_paygo_app_config
page_id: schema-spectrum-config-paygo-app-config-042f0900
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# spectrum-config_paygo_app_config

```yaml
{"allOf": [{"$ref": "#/components/schemas/spectrum-config_base_app_config"}, {"properties": {"dns": {"$ref": "#/components/schemas/spectrum-config_dns"}, "origin_direct": {"$ref": "#/components/schemas/spectrum-config_origin_direct"}, "protocol": {"$ref": "#/components/schemas/spectrum-config_protocol"}}, "required": ["protocol", "dns"], "type": "object"}]}
```
