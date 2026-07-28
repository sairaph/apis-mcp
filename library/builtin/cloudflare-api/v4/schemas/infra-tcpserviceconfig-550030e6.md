---
title: infra_TcpServiceConfig
page_id: schema-infra-tcpserviceconfig-550030e6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# infra_TcpServiceConfig

```yaml
{"example": {"host": {"ipv4": "10.0.0.1", "network": {"tunnel_id": "0191dce4-9ab4-7fce-b660-8e5dec5172da"}}, "name": "postgres-db", "tcp_port": 5432, "type": "tcp"}, "allOf": [{"$ref": "#/components/schemas/infra_ServiceCommon"}, {"properties": {"app_protocol": {"type": "string", "example": "postgresql", "enum": ["postgresql", "mysql"], "nullable": true}, "tcp_port": {"type": "integer", "format": "int32", "example": 5432, "minimum": 1, "nullable": true}}, "type": "object"}]}
```
