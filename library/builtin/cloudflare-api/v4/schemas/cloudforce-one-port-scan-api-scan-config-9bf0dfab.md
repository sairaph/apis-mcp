---
title: cloudforce-one-port-scan-api_scan-config
page_id: schema-cloudforce-one-port-scan-api-scan-config-9bf0dfab
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one-port-scan-api_scan-config

```yaml
{"type": "object", "properties": {"account_id": {"type": "string", "example": "abcd1234abcd1234abcd1234abcd1234"}, "frequency": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_frequency"}, "id": {"description": "Defines the Config ID.", "type": "string", "example": "uuid"}, "ips": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_ips"}, "ports": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_ports"}}, "required": ["id", "account_id", "ips", "frequency", "ports"], "title": "Config"}
```
