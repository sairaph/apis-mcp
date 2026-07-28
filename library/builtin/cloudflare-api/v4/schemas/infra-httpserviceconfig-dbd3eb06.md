---
title: infra_HttpServiceConfig
page_id: schema-infra-httpserviceconfig-dbd3eb06
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# infra_HttpServiceConfig

```yaml
{"example": {"host": {"ipv4": "10.0.0.1", "network": {"tunnel_id": "0191dce4-9ab4-7fce-b660-8e5dec5172da"}}, "http_port": 8080, "https_port": 8443, "name": "web-app", "type": "http"}, "allOf": [{"$ref": "#/components/schemas/infra_ServiceCommon"}, {"properties": {"http_port": {"type": "integer", "format": "int32", "example": 8080, "minimum": 1, "nullable": true}, "https_port": {"type": "integer", "format": "int32", "example": 8443, "minimum": 1, "nullable": true}}, "type": "object"}]}
```
